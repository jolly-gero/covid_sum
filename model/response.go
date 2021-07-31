package model

type ageGroup struct {
	Young     int `json:"0-30"`
	Middle    int `json:"31-60"`
	Old       int `json:"61+"`
	Undefined int `json:"N/A"`
}

type Response struct {
	Province interface{} `json:"Province"`
	AgeGroup ageGroup    `json:"AgeGroup"`
}
