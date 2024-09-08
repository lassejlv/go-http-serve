package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {

	dir := flag.String("dir", ".", "the directory to serve files from. Defaults to the current dir")
	port := flag.String("port", "3000", "the port to serve HTTP on. Defaults to 3000")

	flag.Parse()

	router := http.NewServeMux()

	// Handle Static Files
	router.Handle("/", http.FileServer(http.Dir(*dir)))

	// Start the server
	fmt.Printf("Go HTTP serve - Server running on http://localhost:%s\n", *port)
	fmt.Printf("Static files path: %s\n", *dir)
	fmt.Printf("Press Ctrl + C to stop the server\n")
	log.Fatal(http.ListenAndServe(":"+*port, router))
}
