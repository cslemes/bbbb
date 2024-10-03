package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	resume "github.com/cslemes/cris_term/cmd/app"
	"github.com/joho/godotenv"

	"tailscale.com/tsnet"
	"tailscale.com/types/logger"
)

// const (
// 	host = "localhost"
// 	port = 42069
// )

// You can wire any Bubble Tea model up to the middleware with a function that
// handles the incoming ssh.Session. Here we just grab the terminal info and
// pass it to the new model. You can also return tea.ProgramOptions (such as
// tea.WithAltScreen) on a session by session basis.
func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, active := s.Pty()
	if !active {
		wish.Fatalln(s, "no active terminal, skipping")
		return nil, nil
	}
	m := resume.InitialModel()

	return m, []tea.ProgramOption{
		tea.WithAltScreen(),

		tea.WithInput(pty.Slave),
		tea.WithOutput(pty.Slave),
	}
}

func main() {

	godotenv.Load("tailscale")
	srv := &tsnet.Server{
		Hostname: "ssh-blog",
		AuthKey:  os.Getenv("TSKEY"),
		Logf:     logger.Discard,
	}

	defer srv.Close()

	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start Tailscale server: %v", err)
	}

	ln, err := srv.Listen("tcp", ":2222")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer ln.Close()

	// -----

	s, err := wish.NewServer(
		wish.WithAddress(ln.Addr().String()),
		// wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),

		// wish.WithPublicKeyAuth(func(_ ssh.Context, key ssh.PublicKey) bool {
		// 	// needed for the public key on the ssh.Session, else it just
		// 	// returns 0s
		// 	return true
		// }),
		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Error("could not start server", "error", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Info("Starting SSH server")
	//log.Info("Starting SSH server", "host", host, "port", port)
	go func() {
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("could not start server", "error", err)
			done <- nil
		}
	}()

	<-done
	log.Info("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("could not stop server", "error", err)
	}
}
