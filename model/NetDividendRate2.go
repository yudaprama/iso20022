package model

// Cash dividend amount per equity after deductions or allowances have been made.
type NetDividendRate2 struct {

	// Type of underlying securities to which the rate is related, eg, underlying security for which an interest is paid.
	RateType *NetDividendRateType1FormatChoice `xml:"RateTp"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`
}

func (n *NetDividendRate2) AddRateType() *NetDividendRateType1FormatChoice {
	n.RateType = new(NetDividendRateType1FormatChoice)
	return n.RateType
}

func (n *NetDividendRate2) SetAmount(value, currency string) {
	n.Amount = NewActiveCurrencyAndAmount(value, currency)
}
