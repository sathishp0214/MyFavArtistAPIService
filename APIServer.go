package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"github.com/go-chi/chi"
)

const ListeningPort = ":10000"

var (
	InputRegionTrack = "India"
	ApiUrl           = "http://localhost" + ListeningPort + "/region/" + InputRegionTrack
	RegionTopTrack   = map[string]Track{}
	AllTracks        = []Track{}
)

func ServerInitListening() {
	router := chi.NewRouter()
	router.Get("/region/{name}", GetRegionTopTrack)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		err := http.ListenAndServe(ListeningPort, router)
		if err != nil {
			panic(err)
		}
	}()

	APIClient()
	wg.Wait()

}

// Gets regions top track and top track's suggested tracks in the http response
func GetRegionTopTrack(response http.ResponseWriter, r *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	regionName := chi.URLParam(r, "name")
	regionName = strings.ToLower(regionName)
	_, found := RegionTopTrack[regionName]

	if !found {
		http.Error(response, "Top track is not found for this Region/Country name - "+string(regionName), http.StatusNotFound)
		return
	}

	result := map[string]interface{}{}
	topTrack := RegionTopTrack[regionName]

	//Adding the region's top track in the response
	result["Top Track of this Region - "+string(regionName)] = topTrack

	//Based on the region's top track, Also adding the suggestion tracks based on the track's tag and artist in the response
	suggestedTracks := []Track{}
	for _, track := range AllTracks {
		//Ensures top track is again not duplicates to the suggested tracks
		if track.Id == topTrack.Id {
			continue
		}
		if track.Tag == topTrack.Tag || track.ArtistInfo.Name == topTrack.ArtistInfo.Name {
			suggestedTracks = append(suggestedTracks, track)
		}
	}

	result["Suggested tracks"] = suggestedTracks
	json.NewEncoder(response).Encode(result)
}
