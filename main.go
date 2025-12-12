package main

import (
	asciiart "ascii-art-web/asciiart"
	"html/template"
	"log"
	"net/http"
)

type pageData struct {
	Result string
	Input  string
	Banner string
	Color  string
}

var bannerStr, resultStr, inputStr, ColorStr string

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Error : Invalid template", http.StatusNotFound)
		log.Print(r.Method, " ", http.StatusNotFound, " ", r.URL.Path)
		return
	}
	err = tmpl.Execute(w, pageData{Result: resultStr, Input: inputStr, Banner: bannerStr, Color: ColorStr})
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Print(r.Method, " ", http.StatusInternalServerError, " ", r.URL.Path)
		return
	}
	log.Print(r.Method, " ", 200, " ", r.URL.Path)
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Print(r.Form)
	text := r.Form["inputText"][0]
	banner := r.Form["banner"][0]
	color := r.Form["color"][0]
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		http.Error(w, "Error : Invalid banner", http.StatusBadRequest)
		log.Print(r.Method, " ", http.StatusBadRequest, " ", r.URL.Path)
		return
	}
	args := []string{text, banner}
	resultStr = asciiart.MainRender(args)
	inputStr = text
	bannerStr = banner
	ColorStr = color
	log.Print("\n" + resultStr)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	log.Print(r.Method, " ", 200, " ", r.URL.Path)
}

func main() {
	log.Print("starting server on : 8080")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/{$}", home)
	http.HandleFunc("POST /ascii-art", asciiHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
