package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/wisnercelucus/udemy-booking/pkg/config"
	"github.com/wisnercelucus/udemy-booking/pkg/models"
)

func addDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

var functions = template.FuncMap{}

/*func RenderTemplate(w http.ResponseWriter, path string, templ string) {
	parsedTemplate, _ := template.ParseFiles(path + templ)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("An error occured while parsing template", err)
	}
}
*/

// RenderTemplate renders a template
var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()

	}

	t, ok := tc[tmpl]

	if !ok {
		log.Fatalln("Could not get template from cache")
	}

	buf := new(bytes.Buffer)

	td = addDefaultData(td)
	_ = t.Execute(buf, td)

	buf.WriteTo(w)
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
