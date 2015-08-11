package main

type Occupancy struct {
    Vacant	[]string `json:"vacant"`
}

type Restroom struct {
	Status		string `json:"status"`
	Location	string `json:"location"`
}