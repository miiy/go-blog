package template

import "html/template"

func unescaped(x string) interface{} {
	return template.HTML(x)
}
