package main

import (
	"covid-stats/api"
)

func main() {
	r := api.RegisterApi()
	r.Run(":8000") // listen and serve on localhost:8000
}