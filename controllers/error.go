package controllers

import (
	"html/template"
	"net/http"
)

var (
	Err500 = func(w http.ResponseWriter) {
		HandleError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	Err404 = func(w http.ResponseWriter) {
		HandleError(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	Err405 = func(w http.ResponseWriter) {
		HandleError(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	Err400 = func(w http.ResponseWriter) {
		HandleError(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	errorTemplate = template.Must(template.ParseFiles("views/errors.html"))
)

type ErrorData struct {
	ErrStatusInt    int
	ErrStatusString string
}

func HandleError(w http.ResponseWriter, statusString string, statusCode int) {
	w.WriteHeader(statusCode)
	data := ErrorData{
		ErrStatusInt:    statusCode,
		ErrStatusString: statusString,
	}
	if err := errorTemplate.ExecuteTemplate(w, "errors.html", data); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
