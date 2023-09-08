package main

import (
	"fmt"
	"github.com/goranvasic/go-rename/pkg/files"
	"os"
	"path/filepath"
)

func main() {
	execute(os.Args)
}

func execute(args []string) {
	if len(args) == 3 {
		oldExt := args[1]
		newExt := args[2]
		if oldExt == "." {
			files.RenameAll(newExt)
		} else {
			files.RenameSpecific(oldExt, newExt)
		}
	} else {
		execName := filepath.Base(args[0])
		fmt.Println("[Usage]")
		fmt.Printf("Rename all files to [.foo] extension:\n  %s . foo\n", execName)
		fmt.Printf("Rename all [.foo] files to [.bar] extension:\n  %s foo bar\n", execName)
		os.Exit(1)
	}
}
