package service

import (
	m "covid-stats/model"
)

func Summarize (data []m.ResponseCovidData, result *m.Response)  error{

	m := make(map[string]int)
	for _, v := range data {

		// province
		p := v.ProvinceEn
		if p == "" {
			p = "N/A"
		}		
		m[p] += 1 

		// ageGroup
		a := v.Age
		switch {
		case 0<a && a<31:   
			result.AgeGroup.Young += 1 

		case 31<=a && a<61:
			result.AgeGroup.Middle += 1 

		case a>=61:
			result.AgeGroup.Old += 1 

		default:    //! define that age : 0 is N/A
			result.AgeGroup.Undefined += 1 

		}
	}
	result.Province = m

	return nil
}

