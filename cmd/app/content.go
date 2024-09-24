package resume

import _ "embed"

//go:embed content/home.md
var homeContent string

//go:embed content/sobre.md
var sobreContent string

//go:embed content/contact.md
var contatoContent string

// TODO Put files on object storages for dynamic update

// func homePage() string {
// 	homeContent, err := os.ReadFile("content/home.md")
// 	if err != nil {
// 		panic(err)
// 	}
// 	return string(homeContent)
// }

// func sobrePage() string {
// 	sobreContent, err := os.ReadFile("content/sobre.md")
// 	if err != nil {
// 		panic(err)
// 	}
// 	return string(sobreContent)
// }

// func contatoPage() string {
// 	contatoContent, err := os.ReadFile("content/contact.md")
// 	if err != nil {
// 		panic(err)
// 	}
// 	return string(contatoContent)
// }

// 	splashContent = `
// ███████╗██████╗ ██╗███████╗     ██████╗  ██╗  ██╗ ██   ██╗
// ██╔════╝██╔══██╗██║██╔════╝     ██╔══██╗ ██║  ██║ ███╗ ██║
// ██║     ██████╔╝██║███████╗     ██████╔╝ ██║  ██║ ██╔████║
// ██║     ██╔══██╗██║╚════██║     ██╔══██╗ ██║  ██║ ██║ ███║
// ███████╗██║  ██║██║███████║ ██╗ ██║  ██║ ███████║ ██║  ██║
// ╚══════╝╚═╝  ╚═╝╚═╝╚══════╝ ╚═╝ ╚═╝  ╚═╝ ╚══════╝ ╚═╝  ╚═╝
//`
