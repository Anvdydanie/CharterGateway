package handlers

import (
	"log"
	"net/http"
)

func FlightHandler(w http.ResponseWriter, req *http.Request) {
	var cityFrom = req.PostForm.Get("cityFrom")
	var cityTo = req.PostForm.Get("cityTo")
	var dateBack = req.PostForm.Get("dateBack")
	var dateTo = req.PostForm.Get("dateTo")
	if cityFrom == "" || cityTo == "" || dateBack == "" || dateTo == "" {
		log.Println("отсутствуют необходимые параметры")
		w.WriteHeader(400)
	} else {
		w.WriteHeader(200)
	}

}
