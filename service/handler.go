package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func OccupancyGetHandler(w http.ResponseWriter, r *http.Request) {
	var vacancy []string
	for key, room := range locations {
		if room.Status == Vacant {
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

	loc, ok := locations[room.Location]
	if ok {
		loc.Status = room.Status
		loc.Name = ""
	}

	log.Printf("%s : %+v", room.Location, *loc)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

func ReservationGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	location := vars["location"]

	loc, ok := locations[location]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(loc); err != nil {
		panic(err)
	}
}

func ReservationPostHandler(w http.ResponseWriter, r *http.Request) {
	var res Reservation

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&res)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	loc, ok := locations[res.Location]
	if !ok {
		// location does not exist
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if loc.Status != Vacant {
		// location is already occupied or reserved
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	loc.Status = Reserved
	loc.Name = res.Name

	log.Printf("%s : %+v", res.Location, *loc)

	w.WriteHeader(http.StatusCreated)
}
