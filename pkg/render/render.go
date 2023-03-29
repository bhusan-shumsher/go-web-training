package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// func RenderTemplate(w http.ResponseWriter, tmpl string) {
// 	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("Error parsing template", err)
// 	}
// }

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	_, inMap := tc[t]
	if !inMap {
		// not in cache
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Print("Serving from template cache")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templates...) // take each entry from templates slice
	if err != nil {
		return err
	}
	tc[t] = tmpl
	return nil
}
