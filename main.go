package main

import (
	"fmt"
	"net/http"
)

const (
	pagename = "http://www.npp.zp.ua"
)

func main() {
	fmt.Println("Start.")
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":80", nil)
	fmt.Println(err)
	return
}


func Handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(pagename)
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, resp)
	return
}