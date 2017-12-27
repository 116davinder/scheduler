package main

type acceptJSON struct {
	Name  string `json:"name"`
	Query string `json:"query"`
}

type queryjson []acceptJSON
