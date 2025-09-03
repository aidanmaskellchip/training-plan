package entities

import "errors"

var (
	TerrainRoad  = Terrain{Type: "road"}
	TerrainTrail = Terrain{Type: "trail"}
)

type Terrain struct {
	Type string `json:"name"`
}

func GetTerrainStrings() []string {
	return []string{
		TerrainRoad.Type,
		TerrainTrail.Type,
	}
}

func (t Terrain) String() string {
	return t.Type
}

func FromTerrainType(t string) (Terrain, error) {
	switch t {
	case "road":
		return TerrainRoad, nil
	case "trail":
		return TerrainTrail, nil
	default:
		return Terrain{}, errors.New("invalid terrain type")
	}
}
