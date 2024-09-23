package resume

import "os"

func homePage() string {
	homeContent, err := os.ReadFile("content/home.md")
	if err != nil {
		panic(err)
	}
	return string(homeContent)
}

func sobrePage() string {
	sobreContent, err := os.ReadFile("content/sobre.md")
	if err != nil {
		panic(err)
	}
	return string(sobreContent)
}

func contatoPage() string {
	contatoContent, err := os.ReadFile("content/contact.md")
	if err != nil {
		panic(err)
	}
	return string(contatoContent)
}

// 	splashContent = `
// ███████╗██████╗ ██╗███████╗     ██████╗  ██╗  ██╗ ██   ██╗
// ██╔════╝██╔══██╗██║██╔════╝     ██╔══██╗ ██║  ██║ ███╗ ██║
// ██║     ██████╔╝██║███████╗     ██████╔╝ ██║  ██║ ██╔████║
// ██║     ██╔══██╗██║╚════██║     ██╔══██╗ ██║  ██║ ██║ ███║
// ███████╗██║  ██║██║███████║ ██╗ ██║  ██║ ███████║ ██║  ██║
// ╚══════╝╚═╝  ╚═╝╚═╝╚══════╝ ╚═╝ ╚═╝  ╚═╝ ╚══════╝ ╚═╝  ╚═╝
//`
