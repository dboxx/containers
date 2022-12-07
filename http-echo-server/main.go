package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	http.HandleFunc("/", hello)

	var wg sync.WaitGroup

	wg.Add(1)
	go runHTTPServer(":8080", &wg)

	wg.Add(1)
	go runHTTPServer(":9090", &wg)

	fmt.Println("Started 2 listeners: [8080|9090]")

	wg.Add(1)
	go runHTTPChecker(&wg)

	wg.Wait()
}

func runHTTPServer(addr string, wg *sync.WaitGroup) {
	defer wg.Done()

	http.ListenAndServe(addr, nil)
}

func runHTTPChecker(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		time.Sleep(3 * time.Second)
		runHTTPClient("http://localhost:8080/test-8080")
		runHTTPClient("http://localhost:9090/test-9090")
	}
}

func runHTTPClient(rawURL string) {
	resp, err := http.Get(rawURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "HTTP GET ERROR %s: %v", rawURL, err)
		return
	}

	defer resp.Body.Close()

	fmt.Fprintf(os.Stdout, "HTTP GET OK %s: ", rawURL)
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		fmt.Fprintf(os.Stderr, "HTTP GET ERROR %s: %v", rawURL, err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello", r.URL.Path[1:])
}
