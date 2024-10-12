
![](/images/bubble-blog.png)

<a href="https://www.freepik.com/free-vector/gradient-candy-pastel-color-text-effect_49681775.htm">Image by pikisuperstar on Freepik</a>

# Bubble Blog

Bubble Blog is a delightful terminal-based blog application built with Go, leveraging the power of Charm libraries to create an interactive and visually appealing user interface right in your terminal.

![](/images/homescreen.png)

## Features

- 🖥️ Terminal-based UI: Enjoy a sleek, responsive interface in your favorite terminal
- 📑 Markdown Rendering: Write your blog posts in Markdown and see them beautifully rendered
- 🗂️ File-based Content Management: Easily manage your blog posts as Markdown files
- 🎨 Customizable Themes: Adapt the look and feel to your preferences
- 🚀 Fast and Lightweight: Enjoy a snappy experience, even on older hardware
- ⚙️ Configurable: Easily customize your blog's behavior and appearance through a YAML configuration file
- 🐳 Docker Support: Quickly deploy your blog using Docker


## Try It Out!

You can experience Bubble Blog right now without any installation! A live instance is running and accessible via SSH:

```
ssh cris.run
```

This will connect you to a public instance of Bubble Blog, allowing you to explore its features and interface immediately. Feel free to browse through the blog posts and navigate the terminal UI to get a feel for how Bubble Blog works.

## Technologies Used

- [Go](https://golang.org/): The core programming language
- [Bubble Tea](https://github.com/charmbracelet/bubbletea): Terminal UI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss): Style definitions for terminal applications
- [Glamour](https://github.com/charmbracelet/glamour): Markdown rendering
- [Wish](https://github.com/charmbracelet/wish): SSH server for running Bubble Tea programs

## Installation

### Standard Installation

1. Ensure you have Go installed on your system.
2. Clone this repository:
   ```
   git clone https://github.com/yourusername/bubble-blog.git
   ```
3. Navigate to the project directory:
   ```
   cd bubble-blog
   ```
4. Install dependencies:
   ```
   go mod tidy
   ```
5. Build the project:
   ```
   go build -o bbbb main.go
   ```

### Docker Installation

Bubble Blog comes with a Dockerfile for easy containerization and deployment. Here's how to use it:

1. Ensure you have Docker installed on your system.
2. Clone the repository (if you haven't already):
   ```
   git clone https://github.com/cslemes/bbbb.git
   ```
3. Navigate to the project directory:
   ```
   cd bbbb
   ```
4. Build the Docker image:
   ```
   docker build -t bubble-blog .
   ```
5. Run the container:
   ```
   docker run -p 42069:42069 bubble-blog
   ```

## Configuration

Bubble Blog uses a `config.yaml` file for configuration. Here's an explanation of the configuration options:

```yaml
server:
  host: 0.0.0.0  # The IP address to bind the server to (0.0.0.0 allows connections from any IP)
  port: 42069    # The port number for the SSH server

theme:
  color: special # The color theme to use. Options: subtle, highlight, special

navigation:
  showingSplash: true # Whether to show the splash screen on startup
```

You can customize these settings to change the server's binding address and port, adjust the color theme, and control whether the splash screen is displayed on startup.

## Content Management

Bubble Blog uses a file-based content management system. Blog posts are stored as Markdown files in the `posts` directory. Here's how it works:

- Each blog post is a separate Markdown file in the `posts` directory.
- The filename (without the `.md` extension) is used as the post's identifier.
- The content of the Markdown file is rendered and displayed in the blog.

To add a new blog post:

1. Create a new Markdown file in the `posts` directory.
2. Name the file descriptively, e.g., `my-first-blog-post.md`.
3. Write your blog post content in Markdown format.
4. Save the file, and it will automatically appear in your Bubble Blog.

## Usage

1. Start the Bubble Blog server:
   ```
   ./bbbb
   ```
2. Connect to the server using SSH:
   ```
   ssh localhost -p 42069  # Or the port you specified in config.yaml
   ```
3. Navigate through the blog using the keyboard:
   - Use arrow keys or `h`, `j`, `k`, `l` to move around
   - Press `Enter` to select
   - Use `q` or `Ctrl+C` to quit

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- The amazing [Charm](https://charm.sh/) team for their fantastic libraries


## TODOs

- [ ] Implement user authentication for admin access
- [ ] Implement a search functionality for blog posts
- [ ] Add support for comments on blog posts
- [ ] Create a configuration file for easy customization of colors and styles
- [ ] Implement pagination for the blog post list
- [ ] Add support for image embedding in blog posts (possibly using ASCII art or terminal graphics)
- [ ] Create a simple plugin system for extending functionality
- [ ] Implement a basic analytics system to track post views
- [ ] Add support for multiple authors with different permissions
- [ ] Implement a tagging system for blog posts
- [ ] Add a "featured posts" section on the home page
- [ ] Frontmatter for blog posts
- [ ] Multiples Layouts to choose


