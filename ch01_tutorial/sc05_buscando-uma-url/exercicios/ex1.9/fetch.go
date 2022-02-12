package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		prefix := "http://"

		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}

		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fecth: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		fmt.Printf("HTTP status: %d\n", resp.StatusCode)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fecth: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
