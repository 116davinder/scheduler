package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
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

func commandStart() {
	cmd := exec.Command("echo", "hola run")
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
}

func adhoc(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		cmd := exec.Command("echo", "sleep")
		stdout, err := cmd.Output()
		w.Header().Set("Content-Type", "application/json")

		if err != nil {
			errOutput := errOutput{outputJson{Result: string(err.Error())}}
			w.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(w).Encode(errOutput)
			if err != nil {
				json.NewEncoder(w).Encode(err.Error())
			}
		} else {
			resOutput := resOutput{outputJson{Result: string(stdout)}}
			err := json.NewEncoder(w).Encode(resOutput)
			if err != nil {
				json.NewEncoder(w).Encode(err.Error())
			}
		}
	case "POST":
		log.Print("Hello Post Request")
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
