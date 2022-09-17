package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/theclassicdev/monolithic-app/pkg/config"
	"github.com/theclassicdev/monolithic-app/pkg/models"
)

/* SIMPLER METHOD TO CREATE A CACHE FOR TEMPLATE*/
// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {

// 	var tmpl *template.Template
// 	var err error

// 	//check if we already have template in the template cache
// 	_, inMap := tc[t]
// 	if !inMap {
// 		//need to create the template
// 		log.Println("creating template and adding it to cache")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}

// 	} else {
// 		//we have the template in the cache
// 		log.Println("using cached template")
// 	}

// 	//assign the template data value to tmpl
// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}

// }

// func createTemplateCache(t string) error {
// 	//store values to be parsed into a slice
// 	templates := []string{fmt.Sprintf("./templates/%s", t), "./templates/base-layout.html"}

// 	//parse template
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	//add template to cache
// 	tc[t] = tmpl

// 	return nil
// }

/* RELATIVELY COMPLEX METHOD TO CREATE A CACHE FOR TEMPLATE*/
var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

//RenderTemplate render templates
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	//create template cache
	var tc map[string]*template.Template

	//if useCache is true build the template cache otherwise re  build the template cache
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	//buf is a variable that will hold bytes: finer grain error checking
	buf := new(bytes.Buffer)
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err) // this will tell you that the error logged out is an error from the map
	}

	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	//get all the files with name .html from from templates
	pages, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return myCache, err
	}

	//range through all the files that end with .html
	for _, page := range pages {
		fileName := filepath.Base(page)
		ts, err := template.New(fileName).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[fileName] = ts

	}

	return myCache, nil

}
