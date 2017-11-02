package utils

import (
	"bytes"
	"text/template"
)

func Execute(tpl string, data *Context) (answer string, err error) {
	c := NewContext(data)
	t := template.Must(template.New("").Parse(tpl))
	var buf bytes.Buffer
	err = t.Execute(&buf, c)
	if err == nil {
		answer = buf.String()
	}
	return answer, nil
}
