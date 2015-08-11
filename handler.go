package main

import (
    "encoding/json"
    "log"
    "net/http"

    //"github.com/gorilla/mux"
)

func OccupancyGetHandler(w http.ResponseWriter, r *http.Request) {
    var vacancy []string
    for key, value := range locations {
        if value == "VACANT" {
            vacancy = append(vacancy, key)
        }
    }

    result := &Occupancy{
        Vacant: vacancy,
    }
    
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(result); err != nil {
        panic(err)
    }
}

func OccupancyPostHandler(w http.ResponseWriter, r *http.Request) {
    var room Restroom
    
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&room)
    if err != nil {
        panic(err)
    }

    _, ok := locations[room.Location]
    if ok {
        locations[room.Location] = room.Status
    }
    log.Printf("%+v", locations)

    // w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    // w.WriteHeader(http.StatusCreated)
    // if err := json.NewEncoder(w).Encode(t); err != nil {
    //     panic(err)
    // }
}