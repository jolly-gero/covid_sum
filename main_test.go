package main

import (
	m "covid-stats/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"covid-stats/api"
	// "covid-stats/pkg/requests"

	"github.com/stretchr/testify/assert"
)

func TestHTTPGetSummaryEndpoint(t *testing.T) {
	r := api.RegisterApi()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/covid/summary", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "expecting http status to be 200")
	assert.NotNil(t, w.Body)

}

func TestTypeResponse(t *testing.T) {
	r := api.RegisterApi()

	w := httptest.NewRecorder()
	// req, _ := requests.Get("https://static.wongnai.com/devinterview/covid-cases.json", nil, nil, 30)
	req, _ := http.NewRequest(http.MethodGet, "/covid/summary", nil)

	r.ServeHTTP(w, req)

	var resp api.Result
	if err := json.Unmarshal([]byte(w.Body.String()), &resp); err != nil {
		t.Errorf("Unmarshal fail")
	}
	var testStruct m.Response

	// assert.Equal(t, 608 , resp.Data.AgeGroup.Middle)
	assert.ObjectsAreEqual(testStruct, resp)

}
