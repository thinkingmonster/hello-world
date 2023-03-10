package render

import (
	"bytes"
	"html/template"
	"log"
	"myapp/pkg/config"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// NewTemplate sets the config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template
	// Create a template cache
	//tc, err := CreateTemplateCache()
	//if err != nil {
	//	log.Fatal(err)
	//}
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template")
	}
	// render the template this is one way of doing the thing, we can use a  buffer for more fine-grained  control
	//err = t.Execute(w, nil)
	//if err != nil {
	//	fmt.Println("Error parsing template", err)
	//	return
	//}
	buf := new(bytes.Buffer)
	err := t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}
	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all the files names *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	// range through all the files ending with *.page.tmpl

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
		}
		if err != nil {
			return myCache, err
		}
		myCache[name] = ts
	}
	return myCache, nil
}

//var tc = make(map[string]*template.Template)
//
//func RenderTemplate(w http.ResponseWriter, t string) {
//	var tmpl *template.Template
//	var err error
//
//	_, inMap := tc[t]
//	if !inMap {
//		err = createTemplateCache(t)
//		if err != nil {
//			log.Println(err)
//		}
//	} else {
//		// we have template in the cache and we will use it
//		log.Println("Using cached template")
//	}
//	tmpl = tc[t]
//	err = tmpl.Execute(w, nil)
//	if err != nil {
//		log.Println(err)
//	}
//}

//func createTemplateCache(t string) error {
//	templates := []string{
//		fmt.Sprintf("./templates/%s", t),
//		"./templates/base.layout.tmpl",
//	}
//	//parse the template
//	tmpl, err := template.ParseFiles(templates...)
//	if err != nil {
//		return err
//	}
//	tc[t] = tmpl
//	return nil
//}
