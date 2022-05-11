package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

type Release struct {
	Version      string   `json:"version"`
	ReleaseDate  string   `json:"releaseDate"`
	ReleaseNotes string   `json:"releaseNotes"`
	TKR          []string `json:"tkr"`
}

type Product struct {
	Slug     string    `json:"slug"`
	Name     string    `json:"name"`
	Docs     string    `json:"docs"`
	Releases []Release `json:"releases"`
}

type TemplateData struct {
	Products []Product
}

func GetAllTKRReleases(products []Product) map[string][]string {
	results := map[string][]string{}

	// For each product
	//	 for each version
	//     for each tkr
	//       get major.minor as key
	//       results[key] = "product slug + product version"

	return results
}

func main() {
	data, err := ioutil.ReadFile("releases.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to read releases file: %s\n", err.Error())
		return
	}

	var products []Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to parse releases file: %s\n", err.Error())
		return
	}

	readmeTemplate, err := template.ParseFiles("README-template.md")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to parse template file: %s\n", err.Error())
		return
	}

	err = readmeTemplate.Execute(os.Stdout, &TemplateData{
		Products: products,
	})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to execute template: %s\n", err.Error())
		return
	}
}
