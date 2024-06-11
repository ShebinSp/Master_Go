package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func checkUrl(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		s := fmt.Sprintf("%s is down\n", url)
		s += fmt.Sprintf("Error: %v", err)
		fmt.Println(s)
		c <- url

	} else {
		defer resp.Body.Close()
		s := fmt.Sprintf("%s -> Status Code: %d\n", url, resp.StatusCode)
		s += fmt.Sprintf("%s is UP\n", url)
		fmt.Println(s)
		c <- url
	}
}

func main() {

	urls := []string{"http://golang.org", "http://www.google.com", "http://netflix.com", "http://www.x.com",
		"http://www.unknownurl123.com", "http://www.medium.com", "http://www.facebook.com"}

	c := make(chan string)

	for _, url := range urls {
		go checkUrl(url, c)
	}

	fmt.Println("No. of goroutines running: ", runtime.NumGoroutine())

	// for i := 0; i < len(urls); i++ {
	// 	fmt.Println(<-c)
	// }

	// for {
	// 	go checkUrl(<-c, c)
	// 	fmt.Println("---------------------------------------")
	// 	time.Sleep(time.Second * 2)
	// }

	// for url := range c {
	// 	time.Sleep(time.Second * 2)
	// 	go checkUrl(url, c)
	// }

	for url := range c {
		go func(u string) {
			time.Sleep(time.Second * 2)
			checkUrl(u, c)
		}(url)
		
	}
}
