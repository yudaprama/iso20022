package model

// Specifies a ratio: amount price per financial instrument quantity.
type AmountPricePerFinancialInstrumentQuantity3 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType1Code `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *ActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`

	// Quantity of financial instrument.
	FinancialInstrumentQuantity *FinancialInstrumentQuantity1Choice `xml:"FinInstrmQty"`
}

func (a *AmountPricePerFinancialInstrumentQuantity3) SetAmountPriceType(value string) {
	a.AmountPriceType = (*AmountPriceType1Code)(&value)
}

func (a *AmountPricePerFinancialInstrumentQuantity3) SetPriceValue(value, currency string) {
	a.PriceValue = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountPricePerFinancialInstrumentQuantity3) AddFinancialInstrumentQuantity() *FinancialInstrumentQuantity1Choice {
	a.FinancialInstrumentQuantity = new(FinancialInstrumentQuantity1Choice)
	return a.FinancialInstrumentQuantity
}
