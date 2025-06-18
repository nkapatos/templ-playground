package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/nkapatos/templ-playground/esbuild-basic-tailwindcss/templ/views"
)

func main() {
	// Route handlers
	http.Handle("/", templ.Handler(views.Home()))
	http.Handle("/about", templ.Handler(views.About()))
	http.Handle("/contact", templ.Handler(views.Contact()))
	http.Handle("/daisyui", templ.Handler(views.DaisyUI()))

	// Static assets
	fs := http.FileServer(http.Dir("dist"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
