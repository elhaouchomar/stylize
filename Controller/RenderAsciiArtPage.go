package Controller

import (
	"ascii-art-web/model"
	"html/template"
	"net/http"
)

type HomePageData struct {
	Message string
}

var (
	Err500 = func(w http.ResponseWriter) {
		HandleError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	Err404 = func(w http.ResponseWriter) {
		HandleError(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	Err400 = func(w http.ResponseWriter) {
		HandleError(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	Err405 = func(w http.ResponseWriter) {
		HandleError(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
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

	if err := ValidateInputAndHandleErrors(w, text); err != nil {
		return
	}

	if !validStyles[style] {
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

	data := HomePageData{
		Message: finalArt,
	}
	temp.Execute(w, data)
}
