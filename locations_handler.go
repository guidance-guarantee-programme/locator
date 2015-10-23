package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Location struct {
	Id                string  `json:"id"`
	BookingLocationId string  `json:"booking_location_id"`
	Title             string  `json:"title"`
	Address           string  `json:"address"`
	Phone             string  `json:"phone"`
	Hours             string  `json:"hours"`
	Lat               float64 `json:"lat"`
	Lng               float64 `json:"lng"`
}

type FeatureGeometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type FeatureProperties struct {
	Title             string `json:"title"`
	Address           string `json:"address"`
	BookingLocationId string `json:"booking_location_id"`
	Phone             string `json:"phone"`
	Hours             string `json:"hours"`
}

type Feature struct {
	Type       string            `json:"type"`
	Id         string            `json:"id"`
	Geometry   FeatureGeometry   `json:"geometry"`
	Properties FeatureProperties `json:"properties"`
}

type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

var (
	geojson []byte
)

func NewFeatureCollection(f []Feature) FeatureCollection {
	return FeatureCollection{
		Type:     "FeatureCollection",
		Features: f,
	}
}

func NewFeatureGeometry(l Location) FeatureGeometry {
	return FeatureGeometry{
		Type:        "Point",
		Coordinates: []float64{l.Lng, l.Lat},
	}
}

func NewFeatureProperties(l Location) FeatureProperties {
	return FeatureProperties{
		Title:             l.Title,
		Address:           l.Address,
		BookingLocationId: l.BookingLocationId,
		Phone:             l.Phone,
		Hours:             l.Hours,
	}
}

func NewFeature(l Location) Feature {
	return Feature{
		Type:       "Feature",
		Id:         l.Id,
		Geometry:   NewFeatureGeometry(l),
		Properties: NewFeatureProperties(l),
	}
}

func loadLocationsJson(path string) []Location {
	var locations []Location

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Can't read %s", path))
	}

	err = json.Unmarshal(data, &locations)
	if err != nil {
		panic(fmt.Sprintf("Can't parse %s", path))
	}

	return locations
}

func loadLocations() (locations []Location) {
	for _, filename := range []string{"nicab.json", "cita.json", "cas.json"} {
		path := fmt.Sprintf("public/js/%s", filename)
		locations = append(locations, loadLocationsJson(path)...)
	}

	return
}

func buildFeatures(locations []Location) (features []Feature) {
	for _, location := range locations {
		feature := NewFeature(location)
		features = append(features, feature)
	}

	return
}

func generateGeojson(features []Feature) []byte {
	geojson, err := json.Marshal(NewFeatureCollection(features))
	if err != nil {
		panic("Can't Generate GeoJSON")
	}

	return geojson
}

func geojsonFromLocations(locations []Location) []byte {
	features := buildFeatures(locations)
	return generateGeojson(features)
}

func init() {
	geojson = geojsonFromLocations(loadLocations())
}

func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(geojson)
}
