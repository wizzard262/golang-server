package handler

import (
	"fmt"      // formatting + printing utilities (Println, Fprintln, Sprintf, etc.)
	"net/http" // HTTP server + client (handlers, ResponseWriter, Request)
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go API  at http://localhost:8080/api/hello locally")
	fmt.Fprintln(w, "Hello from Go API on Vercel at https://golang-server-beta.vercel.app/api/hello")
}

// Note: as we are hosting on Vercel and it will handle the HTTP server we dont
// include things like: http.Handle() and http.ListenAndServe()
