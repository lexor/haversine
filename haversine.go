package haversine

import "math"

// Coord ...
type Coord struct {
	Lat float64
	Lon float64
}

// Metres ...
type Metres float64

const (
	// Earth Radius in Metres
	earthRadiusMetres = 6371000
)

// degreesToRadians ...
func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// toRadians ...
func (coord Coord) toRadians() Coord {
	return Coord{
		Lat: degreesToRadians(coord.Lat),
		Lon: degreesToRadians(coord.Lon),
	}
}

// Delta ...
func (origin Coord) Delta(point Coord) Coord {
	return Coord{
		Lat: point.Lat - origin.Lat,
		Lon: point.Lon - origin.Lon,
	}
}

// Distance ...
func Distance(origin, position Coord) Metres {
	origin = origin.toRadians()
	position = position.toRadians()

	change := origin.Delta(position)

	a := math.Sin(change.Lat/2)*math.Sin(change.Lat/2) +
		math.Cos(origin.Lat)*math.Cos(position.Lat)*
			math.Sin(change.Lon/2)*math.Sin(change.Lon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return Metres(c * earthRadiusMetres)
}
