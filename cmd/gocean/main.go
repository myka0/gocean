package main

import (
	"log"

	"github.com/myka0/gocean"
)

func main() {
	p := gocean.NewProgram()

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
