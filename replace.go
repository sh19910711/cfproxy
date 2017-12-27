package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net/http"
)

func Replace(res *http.Response) io.ReadCloser {
	doc, _ := goquery.NewDocumentFromResponse(res)

	doc.Find("div#sidebar").Remove()

	doc.Find(".second-level-menu-list").Children().Each(func(_ int, el *goquery.Selection) {
		html, _ := el.Html()
		doc.Find("body").PrependHtml(html + " / ")
	}).Remove()

	doc.Find(".main-menu-list").Children().Each(func(_ int, el *goquery.Selection) {
		html, _ := el.Html()
		doc.Find("body").PrependHtml(html + " / ")
	}).Remove()

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
