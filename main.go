package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/html"
	"io/ioutil"
	"bytes"
)

const (
	pagename = "http://www.npp/"
)

func main() {
	fmt.Println("Start.")
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
	return
}


func Handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(pagename)
	if err != nil {
		fmt.Fprint(w, err)
	}
	rob, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	doc, err := html.Parse(bytes.NewReader(rob))
	par(doc, w)

	return
}


func par (n *html.Node, w http.ResponseWriter) {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Fprintln(w, a.Val)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		par(c, w)
	}
}
