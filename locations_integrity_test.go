package main

import (
	"testing"
)

func TestLocationIdsAreUnique(t *testing.T) {
	idCounts := make(map[string]int)

	for _, location := range loadLocations() {
		id := location.Id
		if idCounts[id] += 1; idCounts[id] > 1 {
			t.Error("Duplicate location Ids found")
		}
	}
}

func TestBookingLocationIdsAreValidLocations(t *testing.T) {
	locations := loadLocations()

	for _, location := range locations {
		if location.BookingLocationId == "" {
			continue;
		}

		bookingLocationFound := false;

		for _, otherLocation := range locations {
			if location.BookingLocationId == otherLocation.Id {
				bookingLocationFound = true;
				break;
			}
		}

		if !bookingLocationFound {
			t.Error("Booking location not found:", location.BookingLocationId)
		}
	}
}
