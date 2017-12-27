package main

import (
	"flag"
	"github.com/elazarl/goproxy"
	"net/http"
	"os"
	"regexp"
)

func main() {
	templatePath := *flag.String("template", "template.cpp", "path/to/template.cpp")
	flag.Parse()

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	proxy.OnResponse(goproxy.DstHostIs("codeforces.com")).DoFunc(
		func(res *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			res.Body = Replace(res)
			return res
		})

	proxy.OnResponse(goproxy.UrlMatches(regexp.MustCompile("/submit$"))).DoFunc(
		func(res *http.Response, ctxt *goproxy.ProxyCtx) *http.Response {
			if _, err := os.Stat(templatePath); err == nil {
				res.Body = ReplaceSourceCode(res, templatePath)
			}
			return res
		})

	http.ListenAndServe("localhost:8181", proxy)
}
