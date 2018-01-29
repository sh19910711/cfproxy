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

	doc.Find("script").Remove()
	doc.Find("style").Remove()
	doc.Find("link").Remove()
	doc.Find("#sidebar").Remove()
	doc.Find(".menu-list-container form").Remove()
	doc.Find(".roundbox-lt").Remove()
	doc.Find(".roundbox-rt").Remove()
	doc.Find(".roundbox-lb").Remove()
	doc.Find(".roundbox-rb").Remove()
	doc.Find(".header-bell").Remove()
	doc.Find(".side-bell").Remove()

	doc.Find("body").PrependHtml("<div id=\"cfproxy-header\"></div>")
	header := doc.Find("#cfproxy-header")
	doc.Find(".main-menu-list").Children().Each(func(_ int, el *goquery.Selection) {
		html, _ := el.Html()
		header.AppendHtml(html + " / ")
	}).Remove()
	header.AppendHtml("<hr>")
	doc.Find(".second-level-menu-list").Children().Each(func(_ int, el *goquery.Selection) {
		html, _ := el.Html()
		header.AppendHtml(html + " / ")
	}).Remove()

	r, _ := doc.Html()
	return ioutil.NopCloser(bytes.NewBufferString(r))
}

func ReplaceSourceCode(res *http.Response, path string) io.ReadCloser {
	doc, _ := goquery.NewDocumentFromResponse(res)

	content, _ := doc.Find("#pageContent form").Html()
	doc.Find("body").Remove()
	doc.AppendHtml(content)

	code, _ := ioutil.ReadFile(path)
	doc.Find("#sourceCodeTextarea").SetText(string(code))

	r, _ := doc.Html()
	return ioutil.NopCloser(bytes.NewBufferString(r))
}
