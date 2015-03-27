var locations = [],
  geocoder,
  map,
  markers = [],
  infoWindows = [],
  lookupMarker;

function getJSON(url, cb) {
  var xmlHttp = new XMLHttpRequest();

  xmlHttp.onreadystatechange = function() {
    if(xmlHttp.readyState == 4 && xmlHttp.status == 200) {
      cb(JSON.parse(xmlHttp.responseText));
   }
  }

  xmlHttp.open("GET", url, true);
  xmlHttp.send();
}

function closeInfoWindows() {
  for(i in infoWindows) {
    infoWindows[i].close();
  }
}

function makeHandler(marker, content) {
  var infoWindow = new google.maps.InfoWindow({
    content: content
  });
  infoWindows.push(infoWindow);

  return function() {
    closeInfoWindows();
    infoWindow.open(map, marker);
  };
}

function showNearestMarker(latLng) {
  var closest,
    closestDistance;

  for(i in markers) {
    m = markers[i];
    var distance = google.maps.geometry.spherical.computeDistanceBetween(latLng, m.position);
    if(!closestDistance || distance < closestDistance) {
      closest = m;
      closestDistance = distance;
    }
  }

  if(closest) {
    closeInfoWindows();
    new google.maps.event.trigger(closest, 'click');
  }
}

function lookupAddress() {
  var address = document.getElementById('address').value;

  geocoder.geocode({'address': address}, function(results, status) {
    if(status == google.maps.GeocoderStatus.OK) {
      map.setCenter(results[0].geometry.location);

      if(lookupMarker) {
        lookupMarker.setMap(null);
      }

      lookupMarker = new google.maps.Marker({
        map: map,
        position: results[0].geometry.location,
        icon: {
          path: google.maps.SymbolPath.BACKWARD_CLOSED_ARROW,
          scale: 5
        }
      });

      showNearestMarker(results[0].geometry.location);
    } else {
      alert('Geocode was not successful for the following reason: ' + status);
    }
  });
}

function placeMarkers() {
  for(i in locations) {
    var l = locations[i];

    markers[i] = new google.maps.Marker({
      position: new google.maps.LatLng(l['lat'], l['lng']),
      map: map,
      title: l['title']
    });

    var content = [
      '<b>', l['title'], '</b>',
      '<p>', l['address'], '</p>',
      '<p>', l['hours'], '</p>',
      l['phone']
    ].join("\n");

    google.maps.event.addListener(markers[i], 'click', makeHandler(markers[i], content));
  }
}

function initialise() {
  geocoder = new google.maps.Geocoder();

  //var latLng = new google.maps.LatLng(51.5073509, -0.12775829999998223);  // London
  var latLng = new google.maps.LatLng(54.59728500000001, -5.930119999999988); // Belfast

  map = new google.maps.Map(document.getElementById('map-canvas'), {
    zoom: 8,
    center: latLng
  });

  getJSON('/js/nicab.json', function(json) {
    locations = locations.concat(json);

    getJSON('/js/cas.json', function(json) {
      locations = locations.concat(json);

      getJSON('/js/cita.json', function(json) {
        locations = locations.concat(json);
        placeMarkers();
      });
    });
  });
}

google.maps.event.addDomListener(window, 'load', initialise);
