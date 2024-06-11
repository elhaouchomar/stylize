package Controller

import (
	"html/template"
	"net/http"
)

func RenderHomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" || r.Method != "GET" {
		Err404(w)
		return
	}

	temp, err := template.ParseFiles("./View/html/index.html")
	err = temp.Execute(w, nil)
	if err != nil {
		Err500(w)
		return
	}
}
