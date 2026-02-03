package models

type Station struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	City   string `json:"city"`
	Region string `json:"region"`
	Status string `json:"status"`
}
