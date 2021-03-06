package chart

import (
	"bytes"
	"text/template"
)

const loadScript = `
{{- $tpls := index . "tpls" -}}
{{- $len := len $tpls -}} 
{{- if eq $len 0 }}
# no sources found in template
{{- else -}}
{{- range $path, $images := $tpls }}
##---
# Source: {{ $path }}
{{- $len = len $images -}} 
{{- if eq $len 0 }}
# no images found in source
{{- else -}}
{{- range $key, $image := $images }}
docker load -i ./{{ replace $image.Name ":" "_" -1 }}.tar
{{- end -}}
{{- end -}}
{{- end -}}
{{- end -}}
`

var loadTemplate = template.Must(template.New("").Funcs(templateFuncs).Parse(loadScript))

// GenerateLoadScript 產生 docker load 指令
func (t *Templates) GenerateLoadScript() ([]byte, error) {
	data := make(map[string]interface{})
	data["tpls"] = t
	var buf bytes.Buffer
	if err := loadTemplate.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
