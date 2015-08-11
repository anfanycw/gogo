package main

import (
	"log"
	"net/http"
)

var locations map[string]*Restroom

func main() {
	locations = map[string]*Restroom{
		"upstairs left":     {Status: Vacant},
		"upstairs right":    {Status: Vacant},
		"downstairs men":    {Status: Vacant},
		"downstairs women":  {Status: Vacant},
		"downstairs shower": {Status: Vacant},
	}

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
