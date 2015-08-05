describe('parseGeoJSON()', function() {
  var geojson, locations;

  beforeEach(function() {
    geojson = {
      "type": "FeatureCollection",
      "features": [
        {
          "type": "Feature",
          "id": "73abedb4-f822-4f5f-b7a9-6aa1f627f41a",
          "geometry": {
            "type": "Point",
            "coordinates": [
              -6.188037599999999,
              54.71857499999999
            ]
          },
          "properties": {
            "title": "Antrim",
            "address": "Farranshane House\n1 Ballygore Road\nAntrim\nBT41 2RN",
            "phone": "028 9442 8176",
            "hours": "Monday to Thursday, 10am to 12pm and 2pm to 4pm"
          }
        },
        {
          "type": "Feature",
          "id": "8862652c-a7e0-4aba-8846-91f836ebbffe",
          "geometry": {
            "type": "Point",
            "coordinates": [
              -5.7019008,
              54.5952592
            ]
          },
          "properties": {
            "title": "Ards",
            "address": "5 West Street\nNewtownards\nBT23 4EN",
            "booking_location_id": "73abedb4-f822-4f5f-b7a9-6aa1f627f41a",
            "phone": "",
            "hours": ""
          }
        }
      ]
    };

    locations = parseGeoJSON(geojson);
  });

  it('returns the correct number of locations', function() {
    expect(locations.length).toEqual(2);
  });

  describe('when the location manages its own booking', function() {
    it('generates the location', function() {
      var
        location = locations[0],
        expected = {
          "title": "Antrim",
          "address": "Farranshane House\n1 Ballygore Road\nAntrim\nBT41 2RN",
          "phone": "028 9442 8176",
          "hours": "Monday to Thursday, 10am to 12pm and 2pm to 4pm",
          "lat": 54.71857499999999,
          "lng": -6.188037599999999
        };

      for(i in expected) {
        expect(location[i]).toEqual(expected[i]);
      }
    });
  });

  describe('when the location does not manages its own booking', function() {
    it('generates the location', function() {
      var
        location = locations[1],
        expected = {
          "title": "Ards",
          "address": "5 West Street\nNewtownards\nBT23 4EN",
          "booking_centre": "Antrim",
          "booking_location_id": "73abedb4-f822-4f5f-b7a9-6aa1f627f41a",
          "phone": "028 9442 8176",
          "hours": "Monday to Thursday, 10am to 12pm and 2pm to 4pm",
          "lat": 54.5952592,
          "lng": -5.7019008
        };

      for(i in expected) {
        expect(location[i]).toEqual(expected[i]);
      }
    });
  });

  describe('when the referenced booking location does not exist', function() {
    it('does not throw an error', function() {
      delete(geojson['features'][0]);

      expect(function() {
        parseGeoJSON(geojson);
      }).not.toThrowError();
    });
  });
});
