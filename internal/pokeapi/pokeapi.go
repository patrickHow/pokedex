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

func (api *PokeAPI) GetAreaData(area string) (AreaData, error) {
	return GetAreaData(api.requestManager, area)
}

func (api *PokeAPI) GetPokemonData(name string) (PokemonData, error) {
	return GetPokemonData(api.requestManager, name)
}
