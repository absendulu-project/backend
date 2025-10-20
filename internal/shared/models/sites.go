package models

type Site struct {
	Name      string   `json:"name"`
	Radius    int      `json:"radius"`
	Address   string   `json:"address"`
	Latitude  float64  `json:"latitude"`
	Longitude float64  `json:"longitude"`
	Companies *Company `json:"companies,omitempty"`
}
