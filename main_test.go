package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"planet.com/route"
)

func TestAddExoplanet(t *testing.T) {
	payload := []byte(`{
	"Name": "GasGiant 1",
	"Description": "i am GasGiant 1",
	"DistanceFromEarth": 400,
	"Radius":3,
	"Mass": 0,
	"Type": "GasGiant"
  }`)
	req, err := http.NewRequest("POST", "/api/exoplanet/add_exoplanet", bytes.NewBuffer(payload))

	assert.NoError(t, err, "error in creating request")
	req.Header.Set("Content-Type", "applicatio/json")
	rr := httptest.NewRecorder()
	r := route.SetupRouter()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, rr.Body.String())
}
func TestListAllExoplanet(t *testing.T) {
	payload := []byte(`{
		"SortByRadius":"desc",
		"FilterBymass":0
	  }`)
	req, err := http.NewRequest("GET", "/api/exoplanet/list_all_exoplanet", bytes.NewBuffer(payload))

	assert.NoError(t, err, "error in creating request")
	req.Header.Set("Content-Type", "applicatio/json")
	rr := httptest.NewRecorder()
	r := route.SetupRouter()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, rr.Body.String())
}
func TestListExoplanetById(t *testing.T) {
	payload := []byte(``)
	req, err := http.NewRequest("GET", "/api/exoplanet/list_exoplanet_byid/2", bytes.NewBuffer(payload))

	assert.NoError(t, err, "error in creating request")
	req.Header.Set("Content-Type", "applicatio/json")
	rr := httptest.NewRecorder()
	r := route.SetupRouter()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, rr.Body.String())
}
func TestUpdateExoplanet(t *testing.T) {
	payload := []byte(`{
		"Id": 3,
		"Name": "GasGiant update again",
		"Description": "i am GasGiant 2",
		"DistanceFromEarth": 400,
		"Radius": 3,
		"Mass": 0,
		"Type": "GasGiant"
	  }`)
	req, err := http.NewRequest("PUT", "/api/exoplanet/update_exoplanet", bytes.NewBuffer(payload))

	assert.NoError(t, err, "error in creating request")
	req.Header.Set("Content-Type", "applicatio/json")
	rr := httptest.NewRecorder()
	r := route.SetupRouter()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, rr.Body.String())
}

func TestFuelEstimation(t *testing.T) {
	payload := []byte(`{
		"ExoPlanetId": 3,
		"CrewCapacity": 3
	  }`)
	req, err := http.NewRequest("GET", "/api/exoplanet/fuel_estimation", bytes.NewBuffer(payload))

	assert.NoError(t, err, "error in creating request")
	req.Header.Set("Content-Type", "applicatio/json")
	rr := httptest.NewRecorder()
	r := route.SetupRouter()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, rr.Body.String())
}

func TestDeleteExoplanetById(t *testing.T) {
	payload := []byte(``)
	req, err := http.NewRequest("DELETE", "/api/exoplanet/delete_exoplanet_byid/9", bytes.NewBuffer(payload))

	assert.NoError(t, err, "error in creating request")
	req.Header.Set("Content-Type", "applicatio/json")
	rr := httptest.NewRecorder()
	r := route.SetupRouter()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, rr.Body.String())
}
