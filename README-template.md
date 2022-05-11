# TKG Releases

## Products
{{range .Products}}
### {{.Name}} ({{.ShortName}})
Docs: {{.Docs}}

|Version|TKR|Date|Notes|
|-------|---|----|-----|
{{- range .Releases}}
|{{.Version}}|{{range .TKR}}{{.}}<br />{{end}}|{{.ReleaseDate}}|[Release notes]({{.ReleaseNotes}})|
{{- end}}
{{end}}

## By Kubernetes Release
|Kubernetes Release|TKG Versions|
|------------------|------------|
{{- range $kubernetesRelease, $tkgReleases := .KubernetesReleases}}
|{{ $kubernetesRelease }}|{{range $tkgReleases }}{{.}}<br />{{end}}|
{{- end}}
