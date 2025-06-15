package main

import (
	"downloading-files/internal/downloader"
	"flag"
)

var opts downloader.Options

func init() {
	flag.StringVar(&(opts.Urls), "urls", "", "defines paths of the downloading")
	flag.StringVar(&(opts.Output), "output", "/downloads", "defines path where stores result")
}

func main() {
	flag.Parse()
	downloader.Run(opts.Urls, opts.Output)
}