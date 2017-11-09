package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/chilts/sid"
)

var uuid = sid.Id()

type inputJson struct {
	Module string `json:"module"`
	Args   string `json:"args"`
}

type outputJson struct {
	Result string `json:"result"`
}

type errOutput []outputJson
type resOutput []outputJson

func adhocGet(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hola")
}

func adhocPost(w http.ResponseWriter, req *http.Request) {
	reqInput := inputJson{}
 	json.NewDecoder(req.Body).Decode(&reqInput)
 	json.NewEncoder(w).Encode(&reqInput)
}


func adhoc(w http.ResponseWriter, req *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
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
	log.Println("Starting Web Server...")
	muxServer()
}
