package template

import (
	"github.com/gin-contrib/multitemplate"
	"html/template"
	"io/fs"
	"path"
)


type Config struct {
	Name string
	Files []string
}

var funcMap = map[string]interface{}{}

func AddFunc(name string, i interface{}) {
	funcMap[name] = i
}

func NewTemplateRender(fs fs.FS, tcs []Config) (multitemplate.Renderer, error) {
	r := multitemplate.NewRenderer()
	for _, tc := range tcs {
		// template name see template.ParseFiles
		tName := path.Base(tc.Files[0])
		funcMap["unescaped"] = unescaped
		t, err := template.New(tName).Funcs(funcMap).ParseFS(fs, tc.Files...)
		if err != nil {
			return nil, err
		}
		r.Add(tc.Name, t)
	}

	return r, nil
}
