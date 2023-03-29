package main

import (
	"fmt"
	"github.com/goranvasic/go-rename/pkg/files"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) == 3 {
		oldExt := os.Args[1]
		newExt := os.Args[2]
		if oldExt == "." {
			files.RenameAll(newExt)
		} else {
			files.RenameSpecific(oldExt, newExt)
		}
	} else {
		execName := filepath.Base(os.Args[0])
		fmt.Println("[Usage]")
		fmt.Printf("Rename all files to [.foo] extension:\n  %s . foo\n", execName)
		fmt.Printf("Rename all [.foo] files to [.bar] extension:\n  %s foo bar\n", execName)
		os.Exit(1)
	}
}
