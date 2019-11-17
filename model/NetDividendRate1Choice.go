package model

// Choice of format to express a net dividend.
type NetDividendRate1Choice struct {

	// The value of the rate is not specified, eg, the rate is unknown.
	NotSpecifiedRate *RateValueType6FormatChoice `xml:"NotSpcfdRate"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Value is expressed as an amount related to an underlying securities, eg, underlying security for which an interest is paid.
	RateTypeAmount *NetDividendRate2 `xml:"RateTpAmt"`
}

func (n *NetDividendRate1Choice) AddNotSpecifiedRate() *RateValueType6FormatChoice {
	n.NotSpecifiedRate = new(RateValueType6FormatChoice)
	return n.NotSpecifiedRate
}

func (n *NetDividendRate1Choice) SetAmount(value, currency string) {
	n.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (n *NetDividendRate1Choice) AddRateTypeAmount() *NetDividendRate2 {
	n.RateTypeAmount = new(NetDividendRate2)
	return n.RateTypeAmount
}
