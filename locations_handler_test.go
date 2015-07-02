package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

var geoJSON = []byte(`
{
  "type": "FeatureCollection",
  "features": [
    {
      "type": "Feature",
      "id": "2554b26e-6683-4de9-9b4f-f32d8e60cfb0",
      "geometry": {
        "type": "Point",
        "coordinates": [
          -6.188037599999999,
          54.71857499999999
        ]
      },
      "properties": {
        "title": "Antrim Citizens Advice Bureau",
        "address": "Farranshane House, 1 Ballygore Road, Antrim, BT41 2RN",
        "booking_centre": "",
        "phone": "028 9442 8176",
        "hours": "Mon-Fri 9-5"
      }
    }
  ]
}
`)

func TestParseLocationsFromJson(t *testing.T) {
	var (
		locations []Location
		location  Location
	)

	locations = loadLocationsJson("fixtures/locations.json")

	if len(locations) != 1 {
		t.Error("loadLocationsJson() should return a slice of 1 Location")
	}

	location = locations[0]

	if location.Id != "2554b26e-6683-4de9-9b4f-f32d8e60cfb0" {
		t.Error("Location id should be set")
	}

	if location.Title != "Antrim Citizens Advice Bureau" {
		t.Error("Location title should be set")
	}

	if location.Address != "Farranshane House, 1 Ballygore Road, Antrim, BT41 2RN" {
		t.Error("Location address should be set")
	}

	if location.Phone != "028 9442 8176" {
		t.Error("Location phone number should be set")
	}

	if location.Hours != "Mon-Fri 9-5" {
		t.Error("Location opening hours should be set")
	}

	if location.Lat != 54.71857499999999 {
		t.Error("Location latitude should be set")
	}

	if location.Lng != -6.188037599999999 {
		t.Error("Location longitude should be set")
	}
}

func TestGeojsonFromLocations(t *testing.T) {
	locations := loadLocationsJson("fixtures/locations.json")

	compactedGeoJson := new(bytes.Buffer)
	_ = json.Compact(compactedGeoJson, geoJSON)

	if !bytes.Equal(compactedGeoJson.Bytes(), geojsonFromLocations(locations)) {
		t.Error("GeoJSON should be generated correctly")
	}

}

func TestLocationIdsAreUnique(t *testing.T) {
	idCounts := make(map[string]int)

	for _, location := range loadLocations() {
		id := location.Id
		if idCounts[id] += 1; idCounts[id] > 1 {
			t.Error("Duplicate location Ids found")
		}
	}
}
