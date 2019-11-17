package model

// Statistical data related to the price change of a security.
type ValuationStatistics3 struct {

	// Currency of the valuation statistics.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy"`

	// Type of price from which the change is calculated, eg, bid, offer, or single.
	PriceTypeChangeBasis *PriceType2 `xml:"PricTpChngBsis"`

	// Change in price since the previous valuation date.
	PriceChange *PriceValueChange1 `xml:"PricChng"`

	// Rate of income from the financial instrument, usually calculated as total dividends or coupon interest available to investors in the last year,divided by the current price.
	Yield *PercentageRate `xml:"Yld,omitempty"`

	// Information related to price variations, expressed using pre-defined periods.
	ByPredefinedTimePeriods *StatisticsByPredefinedTimePeriods2 `xml:"ByPrdfndTmPrds,omitempty"`

	// Information related to price variations, expressed using user-defined periods.
	ByUserDefinedTimePeriod []*StatisticsByUserDefinedTimePeriod2 `xml:"ByUsrDfndTmPrd,omitempty"`
}

func (v *ValuationStatistics3) SetCurrency(value string) {
	v.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (v *ValuationStatistics3) AddPriceTypeChangeBasis() *PriceType2 {
	v.PriceTypeChangeBasis = new(PriceType2)
	return v.PriceTypeChangeBasis
}

func (v *ValuationStatistics3) AddPriceChange() *PriceValueChange1 {
	v.PriceChange = new(PriceValueChange1)
	return v.PriceChange
}

func (v *ValuationStatistics3) SetYield(value string) {
	v.Yield = (*PercentageRate)(&value)
}

func (v *ValuationStatistics3) AddByPredefinedTimePeriods() *StatisticsByPredefinedTimePeriods2 {
	v.ByPredefinedTimePeriods = new(StatisticsByPredefinedTimePeriods2)
	return v.ByPredefinedTimePeriods
}

func (v *ValuationStatistics3) AddByUserDefinedTimePeriod() *StatisticsByUserDefinedTimePeriod2 {
	newValue := new(StatisticsByUserDefinedTimePeriod2)
	v.ByUserDefinedTimePeriod = append(v.ByUserDefinedTimePeriod, newValue)
	return newValue
}
