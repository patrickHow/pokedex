package pokeapi

import (
	"encoding/json"
	"fmt"
)

type MapData struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const mapBaseURL string = "https://pokeapi.co/api/v2/location-area"

type MapDataManager struct {
	nextURL string
	prevURL *string
}

// Get a new map data manager
// Ensures the request URL is initalized to the base map URL
func NewMapDataManager() *MapDataManager {
	return &MapDataManager{
		nextURL: mapBaseURL,
		prevURL: nil,
	}
}

func (mgr *MapDataManager) NextMapData(rqm *RequestManager) (MapData, error) {
	mapdata, err := requestMapData(rqm, mgr.nextURL)

	// If request succeeded, update the URLs
	if err == nil {
		mgr.nextURL = mapdata.Next
		mgr.prevURL = mapdata.Previous
	}

	return mapdata, err
}

func (mgr *MapDataManager) PrevMapData(rqm *RequestManager) (MapData, error) {
	if mgr.prevURL == nil {
		return MapData{}, fmt.Errorf("error: no previous map data")
	}

	mapdata, err := requestMapData(rqm, *mgr.prevURL)

	// If request succeeded, update the URLs
	if err == nil {
		mgr.nextURL = mapdata.Next
		mgr.prevURL = mapdata.Previous
	}

	return mapdata, err
}

func requestMapData(rqm *RequestManager, url string) (MapData, error) {

	data, err := rqm.GetData(url)
	if err != nil {
		return MapData{}, err
	}

	// Unmarshal raw bytes to struct
	var mapdata MapData
	err = json.Unmarshal(data, &mapdata)
	if err != nil {
		return MapData{}, err
	}

	return mapdata, nil
}
