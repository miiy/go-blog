package template

import (
	"github.com/gin-contrib/multitemplate"
	"goblog.com/web/resources/templates"
	"html/template"
	"path"
)

func NewTemplateRender(templatesMap map[string][]string) (multitemplate.Renderer, error) {
	r := multitemplate.NewRenderer()
	for name, tps := range templatesMap {
		// template name see template.ParseFiles
		tName := path.Base(tps[0])
		t, err := template.New(tName).Funcs(template.FuncMap{
			"unescaped": unescaped,
		}).ParseFS(templates.FS,tps...)
		if err != nil {
			return nil, err
		}
		r.Add(name, t)
	}

	return r, nil
}
