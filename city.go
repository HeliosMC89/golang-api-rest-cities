package main

type City struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"country"`
	District    string `json:"district"`
	Population  int    `json:"population"`
}

type Cities []City
