package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"baliance.com/gooxml/document"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: unprotect <file1,[file2],...>")
		os.Exit(1)
	}
	for _, file := range os.Args[1:] {
		doc, err := document.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}
		doc.Settings.X().DocumentProtection = nil
		filename := path.Base(file)
		ext := path.Ext(filename)
		trunk := strings.TrimSuffix(filename, ext)
		doc.SaveToFile(trunk + "-unprotected" + ext)
	}
}
