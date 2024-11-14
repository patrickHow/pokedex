package pokeapi

type PokeAPI struct {
	requestManager *RequestManager
	mapDataManager *MapDataManager
}

func NewPokeAPI() *PokeAPI {
	return &PokeAPI{
		requestManager: NewRequestManager(),
		mapDataManager: NewMapDataManager(),
	}
}

func (api *PokeAPI) GetNextMapData() (MapData, error) {
	return api.mapDataManager.NextMapData(api.requestManager)
}

func (api *PokeAPI) GetPrevMapData() (MapData, error) {
	return api.mapDataManager.PrevMapData(api.requestManager)
}
