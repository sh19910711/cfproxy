package main

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net/http"
)

func Replace(res *http.Response) io.ReadCloser {
	doc, _ := goquery.NewDocumentFromResponse(res)

	doc.Find("div#sidebar").Remove()

	r, _ := doc.Html()
	return ioutil.NopCloser(bytes.NewBufferString(r))
}

func ReplaceSourceCode(res *http.Response, path string) io.ReadCloser {
	doc, _ := goquery.NewDocumentFromResponse(res)

	code, _ := ioutil.ReadFile(path)
	doc.Find("#sourceCodeTextarea").SetText(string(code))

	r, _ := doc.Html()
	return ioutil.NopCloser(bytes.NewBufferString(r))
}
