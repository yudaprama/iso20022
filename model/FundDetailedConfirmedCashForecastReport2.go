package model

// Information about a cash forecast report.
type FundDetailedConfirmedCashForecastReport2 struct {

	// Information related to the cash-in and cash-out flows for a specific trade date as a result of investment fund transactions, for example, subscriptions, redemptions or switches to/from a specified investment fund. The information provided is sorted by pre-defined criteria such as country, institution, currency or user defined criteria.
	FundCashForecastDetails []*FundCashForecast4 `xml:"FndCshFcstDtls"`

	// Estimated net cash as a result of the cash-in and cash-out flows.
	ConsolidatedNetCashForecast *NetCashForecast3 `xml:"CnsltdNetCshFcst,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (f *FundDetailedConfirmedCashForecastReport2) AddFundCashForecastDetails() *FundCashForecast4 {
	newValue := new(FundCashForecast4)
	f.FundCashForecastDetails = append(f.FundCashForecastDetails, newValue)
	return newValue
}

func (f *FundDetailedConfirmedCashForecastReport2) AddConsolidatedNetCashForecast() *NetCashForecast3 {
	f.ConsolidatedNetCashForecast = new(NetCashForecast3)
	return f.ConsolidatedNetCashForecast
}

func (f *FundDetailedConfirmedCashForecastReport2) AddExtension() *Extension1 {
	newValue := new(Extension1)
	f.Extension = append(f.Extension, newValue)
	return newValue
}
