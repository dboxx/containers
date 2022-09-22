package main

import (
    "sync"
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", hello)

    var wg sync.WaitGroup

    wg.Add(1)
    go runHTTP(":8080", &wg)

    wg.Add(1)
    go runHTTP(":9090", &wg)

    fmt.Println("Started 2 listeners: [8080|9090]")
    wg.Wait()
}

func runHTTP(addr string, wg *sync.WaitGroup) {
    defer wg.Done()

    http.ListenAndServe(addr, nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello", r.URL.Path[1:])
}
