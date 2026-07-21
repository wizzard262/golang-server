package main

import (
	"log"
	"net/http"

	// This import pulls in your Vercel serverless function package. Used ONLY for local development and testing.
	handler "example.com/api" // Import your local API package. Path = module root (example.com) + folder (/api)
)

func main() {
	// LOCAL‑ONLY: Serve static files from the ./public directory
	// This mimics how Vercel serves your static assets via its CDN (Content Delivery Network).
	fs := http.FileServer(http.Dir("./public"))

	// LOCAL‑ONLY: Wire up the static file server to the root path ("/") to test at: http://localhost:8080
	http.Handle("/", fs)

	// LOCAL‑ONLY: Wire up your Vercel functions as a normal HTTP handler to test at:
	// http://localhost:8080/api/hello & http://localhost:8080/api/weather
	http.HandleFunc("/api/hello", handler.Hello)
	http.HandleFunc("/api/weather", handler.Weather)

	log.Println("Local test server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
