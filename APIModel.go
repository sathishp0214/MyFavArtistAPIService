package main

type Track struct {
	Id         int               `json:"id"`
	Name       string            `json:"name"`
	Lyrics     string            `json:"lyrics"`
	ArtistInfo ArtistInformation `json:"artistinfo"`
	Tag        string            `json:"tag"`
}

type ArtistInformation struct {
	Name      string `json:"name"`
	ImageLink string `json:"imagelink"`
}
