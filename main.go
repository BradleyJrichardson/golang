package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	sagesMap := map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhammad",
	}

	err = tpl.ExecuteTemplate(os.Stdout, "00-variable.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "01-range-slice.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "00-variable.gohtml", sagesMap)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "02-range-map.gohtml", sagesMap)
	if err != nil {
		log.Fatalln(err)
	}

}
