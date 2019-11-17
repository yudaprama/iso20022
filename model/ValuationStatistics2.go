package model

// Statistical data related to the price change of a security.
type ValuationStatistics2 struct {

	// Currency of the valuation statistics.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy"`

	// Type of price from which the change is calculated, eg, bid, offer, or single.
	PriceTypeChangeBasis *PriceType2 `xml:"PricTpChngBsis"`

	// Change in price since the last valuation.
	PriceChange *PriceValue2 `xml:"PricChng"`

	// Rate of income from the financial instrument, usually calculated as total dividends or coupon interest available to investors in the last year,divided by the current price.
	Yield *PercentageRate `xml:"Yld,omitempty"`

	// Information related to price variations, expressed using pre-defined periods.
	ByPredefinedTimePeriods *StatisticsByPredefinedTimePeriods1 `xml:"ByPrdfndTmPrds,omitempty"`

	// Information related to price variations, expressed using user-defined periods.
	ByUserDefinedTimePeriod []*StatisticsByUserDefinedTimePeriod1 `xml:"ByUsrDfndTmPrd,omitempty"`
}

func (v *ValuationStatistics2) SetCurrency(value string) {
	v.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (v *ValuationStatistics2) AddPriceTypeChangeBasis() *PriceType2 {
	v.PriceTypeChangeBasis = new(PriceType2)
	return v.PriceTypeChangeBasis
}

func (v *ValuationStatistics2) AddPriceChange() *PriceValue2 {
	v.PriceChange = new(PriceValue2)
	return v.PriceChange
}

func (v *ValuationStatistics2) SetYield(value string) {
	v.Yield = (*PercentageRate)(&value)
}

func (v *ValuationStatistics2) AddByPredefinedTimePeriods() *StatisticsByPredefinedTimePeriods1 {
	v.ByPredefinedTimePeriods = new(StatisticsByPredefinedTimePeriods1)
	return v.ByPredefinedTimePeriods
}

func (v *ValuationStatistics2) AddByUserDefinedTimePeriod() *StatisticsByUserDefinedTimePeriod1 {
	newValue := new(StatisticsByUserDefinedTimePeriod1)
	v.ByUserDefinedTimePeriod = append(v.ByUserDefinedTimePeriod, newValue)
	return newValue
}
