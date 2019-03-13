package haversine_test

import (
	"testing"

	haversine "github.com/lexor/haversine"
)

var exampleLocations = []struct {
	from           haversine.Coord
	to             haversine.Coord
	expectedMetres float64
}{
	{
		haversine.Coord{Lat: 38.898556, Lon: -77.037852},
		haversine.Coord{Lat: 38.897147, Lon: -77.043934},
		549.1557912048178,
	},
	{
		haversine.Coord{Lat: 39.925533, Lon: 32.866287},
		haversine.Coord{Lat: 59.9343, Lon: 30.3351},
		2.231842573694838e+06,
	},
	{
		haversine.Coord{Lat: 51.7520, Lon: 1.2577},
		haversine.Coord{Lat: 52.2053, Lon: 0.1218},
		92698.47437222341,
	},
}

func TestDistance(t *testing.T) {
	for _, input := range exampleLocations {
		metres := float64(haversine.Distance(input.from, input.to))

		if input.expectedMetres != metres {
			t.Errorf("FAIL: From: %v - To: %v - To Be: %v - Result: %v",
				input.from,
				input.to,
				input.expectedMetres,
				metres,
			)
		}
	}
}
