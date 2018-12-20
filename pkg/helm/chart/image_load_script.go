package chart

import (
	"bytes"
	"text/template"
)

const loadScript = `
{{- range $path, $images := index . "tpls" }}
##---
# Source: {{ $path }}
{{- range $key, $image := $images }}
docker load -i ./{{ $image.Name }}.tar
{{- end }}
{{- end }}
`

var loadTemplate = template.Must(template.New("").Parse(loadScript))

func (i *Templates) GenerateLoadScript() (buf bytes.Buffer, err error) {
	data := make(map[string]interface{})
	data["tpls"] = i
	err = loadTemplate.Execute(&buf, data)
	return
}
