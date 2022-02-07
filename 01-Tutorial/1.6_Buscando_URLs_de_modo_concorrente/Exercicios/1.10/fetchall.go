package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	data := make([]string, 0)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) //inicia uma gorotina
	}

	for range os.Args[1:] {
		content := <-ch
		data = append(data, content)
	}
	writeLog(data)
	fmt.Printf("%.2fs decorrido\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // envia para o canal ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() //evita vazamento de recursos

	if err != nil {
		ch <- fmt.Sprintf("enquanto lia %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	content := fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
	ch <- content
}

func writeLog(data []string) error {
	file, err := os.Create("log.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	for _, info := range data {
		_, err := file.WriteString(info + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
