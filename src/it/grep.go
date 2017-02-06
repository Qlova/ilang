package main

import (
	"bufio"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var mainFile string

func grep(wg *sync.WaitGroup, path string) {
	defer wg.Done()

	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {	
		if mainFile != "" {
			return
		}
		if strings.Contains(scanner.Text(), "software") {
			mainFile = path
			return
		}
		if strings.Contains(scanner.Text(), "ソフトウェア") {
			mainFile = path
			return
		}
		if strings.Contains(scanner.Text(), "программного") {
			mainFile = path
			return
		}
		if strings.Contains(scanner.Text(), "软件") {
			mainFile = path
			return
		}
		if strings.Contains(scanner.Text(), "grate") {
			mainFile = path
			return
		}
		if strings.Contains(scanner.Text(), "gui") {
			if strings.Contains(scanner.Text(), "{") {
				mainFile = path
				return
			}
		}
	}
}
