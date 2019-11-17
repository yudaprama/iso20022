package model

// Ratio expressed as amount per quantity.
type AmountAndQuantityRatio5 struct {

	// Cash amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Quantity expressed as number.
	Quantity *RestrictedFINDecimalNumber `xml:"Qty"`
}

func (a *AmountAndQuantityRatio5) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountAndQuantityRatio5) SetQuantity(value string) {
	a.Quantity = (*RestrictedFINDecimalNumber)(&value)
}
