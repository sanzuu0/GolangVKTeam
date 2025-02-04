package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	createTree(out, path, "", printFiles)
	return nil
}

func createTree(out io.Writer, path string, prefix string, printFiles bool) {

	entries, err := os.ReadDir(path)

	if err != nil {
		log.Fatalf("cannot read directory %s: %v", path, err)
	}

	entries = sorting(entries, printFiles)

	for index, entry := range entries {

		isLastEntry := index == len(entries)-1
		connector := "├───"

		if isLastEntry {
			connector = "└───"
		}

		fullPath := filepath.Join(path, entry.Name())

		if entry.IsDir() {

			if isLastEntry {
				fmt.Fprintf(out, "%s%s%s\n", prefix, connector, entry.Name())
				newPrefix := prefix + "\t"
				createTree(out, fullPath, newPrefix, printFiles)
			} else {
				fmt.Fprintf(out, "%s%s%s\n", prefix, connector, entry.Name())
				newPrefix := prefix + "│\t"
				createTree(out, fullPath, newPrefix, printFiles)
			}

		} else if printFiles {

			info, err := entry.Info()
			if err != nil {
				log.Fatal(err)
			}
			fileSize := fmt.Sprintf("%db", info.Size())
			if info.Size() == 0 {
				fileSize = "empty"
			}

			fmt.Fprintf(out, "%s%s%s (%s)\n", prefix, connector, entry.Name(), fileSize)

		}
	}

}

func sorting(entries []os.DirEntry, printFiles bool) []os.DirEntry {
	var sorted []os.DirEntry
	if !printFiles {
		for _, entry := range entries {
			if entry.IsDir() {
				sorted = append(sorted, entry)
			}
		}
	} else {
		sorted = entries
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Name() < sorted[j].Name()
	})
	return sorted
}
