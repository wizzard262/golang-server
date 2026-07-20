package main

import (
	"log"
	"net/http"

	// This import pulls in your Vercel serverless function package.
	// Used ONLY for local development and testing.
	handler "example.com/hello/api"
)

func main() {
	// LOCAL‑ONLY: Serve static files from the ./static directory
	// This mimics how Vercel serves your static assets via its CDN.
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// LOCAL‑ONLY: Wire up your Vercel function as a normal HTTP handler
	// so you can test it at http://localhost:8080/api/hello
	http.HandleFunc("/api/hello", handler.Handler)

	log.Println("Local test server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
