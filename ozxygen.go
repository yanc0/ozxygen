// Licence MIT
// Author Yann Coleu <y@nn-col.eu>

package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Rad convert degrees to radians
func Rad(d float64) float64 {
	return d * (math.Pi / 180)
}

// NumTiles return number of tiles in a given zoom
func NumTiles(z float64) float64 {
	return math.Pow(2, z)
}

// Sec is the secant function
func Sec(x float64) float64 {
	return 1 / math.Cos(x)
}

// LatLon2RelativeXY return X, Y from a given lat, lon
func LatLon2RelativeXY(lat, lon float64) (float64, float64) {
	x := (lon + 180) / 360
	y := (1 - math.Log(math.Tan(Rad(lat))+Sec(Rad(lat)))/math.Pi) / 2
	return x, y
}

// LatLon2XY convert x y in a osm way with zoom
func LatLon2XY(lat, lon float64, z int) (int, int) {
	n := NumTiles(float64(z))
	x, y := LatLon2RelativeXY(lat, lon)
	return int(n * x), int(n * y)
}

func random(min, max float64) float64 {
	
	diff := max - min
	return rand.Float64()*diff + min
}

func main() {
	baseurl := flag.String("baseurl", "http://127.0.0.1/osm/", "HTTP base url")
	extension := flag.String("extension", ".png", "End of the URI")
	z := flag.Int("z", 18, "Zoom level")
	minlat := flag.Float64("minlat", 48.5, "Min latitude")
	maxlat := flag.Float64("maxlat", 49.2, "Max latitude")
	minlon := flag.Float64("minlon", 1.69, "Min longitude")
	maxlon := flag.Float64("maxlon", 2.98, "Max longitude")
	num := flag.Int("num", 1, "Number of results")
	flag.Parse()
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < *num; i++ {
		randlat := random(*minlat, *maxlat)
		randlon := random(*minlon, *maxlon)
		x, y := LatLon2XY(randlat, randlon, *z)
		fmt.Print(*baseurl, *z, "/", x, "/", y, *extension, "\n")
	}
}
