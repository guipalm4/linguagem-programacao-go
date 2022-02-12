// ex1.4 prints the count and text of lines that appear more than once in the
// input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	foundIn := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, foundIn)
	} else {
		for _, arg := range files {
			file, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(file, counts, foundIn)
			file.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%v\t%s\n", n, foundIn[line], line)
		}
	}
}

func countLines(file *os.File, counts map[string]int, foundIn map[string][]string) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !in(file.Name(), foundIn[line]) {
			foundIn[line] = append(foundIn[line], file.Name())
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

func in(needle string, strings []string) bool {
	for _, s := range strings {
		if needle == s {
			return true
		}
	}
	return false
}
