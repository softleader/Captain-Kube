package chart

import (
	"bytes"
	"text/template"
)

const retagScript = `
{{- $from := index . "from" -}}
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
docker tag {{ $from }}/{{ $image.Name }} {{ $image.Host }}/{{ $image.Name }} && docker push {{ $image.Host }}/{{ $image.Name }}
{{- end -}}
{{- end -}}
{{- end -}}
{{- end -}}
`

var retagTemplate = template.Must(template.New("").Parse(retagScript))

// GenerateReTagScript 產生 docker tag 及 docker push 指令
func (t *Templates) GenerateReTagScript(from, to string) ([]byte, error) {
	retags := make(map[string][]*Image)
	for src, images := range *t {
		for _, image := range images {
			if image.Host == from {
				image.ReTag(from, to)
				retags[src] = append(retags[src], image)
			}
		}
	}
	data := make(map[string]interface{})
	data["from"] = from
	data["tpls"] = retags
	var buf bytes.Buffer
	if err := retagTemplate.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
