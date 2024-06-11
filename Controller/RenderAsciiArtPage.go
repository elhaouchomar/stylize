package Controller

import (
	"ascii-art-web/model"
	"html/template"
	"net/http"
)

var (
	Err400 = func(w http.ResponseWriter) {
		HandleError(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	Err404 = func(w http.ResponseWriter) {
		HandleError(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	Err405 = func(w http.ResponseWriter) {
		HandleError(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	Err500 = func(w http.ResponseWriter) {
		HandleError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
)

func RenderAsciiArtPage(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	style := r.FormValue("style")
	validStyles := map[string]bool{
		"shadow":     true,
		"standard":   true,
		"thinkertoy": true,
	}

	if text == "" || !validStyles[style] {
		Err400(w)
		return
	}

	if r.Method != "POST" {
		Err405(w)
		return
	}

	temp, err1 := template.ParseFiles("./View/html/asciiArt.html")
	if err1 != nil {
		Err500(w)
		return
	}

	finalArt, err2 := model.GenerateAsciiArt(text, style)
	if err2 != nil {
		Err500(w)
		return
	}

	if err := ValidateInputAndHandleErrors(w, text); err != nil {
		return
	}
	data := HomePageData{
		Message: finalArt,
	}
	temp.Execute(w, data)
}
