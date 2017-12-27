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
	doc.Find(".menu-list-container form").Remove()
	doc.Find(".roundbox-lt").Remove()
	doc.Find(".roundbox-rt").Remove()
	doc.Find(".roundbox-lb").Remove()
	doc.Find(".roundbox-rb").Remove()
	doc.Find(".header-bell").Remove()
	doc.Find(".side-bell").Remove()

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
