package model

// Information about a cash forecast report.
type FundConfirmedCashForecastReport1 struct {

	// Information related to the cash-in and cash-out flows for a specific trade date as a result of investment fund transactions, for example, subscriptions, redemptions or switches to/from a specified investment fund.
	FundCashForecastDetails []*FundCashForecast1 `xml:"FndCshFcstDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (f *FundConfirmedCashForecastReport1) AddFundCashForecastDetails() *FundCashForecast1 {
	newValue := new(FundCashForecast1)
	f.FundCashForecastDetails = append(f.FundCashForecastDetails, newValue)
	return newValue
}

func (f *FundConfirmedCashForecastReport1) AddExtension() *Extension1 {
	newValue := new(Extension1)
	f.Extension = append(f.Extension, newValue)
	return newValue
}
