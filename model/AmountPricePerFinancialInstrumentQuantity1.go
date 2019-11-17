package model

// Specifies a ratio: amount price per financial instrument quantity.
type AmountPricePerFinancialInstrumentQuantity1 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType1FormatChoice `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *ActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`

	// Quantity of financial instrument.
	FinancialInstrumentQuantity *UnitOrFaceAmount1Choice `xml:"FinInstrmQty"`
}

func (a *AmountPricePerFinancialInstrumentQuantity1) AddAmountPriceType() *AmountPriceType1FormatChoice {
	a.AmountPriceType = new(AmountPriceType1FormatChoice)
	return a.AmountPriceType
}

func (a *AmountPricePerFinancialInstrumentQuantity1) SetPriceValue(value, currency string) {
	a.PriceValue = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountPricePerFinancialInstrumentQuantity1) AddFinancialInstrumentQuantity() *UnitOrFaceAmount1Choice {
	a.FinancialInstrumentQuantity = new(UnitOrFaceAmount1Choice)
	return a.FinancialInstrumentQuantity
}
