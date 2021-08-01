package main

import (
	m "covid-stats/model"
	s "covid-stats/service"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"covid-stats/api"
	"covid-stats/pkg/requests"

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
	req, _ := http.NewRequest(http.MethodGet, "/covid/summary", nil)

	r.ServeHTTP(w, req)

	var resp api.Result
	if err := json.Unmarshal([]byte(w.Body.String()), &resp); err != nil {
		t.Errorf("Unmarshal fail")
	}
	var testStruct m.Response

	// assert.Equal(t, 608 , resp.Data."Age"Group.Middle)
	assert.ObjectsAreEqual(testStruct, resp)

}

func TestGetCovidData(t *testing.T) {
	req, err := requests.Get("https://static.wongnai.com/devinterview/covid-cases.json", nil, nil, 30)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, req.Code, "expecting http status to be 200")
	assert.NotNil(t, req.Body)
}

func TestSummarize(t *testing.T) {

	mockData := []m.ResponseCovidData{
		{
			ConfirmDate:    "2021-05-04",
			No:             "",
			Age:            51,
			Gender:         "หญิง",
			GenderEn:       "Female",
			Nation:         "",
			NationEn:       "China",
			Province:       "Phrae",
			ProvinceId:     46,
			District:       "",
			ProvinceEn:     "Phrae",
			StatQuarantine: 5,
		},
		{
			ConfirmDate:    "2021-05-04",
			No:             "",
			Age:            1,
			Gender:         "หญิง",
			GenderEn:       "Female",
			Nation:         "",
			NationEn:       "India",
			Province:       "Roi Et",
			ProvinceId:     53,
			District:       "",
			ProvinceEn:     "Roi Et",
			StatQuarantine: 1,
		},
		{
			ConfirmDate:    "2021-05-04",
			No:             "",
			Age:            91,
			Gender:         "หญิง",
			GenderEn:       "Female",
			Nation:         "",
			NationEn:       "China",
			Province:       "Phrae",
			ProvinceId:     46,
			District:       "",
			ProvinceEn:     "Phrae",
			StatQuarantine: 5,
		},
		{
			ConfirmDate:    "2021-05-04",
			No:             "",
			Age:            0,
			Gender:         "",
			GenderEn:       "",
			Nation:         "",
			NationEn:       "Thailand",
			Province:       "",
			ProvinceId:     0,
			District:       "",
			ProvinceEn:     "",
			StatQuarantine: 5,
		},
	}

	var sum m.Response
	if err := s.Summarize(mockData, &sum); err != nil {
		t.Error(err)
	}

	assert.Equal(t, 1, sum.AgeGroup.Middle, "expecting middle age group to be 1")
	assert.Equal(t, 3, len(sum.Province), "expecting total province to be 3")
	assert.Equal(t, 2, sum.Province["Phrae"], "expecting sum of Phrae province to be 2")
	assert.Equal(t, 1, sum.AgeGroup.Undefined, "expecting Undefined age group to be 1")

}
