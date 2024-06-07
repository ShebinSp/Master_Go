package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func checkAndSaveBody(url string) {
	resp, err := http.Get(url)

	if err != nil {
		log.Println(err)
		log.Printf("%s is down\n", url)
	} else {
		defer resp.Body.Close()
		log.Printf("%s -> Status Code: %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := io.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1]
			file += ".txt"

			log.Printf("Writing response body to %s\n", file)
			err = os.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	fmt.Println(strings.Repeat("-", 21))
}

func main() {
	t := time.Now().Second()

	urls := []string{"http://golang.org", "http://www.google.com", "http://netflix.com", "http://www.x.com",
		"http://www.unknownurl123.com", "http://www.youtube.com", "http://www.medium.com", "http://www.facebook.com"}

	for _, url := range urls {
		checkAndSaveBody(url)
	}

	// Imagine what happens if we are testing 1000 servers and each HTTP request takes two seconds. We'll end up with a
	// program that takes 2000 seconds or 34 minutes to complete, and this is unacceptable. A better approch would be
	// to create and launch a goroutine for ecah website we want to test. Instead of 34 minutes, it will take only a few
	// seconds and this makes a really big difference. Here the execution time is 9 seconds.
	fmt.Println("Time consumed: ", time.Now().Second()-t, "seconds") // total time taken to execute the program
}
