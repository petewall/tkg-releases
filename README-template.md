# TKG Releases

## Products
{{range .Products}}
### {{.Name}}
Docs: {{.Docs}}

|Version|TKR|Date|Notes|
|-------|---|----|-----|
{{- range .Releases}}
|{{.Version}}|
    {{- range .TKR}}{{.}}<br />{{end -}}
|{{.ReleaseDate}}|[Release notes]({{.ReleaseNotes}})|
{{- end}}
{{end}}
