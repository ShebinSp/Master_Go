package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func checkAndSaveBody(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		s := fmt.Sprintf("%s is down\n", url)
		s += fmt.Sprintf("Error: %v", err)
		c <- s

	} else {
		defer resp.Body.Close()
		s := fmt.Sprintf("%s -> Status Code: %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := io.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1]
			file += ".txt"

			s += fmt.Sprintf("Writing response body to %s", file)
			err = os.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				s += fmt.Sprintf("Error writing file\n")
				c <- s
			}
			s += fmt.Sprintf("%s is UP\n", url)
		}
		c <- s
	}

}
func main() {
	urls := []string{"http://golang.org", "http://www.google.com", "http://netflix.com", "http://www.x.com",
		"http://www.unknownurl123.com", "http://www.medium.com", "http://www.facebook.com"}

	//1.
	c := make(chan string)

	for _, url := range urls {
		// We have launched 7 goroutines + main equals 8, and each goroutine is checking a web server if it's up or down.
		go checkAndSaveBody(url, c)
	}
	fmt.Println("No. of goroutines: ",runtime.NumGoroutine())

	// There are as many goroutines that send over the channel as the URLs so it's safe to iterate like this
	
	for i := 0; i < len(urls); i++ {
		fmt.Println(<- c)
	}
}
