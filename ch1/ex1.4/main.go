package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countlines(os.Stdin, "<os.Stdin>", counts, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex1.4 : %v\n", err)
				continue
			}
			countlines(f, arg, counts, filenames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			files := filenames[line]
			fmt.Printf("%d\t%s\t%v\n", n, line, files)
		}
	}
}

func contains(slice []string, element string) bool {
	for _, str := range slice {
		if str == element {
			return true
		}
	}
	return false
}

func countlines(f *os.File, filename string, counts map[string]int, filenames map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		key := input.Text()
		counts[key]++

		if !contains(filenames[key], filename) {
			filenames[key] = append(filenames[key], filename)
		}
	}
}
