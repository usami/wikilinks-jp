package main

import (
	"fmt"
	"os"

	"github.com/usami/wikilinks-jp/internal/linker"
)

func Usage() {
	fmt.Printf("Usage: %s [category] [annotation-file] [html-dir] [title-pageid-file] [output-file]\n", os.Args[0])
}

func main() {
	if len(os.Args) != 6 {
		Usage()
		os.Exit(1)
	}

	l := linker.NewLinker(os.Args[1])
	l.Load(os.Args[2], os.Args[3], os.Args[4])

	l.Run()

	l.Output(os.Args[5])
}
