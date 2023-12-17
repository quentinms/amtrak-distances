package main

import (
	"encoding/json"
	"html/template"
	"os"
)

type routeData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Source    string `json:"source"`
	Distances []struct {
		Station  string `json:"station"`
		Distance int    `json:"distance"`
	} `json:"distances"`
}

func main() {
	b, rfErr := os.ReadFile("routes.json")
	if rfErr != nil {
		panic(rfErr)
	}
	var routes []routeData
	json.Unmarshal(b, &routes)

	tmpl := template.Must(template.ParseFiles("index.html.tmpl"))
	tErr := tmpl.Execute(os.Stdout, routes)
	if tErr != nil {
		panic(tErr)
	}
}
