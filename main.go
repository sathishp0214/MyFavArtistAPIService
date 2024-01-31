package main

func init() {
	uploadTrackSampleValues()
}

func uploadTrackSampleValues() {
	track1 := Track{1, "Without You", "Without You track's lyrics", ArtistInformation{Name: "Harry Nilsson", ImageLink: "https://lastfm.freetls.fastly.net/i/u/34s/2a96cbd8b46e442fc41c2b86b821562f.png"}, "pop"}
	track2 := Track{2, "You’re So Vain", "You’re So Vain track's lyrics", ArtistInformation{Name: "Carly Simon", ImageLink: "https://lastfm.freetls.fastly.net/i/u/34s/2a96cbd8b46e442fc41c2b86b821562f.png"}, "pop"}
	track3 := Track{3, "Time After Time", "Time After Time track's lyrics", ArtistInformation{Name: "Cyndi Lauper", ImageLink: "https://lastfm.freetls.fastly.net/i/u/34s/2a96cbd8b46e442fc41c2b86b821562f.png"}, "pop"}
	track4 := Track{4, "Where Is My Mind?", "Where Is My Mind? track's lyrics", ArtistInformation{Name: "The Pixies", ImageLink: "https://lastfm.freetls.fastly.net/i/u/34s/2a96cbd8b46e442fc41c2b86b821562f.png"}, "melody"}
	track5 := Track{5, "So What", "So What track's lyrics", ArtistInformation{Name: "Cyndi Lauper", ImageLink: "https://lastfm.freetls.fastly.net/i/u/34s/2a96cbd8b46e442fc41c2b86b821562f.png"}, "melody"}
	track6 := Track{6, "Welcome to the Jungle", "Welcome to the Jungle track's lyrics", ArtistInformation{Name: "Guns N Roses", ImageLink: "https://lastfm.freetls.fastly.net/i/u/34s/2a96cbd8b46e442fc41c2b86b821562f.png"}, "beat"}
	track7 := Track{7, "Old Town Road", "Old Town Road track's lyrics", ArtistInformation{Name: "Lil Nas X", ImageLink: "https://lastfm.freetls.fastly.net/i/u/34s/2a96cbd8b46e442fc41c2b86b821562f.png"}, "beat"}
	track8 := Track{8, "Cannonball", "Cannonball track's lyrics", ArtistInformation{Name: "The Breeders", ImageLink: "https://lastfm.freetls.fastly.net/i/u/34s/2a96cbd8b46e442fc41c2b86b821562f.png"}, "Harry Nilsson"}
	track9 := Track{9, "House of Balloons", "House of Balloons track's lyrics", ArtistInformation{Name: "Miless Davis", ImageLink: "https://lastfm.freetls.fastly.net/i/u/34s/2a96cbd8b46e442fc41c2b86b821562f.png"}, "pop"}

	AllTracks = append(AllTracks, track1, track2, track3, track4, track5, track6, track7, track8, track9)

	RegionTopTrack["india"] = track3
	RegionTopTrack["italy"] = track5
	RegionTopTrack["usa"] = track4
	RegionTopTrack["russia"] = track3

}

func main() {
	//running the API server and listening
	ServerInitListening()
}
