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
type resOutput []outputJson

func adhoc(w http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("echo1","sleep")
	stdout, err := cmd.Output()

	if err != nil {
		errOutput := errOutput{outputJson{Name: "Error",result: "hola"}}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errOutput)
	}else {
		w.Header().Set("Content-Type", "application/json")
		resOutput := resOutput{outputJson{Name: "Output", result: string(stdout)}}
		json.NewEncoder(w).Encode(resOutput)
		//w.Write([]byte(string(stdout)))
	}

}

func main() {
	router := &http.ServeMux{}
	router.HandleFunc("/adhoc", adhoc)
	http.ListenAndServe(":5000", router)
}
