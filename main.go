package main

import (
	"GenSecureSign/api"
	"crypto/tls"
	"log"
	"net/http"
	"os"
)

func ensureDirectoryExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	return err
}

func main() {

	cert, err := tls.LoadX509KeyPair("./newcertificate.pem", "./newprivatekey.pem")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}

	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}

	if err := ensureDirectoryExists("./static/storage/signed"); err != nil {
		panic(err)
	}

	if err := ensureDirectoryExists("./static/storage/uploads"); err != nil {
		panic(err)
	}

	router := api.InitRoutes()

	server := &http.Server{
		Addr:      ":8080",
		Handler:   router,
		TLSConfig: tlsConfig,
	}

	log.Fatal(server.ListenAndServeTLS("", ""))
}
