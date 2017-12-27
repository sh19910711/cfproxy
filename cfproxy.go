package main

import (
	"github.com/elazarl/goproxy"
	"net/http"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	proxy.OnResponse(goproxy.DstHostIs("codeforces.com")).DoFunc(
		func(res *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			res.Body = Replace(res)
			return res
		})

	http.ListenAndServe("localhost:8181", proxy)
}
