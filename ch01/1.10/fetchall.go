package main

import (
	"fmt"
	"io"
	"net/http"
	url2 "net/url"
	"os"
	"time"
)

func main() {
	for i := 1; i <= 2; i++ { // 2回実行
		start := time.Now()
		ch := make(chan string)
		for _, url := range os.Args[1:] {
			go fetch(url, fileNameOf(url, i), ch) // ゴルーチンを開始
		}
		for range os.Args[1:] {
			fmt.Println(<-ch) // ch チャネルから受信
		}
		fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	}
}

func fetch(url string, fileName string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // ch チャネルに送信
		return
	}

	file, err := os.Create(fileName)
	if err != nil {
		ch <- fmt.Sprint(err) // ch チャネルに送信
		return
	}
	defer file.Close()

	nbytes, err := io.Copy(file, resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs	%7d	%s", secs, nbytes, url)
}

func fileNameOf(url string, th int) string {
	u, err := url2.Parse(url)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s_%dth.txt", u.Host, th)
}
