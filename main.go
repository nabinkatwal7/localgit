package main

import (
	"flag"
	"fmt"
)

func scan(path string) {
	fmt.Printf("Found folders: \n\n")
	repositories := recursiveScanFolder(path)
	filePath := getDotFilePath()
	addNewSliceElementsToFile(filePath, repositories)
	fmt.Printf("\n\nSuccessfully added\n\n")
}

func stats(email string) {
	print(email)
}

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "add a new folder to scan for git repos.")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)
}
