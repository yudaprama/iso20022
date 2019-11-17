package model

// Parameters used to report cash movements,eg, country code, currency code, BIC or a user defined parameter.
type ForecastParameter2 struct {

	// Type of parameter used for grouping the information in a report, eg, country code, currency code, BIC or a user defined parameter.
	ReportParameter *ReportParameter2Choice `xml:"RptParam"`

	// Cash movement in to of a fund as a result of investment funds transactions, eg, subscriptions or switch-out.
	CashInForecastDetails []*CashInForecast1 `xml:"CshInFcstDtls,omitempty"`

	// Cash movement out of a fund as a result of investment funds transactions, eg, redemptions or switch-out.
	CashOutForecastDetails []*CashOutForecast1 `xml:"CshOutFcstDtls,omitempty"`

	// Net cash movements to a fund as a result of investment funds transactions.
	NetCashForecastDetails []*NetCashForecast1 `xml:"NetCshFcstDtls,omitempty"`
}

func (f *ForecastParameter2) AddReportParameter() *ReportParameter2Choice {
	f.ReportParameter = new(ReportParameter2Choice)
	return f.ReportParameter
}

func (f *ForecastParameter2) AddCashInForecastDetails() *CashInForecast1 {
	newValue := new(CashInForecast1)
	f.CashInForecastDetails = append(f.CashInForecastDetails, newValue)
	return newValue
}

func (f *ForecastParameter2) AddCashOutForecastDetails() *CashOutForecast1 {
	newValue := new(CashOutForecast1)
	f.CashOutForecastDetails = append(f.CashOutForecastDetails, newValue)
	return newValue
}

func (f *ForecastParameter2) AddNetCashForecastDetails() *NetCashForecast1 {
	newValue := new(NetCashForecast1)
	f.NetCashForecastDetails = append(f.NetCashForecastDetails, newValue)
	return newValue
}
