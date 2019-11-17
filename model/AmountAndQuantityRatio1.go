package model

// Ratio expressed as amount per quantity.
type AmountAndQuantityRatio1 struct {

	// Cash amount.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Quantity expressed as number.
	Quantity *DecimalNumber `xml:"Qty"`
}

func (a *AmountAndQuantityRatio1) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndQuantityRatio1) SetQuantity(value string) {
	a.Quantity = (*DecimalNumber)(&value)
}
