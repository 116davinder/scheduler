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
	Result string `json:"result"`
}

type errOutput []outputJson
type resOutput []outputJson

func adhoc(w http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("echo","sleep")
	stdout, err:= cmd.Output()
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		errOutput := errOutput{outputJson{Result: string(err.Error())}}
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(errOutput)
		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		}
	}else {
		resOutput := resOutput{outputJson{Result: string(stdout)}}
		err := json.NewEncoder(w).Encode(resOutput)
		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		}
	}
}

func main() {
	router := &http.ServeMux{}
	router.HandleFunc("/adhoc", adhoc)
	http.ListenAndServe(":5000", router)
}
