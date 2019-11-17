package model

// Price expressed as an amount.
type AmountPrice3 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType1Code `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *ActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`
}

func (a *AmountPrice3) SetAmountPriceType(value string) {
	a.AmountPriceType = (*AmountPriceType1Code)(&value)
}

func (a *AmountPrice3) SetPriceValue(value, currency string) {
	a.PriceValue = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}
