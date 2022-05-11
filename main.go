package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/Masterminds/semver"
)

type Release struct {
	Version      string   `json:"version"`
	ReleaseDate  string   `json:"releaseDate"`
	ReleaseNotes string   `json:"releaseNotes"`
	TKR          []string `json:"tkr"`
}

type Product struct {
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	ShortName string    `json:"shortName"`
	Docs      string    `json:"docs"`
	Releases  []Release `json:"releases"`
}

type TemplateData struct {
	Products           []Product
	KubernetesReleases map[string]map[string]string
}

func GetAllTKRReleases(products []Product) map[string]map[string]string {
	results := map[string]map[string]string{}
	for _, product := range products {
		for _, release := range product.Releases {
			for _, tkr := range release.TKR {
				k8sVersionObject, err := semver.NewVersion(tkr)
				if err != nil {
					_, _ = fmt.Fprintf(os.Stderr, "TKR version \"%s\" is not a valid semver. Found in %s %s", tkr, product.Slug, release.Version)
				} else {
					k8sVersion := fmt.Sprintf("%d.%d", k8sVersionObject.Major(), k8sVersionObject.Minor())
					if results[k8sVersion] == nil {
						results[k8sVersion] = map[string]string{}
					}
					tkgVersion := fmt.Sprintf("%s %s", product.ShortName, release.Version)
					results[k8sVersion][tkgVersion] = tkgVersion
				}
			}
		}
	}

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
		Products:           products,
		KubernetesReleases: GetAllTKRReleases(products),
	})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to execute template: %s\n", err.Error())
		return
	}
}
