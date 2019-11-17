package model

// Price expressed as an actual amount.
type AmountPrice2 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType2Code `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *ActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`
}

func (a *AmountPrice2) SetAmountPriceType(value string) {
	a.AmountPriceType = (*AmountPriceType2Code)(&value)
}

func (a *AmountPrice2) SetPriceValue(value, currency string) {
	a.PriceValue = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}
