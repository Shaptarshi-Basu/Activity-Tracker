package main

import (
	actHandler "activity-tracker"
	"log"
	"net/http"
)

type server struct{
	handler actHandler.ActivityHandler
}


func main() {
	mux := http.NewServeMux()
	s := &server{
		handler:actHandler.ActivityHandler{},
	}
	mux.HandleFunc("/add/activity", s.handler.AddActivity)
	mux.HandleFunc("/get/activity", s.handler.GetActivity)
	err := http.ListenAndServe(":7070", mux)
	if err != nil {
		log.Print(err)
	}
}
