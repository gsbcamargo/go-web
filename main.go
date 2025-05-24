package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site. I mean, full-fledged-project-dude!</h1>")
}

func contactInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>Contact Page</h1><p>To get in touch, e-mail me 
	at <a href="mailto:gabriel@peticiona.ai">gabriel@peticiona.ai</a>.</p>`)
}

func pageNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>Sorry, the page you requested doesn't exist here.</h1>`)

}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactInfoHandler(w, r)
	default:
		pageNotFoundHandler(w, r)
	}
}

func main() {
	http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/contact", contactInfoHandler)
	fmt.Println("Starting server on port 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
