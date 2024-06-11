package Controller

import (
	"html/template"
	"net/http"
)

func RenderHomePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" || req.Method != "GET" {
		Err404(res)
		return
	}

	temp, err := template.ParseFiles("./View/html/index.html")
	err = temp.Execute(res, nil)
	if err != nil {
		Err500(res)
		return
	}
}
