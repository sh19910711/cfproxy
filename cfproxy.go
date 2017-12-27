package main

import (
	"github.com/elazarl/goproxy"
	"net/http"
	"os"
	"regexp"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	proxy.OnResponse(goproxy.DstHostIs("codeforces.com")).DoFunc(
		func(res *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			res.Body = Replace(res)
			return res
		})

	proxy.OnResponse(goproxy.UrlMatches(regexp.MustCompile("/submit$"))).DoFunc(
		func(res *http.Response, ctxt *goproxy.ProxyCtx) *http.Response {
			template := "./template"
			if _, err := os.Stat(template); err == nil {
				res.Body = ReplaceSourceCode(res, template)
			}
			return res
		})

	http.ListenAndServe("localhost:8181", proxy)
}
