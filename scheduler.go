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


	if err != nil {
		errOutput := errOutput{outputJson{Result: string(err.Error())}}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errOutput)
	}else {
		w.Header().Set("Content-Type", "application/json")
		resOutput := resOutput{outputJson{Result: string(stdout)}}
		json.NewEncoder(w).Encode(resOutput)

	}
}

func main() {
	router := &http.ServeMux{}
	router.HandleFunc("/adhoc", adhoc)
	http.ListenAndServe(":5000", router)
}
