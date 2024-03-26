package main

import (
	"fmt"
	"github.com/go-redis/redis/v8" // Import the Redis package
	"net/http"
	"os"
	"time"
)

// Conversation represents a single message in the conversation,
// including the role of the sender and the content of the message.
type Conversation struct {
	Role    string // Role can be "user", "bot", etc., depending on your application's needs.
	Content string // Content stores the text of the message.
}

// ConversationHistory holds a list of conversations.
type ConversationHistory struct {
	Conversations []Conversation
}

// AddConversation adds a new conversation to the history.
func (ch *ConversationHistory) AddConversation(convo Conversation) {
	ch.Conversations = append(ch.Conversations, convo)
}

// PrintConversations prints all conversations in the history.
func (ch *ConversationHistory) PrintConversations() {
	for i, convo := range ch.Conversations {
		fmt.Printf("%d: [%s] %s\n", i, convo.Role, convo.Content)
	}
}

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

func get_form_value(r *http.Request, w http.ResponseWriter, form_value string) (string, string) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return "", "failed to parse string"
	}

	message := r.FormValue(form_value)
	fmt.Printf("Received message: %s\n", message)

	return message, ""
}

func main() {
	fs := http.FileServer(http.Dir("."))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// Check if the request is for the root or for index.html, and not a request for a static file
		if path == "/" || path == "/index.html" {
			// Attempt to retrieve the cookie by name
			cookie, err := r.Cookie("phone_number")
			// Check if the cookie was found
			if err != nil {
				if err == http.ErrNoCookie {
					// Cookie not found, handle accordingly
					fmt.Printf("Cookie is not set.")

					component := login_page()
					component.Render(r.Context(), w)
					// Optionally, you can set the cookie here if it's necessary
				} else {
					// Some other error occurred
					http.Error(w, "An error occurred while checking the cookie", http.StatusInternalServerError)
				}
			} else {
				// Cookie is found, handle accordingly
				fmt.Printf("Cookie is set: %s", cookie.Value)
				component := chat_page()
				component.Render(r.Context(), w)
			}

			// Check to see if there is a cookie set named phone number:

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

	http.HandleFunc("/submit_chat", func(w http.ResponseWriter, r *http.Request) {
		message, _ := get_form_value(r, w, "user_prompt")

		component := user_input_submitted(message)
		component.Render(r.Context(), w)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		phone_number, _ := get_form_value(r, w, "phone_number")

		cookie := &http.Cookie{
			Name:     "phone_number",                 // Cookie name
			Value:    phone_number,                   // Cookie value
			Expires:  time.Now().Add(24 * time.Hour), // Set the expiration time for 1 day
			Path:     "/",                            // Path for the cookie
			HttpOnly: true,                           // Makes the cookie inaccessible to JavaScript running in the browser
		}

		http.SetCookie(w, cookie)

		chat_body().Render(r.Context(), w)

	})

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}

}
