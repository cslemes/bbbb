package pages

import (
	"os"
)

// TODO: Put files on object storages for dynamic update

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

func splashContent(xTerm int) string {

	splashHD, err := os.ReadFile("content/splash.txt")
	if err != nil {
		panic(err)
	}

	splashLow, err := os.ReadFile("content/splashlow.txt")
	if err != nil {
		panic(err)
	}
	if xTerm == 256 {
		return string(splashHD)

	} else {
		return string(splashLow)
	}

}
