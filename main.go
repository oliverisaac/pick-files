package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func isFile(fname string) bool {
	info, err := os.Stat(fname)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func isDir(dname string) bool {
	info, err := os.Stat(dname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func main() {
	helpShorthand := " (shorthand)"

	showDirs := false
	helpDirs := "Show directories"
	flag.BoolVar(&showDirs, "dir", false, helpDirs)
	flag.BoolVar(&showDirs, "d", false, helpDirs+helpShorthand)

	showFiles := false
	helpFiles := "Show files (default)"
	flag.BoolVar(&showFiles, "files", false, helpFiles)
	flag.BoolVar(&showFiles, "f", false, helpFiles+helpShorthand)

	showAll := false
	helpAll := "Show both files and directories"
	flag.BoolVar(&showFiles, "all", false, helpAll)
	flag.BoolVar(&showFiles, "a", false, helpAll+helpShorthand)

	flag.Parse()

	// If user wants all, show all
	if showAll {
		showFiles = true
		showDirs = true
	}

	// If user specs neither files nor dirs, show files
	if !showFiles && !showDirs {
		showFiles = true
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Println(err)
			continue
		}
		line := scanner.Text()

		if showDirs && isDir(line) {
			fmt.Println(line)
		}
		if showFiles && isFile(line) {
			fmt.Println(line)
		}
	}
}
