package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
)

func main() {
	if err := license.SetLegacyLicenseKey(unidocLicense); err != nil {
		fmt.Println("error installing unioffice license:", err)
		os.Exit(1)
	}

	unioffice.DisableLogging()

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: unprotect <file1,[file2],...>")
		os.Exit(1)
	}
	for _, file := range os.Args[1:] {
		doc, err := document.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error opening %s: %v\n", file, err)
			continue
		}
		doc.Settings.X().DocumentProtection = nil
		filename := path.Base(file)
		ext := path.Ext(filename)
		trunk := strings.TrimSuffix(filename, ext)
		doc.SaveToFile(trunk + "-unprotected" + ext)
		fmt.Println(trunk + "-unprotected" + ext)
	}
}
