package main

import (
	"fmt"
	"net/http"
	"os"

	"time"
)

func getUsername(w http.ResponseWriter, r *http.Request) (string, error) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return "", nil
	}

	if err := r.ParseForm(); err != nil {
		return "", err
	}

	username := r.FormValue("username")
	return username, nil
}

func handleDelayedReturn(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)

	time.Sleep(5 * time.Second)
	fmt.Fprintf(w, "<div>return</div>")
}

func main() {
	fs := http.FileServer(http.Dir("."))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// Check if the request is for the root or for index.html, and not a request for a static file
		if path == "/" || path == "/index.html" {
			component := chat_page()
			component.Render(r.Context(), w)
		} else {
			// Check if the file exists in the current directory
			if _, err := os.Stat("." + path); err == nil || !os.IsNotExist(err) {
				// File exists, let the FileServer handler serve it
				fs.ServeHTTP(w, r)
			} else {
				// File doesn't exist, return a 404
				http.NotFound(w, r)
			}
		}
	})

	http.HandleFunc("/submit_username", func(w http.ResponseWriter, r *http.Request) {
		username, err := getUsername(w, r)
		if err == nil {
			// Placeholder for where you'd use the username and render a response
			fmt.Fprintf(w, "<div>Hello, %s</div>", username)
		}
	})

	http.HandleFunc("/delayed_return", handleDelayedReturn)

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
