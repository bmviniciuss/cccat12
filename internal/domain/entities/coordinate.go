package entities

import "math"

const (
	earthRadiusMeters = 6371000
)

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

func NewCoordinate(latitude, longitude float64) *Coordinate {
	return &Coordinate{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

func (c Coordinate) DistanceInMeters(c2 Coordinate) float64 {
	c1LatRad, c1LonRad := toRadians(c.Latitude), toRadians(c.Longitude)
	c2LatRad, c2LonRad := toRadians(c2.Latitude), toRadians(c2.Longitude)

	deltaLat := c2LatRad - c1LatRad
	deltaLon := c2LonRad - c1LonRad

	a := math.Pow(math.Sin(deltaLat/2), 2) + math.Cos(c1LatRad)*math.Cos(c2LatRad)*math.Pow(math.Sin(deltaLon/2), 2)
	dist := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return earthRadiusMeters * dist
}

func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
