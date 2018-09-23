package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var WorkQueue = make(chan CSPRequest, 100)

func Collector(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/" {
		message := fmt.Sprintf("Path \"%s\" not found.", r.RequestURI)
		response(w, http.StatusNotFound, message)
		return
	}

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		message := fmt.Sprintf("Method \"%s\" not allowed.", r.Method)
		response(w, http.StatusMethodNotAllowed, message)
		return
	}

	contentType := r.Header.Get("Content-Type")
	if contentType != "application/csp-report" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		message := fmt.Sprintf("Unsupported Media Type \"%s\".", contentType)
		response(w, http.StatusUnsupportedMediaType, message)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		message := err.Error()
		response(w, http.StatusInternalServerError, message)
		return
	}

	data := NewCSPRequest()
	err = json.Unmarshal(body, &data)
	if err != nil {
		message := err.Error()
		response(w, http.StatusBadRequest, message)
		return
	}

	WorkQueue <- data

	message := "Thanks for reporting."
	response(w, http.StatusCreated, message)
	return
}

func response(w http.ResponseWriter, status int, message string) {
	log.Print(message)

	d := map[string]string{
		"message": message,
	}

	j, _ := json.MarshalIndent(d, "", "    ")

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	w.Write(j)
}
