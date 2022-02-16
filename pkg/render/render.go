package render

import (
	"bytes"
	"fmt"
	"github.com/thinkingmonster/hello-world/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc := app.TemplateCache

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Empty cache")
	}

	buff := new(bytes.Buffer)
	data := t.Execute(buff, nil)
	fmt.Println(data)

	_, err := buff.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

// CreateTemplateCache  creates a template cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
