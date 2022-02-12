// Dup1 exibe contagem e o texto das linhas que aparecem mais de uma vez na entrada.
// Ele lê de stdiun ou de uma lista de arquivos nomeados.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make((map[string]int))
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			file, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(file, counts)
			file.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(file *os.File, counts map[string]int) {

	input := bufio.NewScanner(file)

	for input.Scan() {
		counts[input.Text()]++
	}
	//NOTA: ignorando erros em potencial de input.Erro()
}
