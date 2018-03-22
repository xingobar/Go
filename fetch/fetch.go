package main

import (
	"io"
	"strings"
	// "io/ioutil"
	"fmt"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url) // 發出 http 請求
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close() // 避免資料洩漏
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Print("[*] status code => %s", resp.Status);
		fmt.Printf("%s",b)
	}
}