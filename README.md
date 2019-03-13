haversine
=========

The haversine formula determines the great-circle distance between two points on a sphere given their longitudes and latitudes. Important in navigation, it is a special case of a more general formula in spherical trigonometry, the law of haversines, that relates the sides and angles of spherical triangles.

Example
-------

```go
package main

import (
	"fmt"

	haversine "github.com/lexor/haversine"
)

func main() {
	whiteHouse := haversine.Coord{Lat: 38.89768, Lon: -77.03653}
	eighteenAndF := haversine.Coord{Lat: 38.89736, Lon: -77.04173}

	fmt.Printf("%f\n", float64(haversine.Distance(whiteHouse, eighteenAndF)))
	// Return: 451.411074 (float64)
}
```