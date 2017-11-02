package main

import (
	"net/http"
	"os/exec"
	"encoding/json"
)

type InputJson struct {
	Name  string `json:"name"`
	Query []byte `json:"query"`
}

type outputJson struct {
	Name  string `json:"name"`
	result string `json:"result"`
}

type errOutput []outputJson

func adhoc(w http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("echo1","sleep x; 121")
	stdout, err := cmd.Output()

	if err != nil {
		errOutput := errOutput{outputJson{Name: "Error",result: err.Error()}}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errOutput)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(stdout)))
}

func main() {
	router := &http.ServeMux{}
	router.HandleFunc("/adhoc", adhoc)
	http.ListenAndServe(":5000", router)
}
