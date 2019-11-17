package model

// Specifies a ratio: amount price per financial instrument quantity.
type AmountPricePerFinancialInstrumentQuantity7 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType1Code `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`

	// Quantity of financial instrument.
	FinancialInstrumentQuantity *FinancialInstrumentQuantity15Choice `xml:"FinInstrmQty"`
}

func (a *AmountPricePerFinancialInstrumentQuantity7) SetAmountPriceType(value string) {
	a.AmountPriceType = (*AmountPriceType1Code)(&value)
}

func (a *AmountPricePerFinancialInstrumentQuantity7) SetPriceValue(value, currency string) {
	a.PriceValue = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountPricePerFinancialInstrumentQuantity7) AddFinancialInstrumentQuantity() *FinancialInstrumentQuantity15Choice {
	a.FinancialInstrumentQuantity = new(FinancialInstrumentQuantity15Choice)
	return a.FinancialInstrumentQuantity
}
