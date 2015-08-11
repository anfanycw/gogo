package main

import (
    "encoding/json"
    "log"
    "net/http"

    //"github.com/gorilla/mux"
)

var locations []string

func OccupancyGetHandler(w http.ResponseWriter, r *http.Request) {
    locations := []string{"here", "here too"}
    occ := Occupancies {
        Occupancy{Locations: locations},
    }
    
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(occ); err != nil {
        panic(err)
    }
}

func OccupancyPostHandler(w http.ResponseWriter, r *http.Request) {
    var status Status
    
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&status)
    if err != nil {
        panic(err)
    }
    log.Println(status.Value)

    // w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    // w.WriteHeader(http.StatusCreated)
    // if err := json.NewEncoder(w).Encode(t); err != nil {
    //     panic(err)
    // }
}