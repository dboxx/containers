package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	port1, port2 int
)

func init() {
	flag.IntVar(&port1, "port1", 8080, "port 1")
	flag.IntVar(&port2, "port2", 9090, "port 2")
	flag.Parse()
}

func main() {
	http.HandleFunc("/", commonHandler)

	var wg sync.WaitGroup

	wg.Add(1)
	go runHTTPServer(fmt.Sprintf(":%d", port1), &wg)

	wg.Add(1)
	go runHTTPServer(fmt.Sprintf(":%d", port2), &wg)

	fmt.Printf("Started 2 listeners: [%d|%d]\n", port1, port2)

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
		runHTTPClient(fmt.Sprintf("http://localhost:%d/test-port1", port1))
		runHTTPClient(fmt.Sprintf("http://localhost:%d/test-port2", port2))
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

type Response struct {
	Server Server `json:"server"`
	Client Client `json:"client"`
}

type Server struct {
	Addrs []string `json:"addrs,omitempty"`
}

type Client struct {
	Request       string `json:"request,omitempty"`
	RemoteAddr    string `json:"remote-addr,omitempty"`
	XRealIP       string `json:"x-real-ip,omitempty"`
	XForwardedFor string `json:"x-forwarded-for,omitempty"`
}

func commonHandler(w http.ResponseWriter, r *http.Request) {
	rs := &Response{}

	// get list of available addresses
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				// check if IPv4 or IPv6 is not nil
				if ipnet.IP.To4() != nil || ipnet.IP.To16() != nil {
					// print available addresses
					rs.Server.Addrs = append(rs.Server.Addrs, ipnet.IP.String())
				}
			}
		}
	}

	rs.Client.Request = r.URL.Path[1:]
	rs.Client.RemoteAddr, _, _ = net.SplitHostPort(r.RemoteAddr)
	if rs.Client.RemoteAddr == "" {
		rs.Client.RemoteAddr = "127.0.0.1"
	}
	rs.Client.XRealIP = r.Header.Get("X-Real-Ip")
	rs.Client.XForwardedFor = r.Header.Get("X-Forwarded-For")

	json.NewEncoder(w).Encode(rs)
}
