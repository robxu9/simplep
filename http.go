package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func makeHTTPTunnel(listen, target string) {
	targetParse, err := url.Parse(target)
	if err != nil {
		errAndExit(err)
	}

	reverseProxy := httputil.NewSingleHostReverseProxy(targetParse)
	go http.ListenAndServe(listen, reverseProxy)
}
