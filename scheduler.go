package main

import (
	"fmt"
	"log"
	"net/http"
)

type inputJson struct {
	Module string `json:"module"`
	Args   string `json:"args"`
}

type outputJson struct {
	Result string `json:"result"`
}

type errOutput []outputJson
type resOutput []outputJson
type inputs []inputJson

func adhocGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola adhocGet Function")
}

func adhocPost(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hola adhocPost Function")
}

func adhoc(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		adhocGet(w, req)
	case "POST":
		adhocPost(w, req)
	}
}

func muxServer() {
	router := &http.ServeMux{}
	router.HandleFunc("/adhoc", adhoc)
	log.Fatal(http.ListenAndServe(":5000", router))
}

func main() {
	muxServer()
}
