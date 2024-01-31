package main

import (
	"fmt"
	"io"
	"net/http"
)

// http client request for accessing the API url and gets the response
func APIClient() {
	fmt.Println("API client request sending for the url --", ApiUrl)
	response, err := http.Get(ApiUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	byteResponse, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("API server Response status -- ", response.Status)
	fmt.Println("API server Response for the Top track and Top track's suggested tracks for region", InputRegionTrack, " -- ", string(byteResponse))
}
