package model

// Information about a cash forecast report.
type FundDetailedConfirmedCashForecastReport3 struct {

	// Information about the fund/sub fund when the report either specifies cash flow for the fund/sub fund or for a share class of the fund/sub fund.
	FundOrSubFundDetails *Fund4 `xml:"FndOrSubFndDtls,omitempty"`

	// Information related to the cash-in and cash-out flows for a specific trade date as a result of transactions in shares in an investment fund, for example, subscriptions, redemptions or switches. The information provided is sorted by pre-defined criteria such as country, institution, currency or user defined criteria.
	FundCashForecastDetails []*FundCashForecast6 `xml:"FndCshFcstDtls"`

	// Estimated net cash as a result of the cash-in and cash-out flows.
	ConsolidatedNetCashForecast *NetCashForecast3 `xml:"CnsltdNetCshFcst,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (f *FundDetailedConfirmedCashForecastReport3) AddFundOrSubFundDetails() *Fund4 {
	f.FundOrSubFundDetails = new(Fund4)
	return f.FundOrSubFundDetails
}

func (f *FundDetailedConfirmedCashForecastReport3) AddFundCashForecastDetails() *FundCashForecast6 {
	newValue := new(FundCashForecast6)
	f.FundCashForecastDetails = append(f.FundCashForecastDetails, newValue)
	return newValue
}

func (f *FundDetailedConfirmedCashForecastReport3) AddConsolidatedNetCashForecast() *NetCashForecast3 {
	f.ConsolidatedNetCashForecast = new(NetCashForecast3)
	return f.ConsolidatedNetCashForecast
}

func (f *FundDetailedConfirmedCashForecastReport3) AddExtension() *Extension1 {
	newValue := new(Extension1)
	f.Extension = append(f.Extension, newValue)
	return newValue
}
