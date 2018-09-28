package main

/*
	basic_main1~basic_main4均来自这个网站
	https://www.alexedwards.net/blog/a-recap-of-request-handling

	了解中间件前需要了解 ServeMux、DefaultServeMux、http.Handler、http.HandlerFunc、mux.HandleFunc、ServeHTTP 等相关知识和它们之间的关系
*/
import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	rh := http.RedirectHandler("http://example.org", 307)
	mux.Handle("/foo", rh)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}