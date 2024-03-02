package main

import (
	"context"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		component := hello("John")
		component.Render(context.Background(), w) // Use w instead of os.Stdout
	})

	http.ListenAndServe(":8080", nil)
}
