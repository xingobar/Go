package main

import (
	"io/ioutil"
	"io"
	"net/http"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)  // create channel
	for _, url := range os.Args[1:]{
		go fetch(url, ch) // goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // 從 channel 接收
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	
	if err != nil {
		ch<- fmt.Sprint(err) // 發送到 channel 
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch<- fmt.Sprint(err) // 發送到 channel
		return
	}

	secs := time.Since(start).Seconds()
	ch<- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}