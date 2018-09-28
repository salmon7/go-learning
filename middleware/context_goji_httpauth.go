package main

import (
	"net/http"

	"github.com/goji/httpauth"
)

func main() {
	finalHandler := http.HandlerFunc(final)
	authHandler := httpauth.SimpleBasicAuth("username", "password")

	http.Handle("/", authHandler(finalHandler))
	http.ListenAndServe(":8080", nil)
}

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

/*
curl -i username:password@localhost:8080
HTTP/1.1 200 OK
Date: Thu, 01 Feb 2018 05:27:33 GMT
Content-Length: 2
Content-Type: text/plain; charset=utf-8

OK

from  https://blog.csdn.net/wangshubo1989/article/details/79227443
*/