package model

// Criterion by which the cash movements are broken down.
type CashSortingCriterion1 struct {

	// Type of criterion by which the estimated cash flow is being broken down, ie, country, institution, currency code or a user defined type, such as a region or distribution channel.
	SortingCriterionType *SortCriteria1Choice `xml:"SrtgCritnTp"`

	// Parameter for which the cash movements are reported.
	ForecastBreakdownDetails []*ForecastParameter1 `xml:"FcstBrkdwnDtls"`
}

func (c *CashSortingCriterion1) AddSortingCriterionType() *SortCriteria1Choice {
	c.SortingCriterionType = new(SortCriteria1Choice)
	return c.SortingCriterionType
}

func (c *CashSortingCriterion1) AddForecastBreakdownDetails() *ForecastParameter1 {
	newValue := new(ForecastParameter1)
	c.ForecastBreakdownDetails = append(c.ForecastBreakdownDetails, newValue)
	return newValue
}
