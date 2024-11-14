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

type AreaData struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
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
	mapData, err := requestMapData(rqm, mgr.nextURL)

	// If request succeeded, update the URLs
	if err == nil {
		mgr.nextURL = mapData.Next
		mgr.prevURL = mapData.Previous
	}

	return mapData, err
}

func (mgr *MapDataManager) PrevMapData(rqm *RequestManager) (MapData, error) {
	if mgr.prevURL == nil {
		return MapData{}, fmt.Errorf("error: no previous map data")
	}

	mapData, err := requestMapData(rqm, *mgr.prevURL)

	// If request succeeded, update the URLs
	if err == nil {
		mgr.nextURL = mapData.Next
		mgr.prevURL = mapData.Previous
	}

	return mapData, err
}

func GetAreaData(rqm *RequestManager, area string) (AreaData, error) {
	url := mapBaseURL + "/" + area

	data, err := rqm.GetData(url)
	if err != nil {
		return AreaData{}, err
	}

	var areaData AreaData
	err = json.Unmarshal(data, &areaData)
	if err != nil {
		return AreaData{}, err
	}

	return areaData, nil
}

func requestMapData(rqm *RequestManager, url string) (MapData, error) {

	data, err := rqm.GetData(url)
	if err != nil {
		return MapData{}, err
	}

	// Unmarshal raw bytes to struct
	var mapData MapData
	err = json.Unmarshal(data, &mapData)
	if err != nil {
		return MapData{}, err
	}

	return mapData, nil
}
