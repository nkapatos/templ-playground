package main

import (
	"fmt"
	"github.com/a-h/templ"
	"github.com/evanw/esbuild/pkg/api"
	"github.com/nkapatos/templ-playground/esbuild-go/templ/views"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func httpserver() {
	// Parse esbuild dev server URL
	esbuildURL, err := url.Parse("http://localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	// Create reverse proxy to esbuild server
	proxy := httputil.NewSingleHostReverseProxy(esbuildURL)

	// Proxy asset requests to esbuild dev server
	http.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = r.URL.Path[len("/assets"):]
		proxy.ServeHTTP(w, r)
	})

	// Proxy /esbuild for EventSource live reload
	http.HandleFunc("/esbuild", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	http.Handle("/", templ.Handler(views.Home()))
	http.Handle("/about", templ.Handler(views.About()))
	http.Handle("/contact", templ.Handler(views.Contact()))
	http.Handle("/daisyui", templ.Handler(views.DaisyUI()))

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func devserver() {

	ctx, err := api.Context(api.BuildOptions{
		EntryPoints: []string{"assets/css/index.css", "assets/ts/main.ts"},
		Outdir:      "dist",
		Bundle:      true,
		LogLevel:    api.LogLevelInfo,
	})

	if err != nil {
		panic("oooops")
	}

	watch_error := ctx.Watch(api.WatchOptions{})

	if watch_error != nil {
		panic("error while watching")
	}

	_, serve_error := ctx.Serve(api.ServeOptions{})
	if serve_error != nil {
		panic(serve_error)
	}
	// TODO: check if/why this is pref. instead of the default <- chan
	select {}
}

func main() {
	go devserver()
	httpserver()
}
