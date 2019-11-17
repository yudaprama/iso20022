package model

// Information about a cash forecast report.
type FundConfirmedCashForecastReport3 struct {

	// Information about the fund/sub fund when the report either specifies cash flow for the fund/sub fund or for a share class of the fund/sub fund.
	FundOrSubFundDetails []*Fund2 `xml:"FndOrSubFndDtls,omitempty"`

	// Information related to the cash-in and cash-out flows for a specific trade date as a result of investment fund transactions, for example, subscriptions, redemptions or switches to/from a specified investment fund.
	FundCashForecastDetails []*FundCashForecast7 `xml:"FndCshFcstDtls,omitempty"`

	// Estimated net cash as a result of the cash-in and cash-out flows.
	ConsolidatedNetCashForecast *NetCashForecast3 `xml:"CnsltdNetCshFcst,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (f *FundConfirmedCashForecastReport3) AddFundOrSubFundDetails() *Fund2 {
	newValue := new(Fund2)
	f.FundOrSubFundDetails = append(f.FundOrSubFundDetails, newValue)
	return newValue
}

func (f *FundConfirmedCashForecastReport3) AddFundCashForecastDetails() *FundCashForecast7 {
	newValue := new(FundCashForecast7)
	f.FundCashForecastDetails = append(f.FundCashForecastDetails, newValue)
	return newValue
}

func (f *FundConfirmedCashForecastReport3) AddConsolidatedNetCashForecast() *NetCashForecast3 {
	f.ConsolidatedNetCashForecast = new(NetCashForecast3)
	return f.ConsolidatedNetCashForecast
}

func (f *FundConfirmedCashForecastReport3) AddExtension() *Extension1 {
	newValue := new(Extension1)
	f.Extension = append(f.Extension, newValue)
	return newValue
}
