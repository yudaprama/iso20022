package model

// Ratio expressed as amount per quantity.
type AmountAndQuantityRatio4 struct {

	// Cash amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Quantity expressed as number.
	Quantity *DecimalNumber `xml:"Qty"`
}

func (a *AmountAndQuantityRatio4) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountAndQuantityRatio4) SetQuantity(value string) {
	a.Quantity = (*DecimalNumber)(&value)
}
