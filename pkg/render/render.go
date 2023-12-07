package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/bongochat/wannachat/pkg/config"
	"github.com/bongochat/wannachat/pkg/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig

// set the configuration for the template
func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	intMap := make(map[string]int)
	intMap["current_year"] = time.Now().Year()
	td.IntMap = intMap

	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	var err error

	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		// create new template cache
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Println(err)
		}
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from the cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	err = t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	template_cache := map[string]*template.Template{}

	// get all of the files named *.page.go.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.go.tmpl")
	if err != nil {
		return template_cache, err
	}

	// range through all of the files ending with *.page.go.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return template_cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.go.tmpl")
		if err != nil {
			return template_cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*layout.go.tmpl")
			if err != nil {
				return template_cache, err
			}
		}

		template_cache[name] = ts
	}

	return template_cache, nil
}
