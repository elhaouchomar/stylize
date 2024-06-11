package Controller

import (
	"errors"
	"html/template"
	"net/http"
	"strings"
)

type ErrorResponseData struct {
	ErrStatusInt    int
	ErrStatusString string
}
type HomePageData struct {
	Message string
}

func ValidateInputAndHandleErrors(res http.ResponseWriter, text string) error {
	arr := strings.Fields(text)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] > 126 || arr[i][j] < 32 {
				Err400(res)
				return errors.New("error")
			}
		}
	}
	return nil
}

func HandleError(res http.ResponseWriter, Errortxt string, Errorstatus int) {
	temp, _ := template.ParseFiles("./View/html/error.html")
	data := ErrorResponseData{
		ErrStatusInt:    Errorstatus,
		ErrStatusString: Errortxt,
	}
	res.WriteHeader(Errorstatus)
	temp.Execute(res, data)
}
