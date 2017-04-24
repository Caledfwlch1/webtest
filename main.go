package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/html"
	"io/ioutil"
	"bytes"
	"html/template"
)

const (
	pagename = "http://golang-book.ru"
)

type Page struct {
	Link	string
}

func main() {
	fmt.Println("Start.")
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
	return
}


func Handler(w http.ResponseWriter, r *http.Request) {
	Linkpage := pagename
	p := &Page{Link: Linkpage}
	renderTemplate(w, p)

	resp, err := http.Get(Linkpage)
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




func renderTemplate(w http.ResponseWriter, p *Page) {
	var err error
	TemplateRep := template.Must(template.ParseFiles("index.html")) // шаблон для первой страницы и страницы с результатом поиска

	err = TemplateRep.ExecuteTemplate(w, "index.html", p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}
