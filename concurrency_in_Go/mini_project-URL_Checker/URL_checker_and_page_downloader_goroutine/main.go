package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

func checkAndSaveBody(wg *sync.WaitGroup, url string) {
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
	wg.Done()
	fmt.Println(strings.Repeat("-", 21))
}
func main() {
	// Here to make the application concurrently, We will use goroutines and waitgroups for synchronization.
	// In fact, I want the application to check the URLs and save the web pages not in a sequential manner,
	// but at the same time. This means that i'll create and launch a goroutine for each url and I'll use
	// waitgroups to synchroize them.

	t := time.Now().Second()

	urls := []string{"http://golang.org", "http://www.google.com", "http://netflix.com", "http://www.x.com",
		"http://www.unknownurl123.com", "http://www.youtube.com", "http://www.medium.com", "http://www.facebook.com"}

	// Creating a new instance of waitgroup
	var wg = sync.WaitGroup{} // wg is used to wait for all the goroutines to finish

	wg.Add(len(urls)) // adding the counter to tell the waitgroup there are `counter` numbers of goroutines.

	// lauching a goroutine. A goroutine is in fact a function invoked in a special manner. When we use `go` keyword
	// just before the function's name, it becomes a goroutine.
	for _, url := range urls {
		go checkAndSaveBody(&wg, url)
	}

	fmt.Println("No. of Goroutines: ", runtime.NumGoroutine())

	wg.Wait() // to block the execution of main() until the goroutines in the waitgroup have successfully completed.

	fmt.Println("Time consumed: ", time.Now().Second()-t, "seconds") // total time taken to execute the program

	// By using goroutines we saved around 4-5 seconds.
}
