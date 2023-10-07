package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func initClient() http.Client {
	caCert, err := os.ReadFile(*certFile)
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	f, err := os.OpenFile("tlskeys.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	// Create TLS configuration with the certificate of the server
	tlsConfig := &tls.Config{
		RootCAs:      caCertPool,
		KeyLogWriter: f,
	}

	client := &http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	return *client
}
