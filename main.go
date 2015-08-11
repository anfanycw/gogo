package main

import (
    "log"
    "net/http"
)

var locations map[string]string

func main() {
	locations = map[string]string {
		"upstairs left": "VACANT",
		"upstairs right": "VACANT",
		"downstairs men": "VACANT",
		"downstairs women": "VACANT",
		"downstairs shower": "VACANT",
	}

    router := NewRouter()
    log.Fatal(http.ListenAndServe(":8080", router))
}