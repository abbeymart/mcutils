package utils

import (
	"html/template"
	"time"
)

var funcMap = template.FuncMap{
	"dateFormat": dateFormat,
}

func dateFormat(layout string, d time.Time) string {
	return d.Format(layout)
}

func ParseTemplateString(name, tpl string) *template.Template {
	t := template.New(name)
	t.Funcs(funcMap)
	t = template.Must(t.Parse(tpl))
	return t
}

func ParseTemplateFile(name, file string) *template.Template {
	t := template.New(name)
	t.Funcs(funcMap)
	//t = template.Must(t.ParseFiles(tpl))
	t = template.Must(template.ParseFiles("templates/" + file))
	return t
}
