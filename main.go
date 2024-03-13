package main

import (
	"context"
	"net/http"
)

func getUsername(w http.ResponseWriter, r *http.Request) {
    // Ensure we're dealing with a POST request
    if r.Method != "POST" {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    // Parse form data
    if err := r.ParseForm(); err != nil {
        return
    }

    // Get the username from the form
    username := r.FormValue("username")

	return username, nil
}


func main() {
	http.HandleFunc("/submit_username", func(w http.ResponseWriter, r *http.Request) {

		username, err := getUsername(w, r)
		if err != nil {
			component := hello(username)
			component.Render(context.Background(), w) // Use w instead of os.Stdout
		}
		else {

		}
	})



	http.ListenAndServe(":8080", nil)
}
