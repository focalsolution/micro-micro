package template

var (
	Plugin = `package main
{{if .Plugins}}
import ({{range .Plugins}}
	_ "github.com/focalsolution/micro-go-plugins/{{.}}"{{end}}
){{end}}
`
)
