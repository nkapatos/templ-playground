package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os/exec"
	"time"

	"github.com/a-h/templ"
	"github.com/evanw/esbuild/pkg/api"
	"github.com/nkapatos/templ-playground/esbuild-go/templ/views"
)

func httpserverv2() {
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

func tailwindCLI() (string, error) {
	cmd := exec.Command("npx", "@tailwindcss/cli", "-i", "assets/css/index.css", "-o", "-", "--minify")
	// cmd := exec.Command("npx", "@tailwindcss/cli", "-i", "./assets/input.css", "-o", "./src/output.css", "--watch")
	var out bytes.Buffer
	cmd.Stderr = &out
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func devserverv2(twcss string) error {

	ctx, err := api.Context(api.BuildOptions{
		EntryPoints: []string{"assets/ts/main.ts"},
		Stdin: &api.StdinOptions{
			Contents:   twcss,
			ResolveDir: "assets/css",
			Sourcefile: "input.css",
			Loader:     api.LoaderCSS,
		},
		Outdir:   "dist",
		Bundle:   true,
		LogLevel: api.LogLevelInfo,
	})

	if err != nil {
		return err
	}

	watch_error := ctx.Watch(api.WatchOptions{})

	if watch_error != nil {
		return watch_error
	}

	_, serve_error := ctx.Serve(api.ServeOptions{})
	if serve_error != nil {
		return serve_error
	}
	select {}
}

func mainv2() {
	twcssChan := make(chan struct {
		css string
		err error
	})
	go func() {
		css, err := tailwindCLI()
		twcssChan <- struct {
			css string
			err error
		}{css, err}

	}()

	select {
	case res := <-twcssChan:
		if res.err != nil {
			log.Fatalf("Tailwind build failed: %v", res.err)
		}
		if err := devserverv2(res.css); err != nil {
			log.Fatalf("esbuild failed: %v", err)
		}
		log.Println("Build completed successfully")
	case <-time.After(10 * time.Second):
		log.Fatal("Timeout waiting for Tailwind build")

	}

	// go devserver()
	httpserverv2()
}
