package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site. I mean, !</h1>")
}

func contactInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>Contact Page</h1><p>To get in touch, e-mail me 
	at <a href="mailto:gabriel@peticiona.ai">gabriel@peticiona.ai</a>.</p>`)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactInfoHandler)
	fmt.Println("Starting server on port 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
