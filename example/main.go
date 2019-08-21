package main

import (
	"fmt"
	"net/http"

	i18n "github.com/middangeard-fiction/rosetta-go"
)

func handler(w http.ResponseWriter, r *http.Request) {
	lang := i18n.DetectLanguage(r.Header.Get("Accept-Language"))
	fmt.Fprintf(w, "<html><head><title> %s </title></head><body>", i18n.Tr(lang, "hello_world"))
	fmt.Fprintf(w, "<h2> %s </h2>", i18n.Tr(lang, "hello_world"))
	githubLink := "https://github.com/bykovme/i18n"
	link := fmt.Sprintf(`<a href="%s">%s</a>`, githubLink, githubLink)
	fmt.Fprintf(w, i18n.Tr(lang, "find_more"), link)
	fmt.Fprint(w, "</body></html>")
}

func main() {
	err := i18n.InitLocales("langs")
	if err != nil {
		panic(err)
	}

	fmt.Println(i18n.GetUILanguage())

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
