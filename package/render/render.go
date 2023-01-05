package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/sanyog2491/bookings/package/config"
	"github.com/sanyog2491/bookings/package/model"
)

var functions = template.FuncMap{}

var app *config.Appconfig

func NewTemplate(a *config.Appconfig) {
	app = a

}
func AddDefaultData(td *model.TemplateData) *model.TemplateData {

	return td
}
func Rendertemplate(w http.ResponseWriter, tmpl string, td *model.TemplateData) {

	//template data holds the data sent from handlers to templates

	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = Createtemplatecache()
	}

	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println(err)
	}

}

func Createtemplatecache() (map[string]*template.Template, error) {

	mycache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return mycache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return mycache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return mycache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return mycache, err
			}
		}
		mycache[name] = ts

	}
	return mycache, nil

}
