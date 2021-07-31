package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterApi() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/covid/summary", GetSummaryEndpoint)

	return r
}