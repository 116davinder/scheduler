package main

import (
	"net/http"
	"os/exec"
)

func adhoc(w http.ResponseWriter, req *http.Request) {
	// app := "echo"
	// arg1 := "Running My first Commandss in Golang"

	cmd := exec.Command("echo", "HOla")
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	println(string(stdout))
}

func main() {
	router := &http.ServeMux{}
	router.HandleFunc("/adhoc", adhoc)
	http.ListenAndServe(":80", router)
}
