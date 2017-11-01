package main

import (
	"net/http"
	"os/exec"
)

type queryJson struct {
	Bash  string `json:"bash"`
	Name  string `json:"name"`
	Query string `json:"query"`
}

func adhoc(w http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("echo1","sleep x; 121")
	stdout, err := cmd.Output()

	if err != nil {
		//println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Error: "+ err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(stdout)))
}

func main() {
	router := &http.ServeMux{}
	router.HandleFunc("/adhoc", adhoc)
	http.ListenAndServe(":5000", router)
}
