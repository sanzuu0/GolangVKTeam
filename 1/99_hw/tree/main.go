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
	createTree(path, "")
	return nil
}

func createTree(path string, prefix string) {

	//TODO + : Считывание директорий и их вывод отсортированно и с правильными префиксами, без отступов.
	//TODO + : настроить отступы
	//TODO + : решить проблему с последним отступом
	//TODO: настроить флаг вывода директорий без файлов

	entries, err := os.ReadDir(path)

	if err != nil {
		log.Fatal("Read not:", err)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	for index, entry := range entries {

		if index == len(entries)-1 {
			fmt.Println(prefix + "└───" + entry.Name())
		} else {
			fmt.Println(prefix + "├───" + entry.Name())
		}

		if entry.IsDir() {
			newPrefix := prefix
			if index == len(entries)-1 {
				newPrefix += "\t"
			} else {
				newPrefix += "│\t"
			}
			createTree(filepath.Join(path, entry.Name()), newPrefix)
		}
	}
}
