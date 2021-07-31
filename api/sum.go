package api

import (
	"covid-stats/pkg/requests"
	"encoding/json"
	"net/http"

	m "covid-stats/model"
	s "covid-stats/service"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetSummaryEndpoint(c *gin.Context) {
	covidData, err := requests.Get("https://static.wongnai.com/devinterview/covid-cases.json", nil, nil, 30)

	if err != nil {
		log.Errorln("get data from covid-cases.json error", err)
		c.JSON(http.StatusInternalServerError, Result{Error: "get data from covid-cases error"})
		return
	}

	var result m.SumResponse
	if err := json.Unmarshal(covidData.Body, &result); err != nil {
		log.Errorln("json unmarshal fail", err)
		c.JSON(http.StatusInternalServerError, Result{Error: "get data from covid-cases error --> json unmarshal fail"})
		return
	}

	var sum m.Response
	if err := s.Summarize(result.Data, &sum); err != nil {
		log.Errorln("system error summary --> ", err)
		c.JSON(http.StatusInternalServerError, Result{Error: "get summary error"})
		return
	}

	c.JSON(http.StatusOK, Result{Message: "get Covid-19 Summary", Count: len(result.Data), Data: sum})
}
