package main

import (
	"github.com/evanw/esbuild/pkg/api"
)

func main() {
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

	select {}

}
