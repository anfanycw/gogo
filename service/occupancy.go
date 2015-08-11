package main

const (
	Reserved = "RESERVED"
	Occupied = "OCCUPIED"
	Vacant   = "VACANT"
)

type Occupancy struct {
	Vacant []string `json:"vacant"`
}

type Restroom struct {
	Status   string `json:"status"` // RESERVED | OCCUPIED | VACANT
	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
}

type Reservation struct {
	Location string `json:"location"`
	Name     string `json:"name"`
}
