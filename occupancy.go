package main

type Occupancy struct {
    Locations	[]string `json:[]string`
}

type Occupancies []Occupancy

type Stat struct {
	Stat	string `json:"string"`
}

type Status []Stat