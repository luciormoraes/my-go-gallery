package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/luciormoraes/my-go-gallery/controllers"
	"github.com/luciormoraes/my-go-gallery/templates"
	"github.com/luciormoraes/my-go-gallery/views"
)

func main() {
	r := chi.NewRouter()

	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml"))))

	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", r)
}
