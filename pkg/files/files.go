package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	allFiles          = getAllFiles()
	existingFileNames []string
)

func RenameAll(newExt string) {
	for _, file := range allFiles {
		existingFileNames = append(existingFileNames, file.Name())
	}
	for _, file := range allFiles {
		oldExt := getExtension(file)
		newName := getNewName(file, oldExt, newExt)
		rename(file, newName)
	}
}

func RenameSpecific(oldExt, newExt string) {
	var filesToRename []os.DirEntry
	for _, file := range allFiles {
		if getExtension(file) == oldExt {
			filesToRename = append(filesToRename, file)
		} else {
			existingFileNames = append(existingFileNames, file.Name())
		}
	}
	for _, file := range filesToRename {
		newName := getNewName(file, oldExt, newExt)
		rename(file, newName)
	}
}

func getNewName(file os.DirEntry, oldExt, newExt string) string {
	oldName := file.Name()
	if oldExt == "" {
		return oldName + "." + newExt
	}
	return oldName[:len(oldName)-len(oldExt)] + newExt
}

func rename(file os.DirEntry, newName string) {
	oldName := file.Name()
	fileExists := false
	for _, existingFile := range existingFileNames {
		if newName == existingFile {
			fileExists = true
			break
		}
	}
	if !fileExists {
		fmt.Printf("Renaming: %s > %s\n", oldName, newName)
		err := os.Rename(oldName, newName)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(1)
		}
		existingFileNames = append(existingFileNames, newName)
	} else {
		fmt.Printf("Skipping: %s > %s - File already exists.\n", oldName, newName)
	}
}

func getExtension(file os.DirEntry) string {
	return strings.Trim(filepath.Ext(file.Name()), ".")
}

func getAllFiles() (allFiles []os.DirEntry) {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("ERROR: Could not read files under the current directory: %v\n", err)
		os.Exit(1)
	}
	for _, file := range files {
		if !file.IsDir() {
			allFiles = append(allFiles, file)
		}
	}
	return
}
