****Go + VS Code: Basic Hello World****  

* Install Go from: https://go.dev/dl  *(Windows installer: go1.26.5.windows-amd64.msi)*  into: *C:\Program Files\Go*   
* Open PowerShell --> Run as Administrator:  
  Verify installation (from any folder): `go version` reports: `go version go1.26.5 windows/amd64`   
* Create a new GitHub repo *(with README.md and a Go .gitignore)*   
* Clone repo to folder: *C:\DEV\Repositories\GitHub\golang-server*  
* Initialize a Go module:  VS Code --> Terminal --> *C:\DEV\Repositories\GitHub\golang-server\code*  
`go mod init example.com/hello` *(Creates the Module **go.mod** in C:\DEV\Repositories\GitHub\golang-server\code)*  
`go mod tidy` *(Adds any required dependencies and removes unused dependencies)*  
* Add file: *C:\DEV\Repositories\GitHub\golang-server\code\static\index.html* with some HTML content.  
* Add file: *C:\DEV\Repositories\GitHub\golang-server\code\api\hello.go*  
	```
	package handler

	import (
		"fmt"      // formatting utilities
		"net/http" // HTTP request + response types
	)

	func Handler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go API  at http://localhost:8080/api/hello locally!")
		fmt.Fprintln(w, "Hello from Go API on Vercel at TODO !")
	}
	// Note: as we are hosting on Vercel and it will handle the HTTP server we dont 
	// include things like: http.Handle() and http.ListenAndServe()
	```
* Add file: *C:\DEV\Repositories\GitHub\golang-server\code\local-main.go*
	```
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
	```
* Add file: *C:\DEV\Repositories\GitHub\golang-server\code\vercel.json*
	```
	{
	  // "rewrites" tells Vercel to internally redirect certain URLs
	  "rewrites": [
		{
		  // When the user visits the root URL: https://your-app.vercel.app/
		  "source": "/",

		  // Serve the static file located at /static/index.html instead
		  // This makes your index.html act like the homepage
		  "destination": "/static/index.html"
		}
	  ]
	}
	```

* Open *C:\DEV\Repositories\GitHub\golang-server\code\api* folder in VS Code  
  Click **hello.go** and allow VS Code to install recommended Go tools.  
  
* Run the program locally in VS Code:  
  VS Code --> Terminal --> `go run local-main.go`  *(Compiles and runs code)*  
  Terminal shows: *2026/07/20 22:59:53 Local test server running at http://localhost:8080*

* Open browser:  
  *http://localhost:8080/* --> shows then content from : **index.html**  
  *http://localhost:8080/api/hello* --> shows dynamic text: *"Hello from Go API at: http://localhost:8080/api/hello"*  #
  
  ---
  **Deploy to Vercel**
 * Login to Vercel: https://vercel.com using my Github Account.
 * Add: New Project --> Import Github Repo --> **golang-server** --> import
 * Application Preset: Go
 * Set **Root Directory** to the subfolder: Code
 * Deploy *(still we need to create deployment and give it the URL to the CODE folder: https://github.com/wizzard262/golang-server/tree/main/code)* 
 * https://golang-server-snowy.vercel.app/ shows: ??