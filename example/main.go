package main

import (
	"fmt"
	"net/http"

	i18n "github.com/middangeard-fiction/rosetta-go"
)

func handler(w http.ResponseWriter, r *http.Request) {
	err := i18n.Init("langs", r.Header.Get("Accept-Language"))
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "<html><head><title> %s </title></head><body>", i18n.GetMessage("hello_world"))
	fmt.Fprintf(w, "<h2> %s </h2>", i18n.GetMessage("hello_world"))
	githubLink := "https://github.com/bykovme/i18n"
	link := fmt.Sprintf(`<a href="%s">%s</a>`, githubLink, githubLink)
	fmt.Fprintf(w, i18n.GetMessage("find_more"), link)
	fmt.Fprint(w, "</body></html>")
}

func main() {
	fmt.Println(i18n.GetUILanguage())

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
