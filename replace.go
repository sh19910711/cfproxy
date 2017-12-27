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
