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

## TODO

* Add TKR section that shows k8s releases mapped to tanzu releases
  * Just probably just high-level:
  * 1.21 --> TKGM 1.4, 1.5
  * 1.22 --> TKGM 1.5, TKGI 1.13