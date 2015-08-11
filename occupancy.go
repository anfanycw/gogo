package main

type Occupancy struct {
    Locations	[]string `json:[]string`
}

type Occupancies []Occupancy

type Status struct {
	Value	string `json:"status"`
}