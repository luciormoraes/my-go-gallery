package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome website</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:lucio@castor.dev\">lucio@castor.dev</a>.</p>")
}

func pageNotFoundHandle(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page not found", http.StatusNotFound)
}

func faqHandle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1> FAQ Pages needs to imprve it</h1>`)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandlerFunc)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandle)
	r.NotFound(pageNotFoundHandle)

	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", r)
}
