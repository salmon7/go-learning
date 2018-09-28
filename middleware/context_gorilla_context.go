package main

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
)

type Key string

const GlobalRequestVariable Key =""

func SetGlobalHandler(w http.ResponseWriter, r *http.Request) {
	context.Set(r, GlobalRequestVariable, "test")

	get, ok := context.GetOk(r, GlobalRequestVariable)
	w.Write([]byte(fmt.Sprintf("GetOK : [%v] and get what :[%v] ", ok, get)))

	InternalGetGlobalHandler(w, r)

}

func InternalGetGlobalHandler(w http.ResponseWriter, r *http.Request) {

	get, ok := context.GetOk(r, GlobalRequestVariable)
	w.Write([]byte(fmt.Sprintf("\nInternal GetOK : [%v] and get what :[%v] ", ok, get)))

}

func GetGlobalHandler(w http.ResponseWriter, r *http.Request) {

	get, ok := context.GetOk(r, GlobalRequestVariable)
	w.Write([]byte(fmt.Sprintf("GetOK : [%v] and get what :[%v] ", ok, get)))

}

func main() {
	mx := mux.NewRouter()

	mx.HandleFunc("/", SetGlobalHandler)
	mx.HandleFunc("/get", GetGlobalHandler)

	http.ListenAndServe(":8080", mx)
}
/*
form https://blog.csdn.net/wangshubo1989/article/details/78910935
 */
