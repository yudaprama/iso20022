package model

// Price expressed as an actual amount.
type AmountPrice4 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType2Code `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`
}

func (a *AmountPrice4) SetAmountPriceType(value string) {
	a.AmountPriceType = (*AmountPriceType2Code)(&value)
}

func (a *AmountPrice4) SetPriceValue(value, currency string) {
	a.PriceValue = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}
