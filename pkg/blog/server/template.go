package server

import (
	"log"
	"net/http"
	"text/template"
)

func templateHandle(file string, w http.ResponseWriter, data interface{}) {
	tpl, err := template.ParseFiles(file)
	if err != nil {
		log.Printf("parse template file failed: %+v", err)
		return
	}
	tpl.Execute(w, data)
}
