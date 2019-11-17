package model

// Price expressed as an amount.
type AmountPrice1 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType1FormatChoice `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *ActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`
}

func (a *AmountPrice1) AddAmountPriceType() *AmountPriceType1FormatChoice {
	a.AmountPriceType = new(AmountPriceType1FormatChoice)
	return a.AmountPriceType
}

func (a *AmountPrice1) SetPriceValue(value, currency string) {
	a.PriceValue = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}
