package model

// Statistical data related to the price change of a security.
type StatisticsByPredefinedTimePeriods2 struct {

	// Highest price for the referenced period.
	HighestPriceValue12Months *PriceValue5 `xml:"HghstPricVal12Mnths,omitempty"`

	// Lowest price for the referenced period.
	LowestPriceValue12Months *PriceValue5 `xml:"LwstPricVal12Mnths,omitempty"`

	// Change in price over a one year period.
	OneYearPriceChange *PriceValueChange1 `xml:"OneYrPricChng,omitempty"`

	// Change in price over a three year period.
	ThreeYearPriceChange *PriceValueChange1 `xml:"ThreeYrPricChng,omitempty"`

	// Change in price over a five year period.
	FiveYearPriceChange *PriceValueChange1 `xml:"FiveYrPricChng,omitempty"`
}

func (s *StatisticsByPredefinedTimePeriods2) AddHighestPriceValue12Months() *PriceValue5 {
	s.HighestPriceValue12Months = new(PriceValue5)
	return s.HighestPriceValue12Months
}

func (s *StatisticsByPredefinedTimePeriods2) AddLowestPriceValue12Months() *PriceValue5 {
	s.LowestPriceValue12Months = new(PriceValue5)
	return s.LowestPriceValue12Months
}

func (s *StatisticsByPredefinedTimePeriods2) AddOneYearPriceChange() *PriceValueChange1 {
	s.OneYearPriceChange = new(PriceValueChange1)
	return s.OneYearPriceChange
}

func (s *StatisticsByPredefinedTimePeriods2) AddThreeYearPriceChange() *PriceValueChange1 {
	s.ThreeYearPriceChange = new(PriceValueChange1)
	return s.ThreeYearPriceChange
}

func (s *StatisticsByPredefinedTimePeriods2) AddFiveYearPriceChange() *PriceValueChange1 {
	s.FiveYearPriceChange = new(PriceValueChange1)
	return s.FiveYearPriceChange
}
