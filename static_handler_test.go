package main

import (
  "net/http"
  "net/http/httptest"
  "strings"
	"testing"
)

func BuildFakeGetenv(vars map[string]string) func(name string) string {
  return func(name string) string {
    return vars[name]
  }
}

func TestGoogleMapsApiKey(t *testing.T) {
  vars := make(map[string]string)
  vars["GOOGLE_MAP_API_KEY"] = "DUMMYGOOGLEMAPSAPIKEY"
  getEnv = BuildFakeGetenv(vars)

  req, err := http.NewRequest("GET", "/", nil)
  if err != nil {
      t.Fatal(err)
  }

  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(StaticHandler)
  handler.ServeHTTP(rr, req)

  expected := "https://maps.googleapis.com/maps/api/js?key=DUMMYGOOGLEMAPSAPIKEYregion=GB&libraries=geometry"
  if !strings.Contains(rr.Body.String(), expected) {
    t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
  }
}
