package model

// Price expressed as an amount.
type AmountPrice5 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType1Code `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`
}

func (a *AmountPrice5) SetAmountPriceType(value string) {
	a.AmountPriceType = (*AmountPriceType1Code)(&value)
}

func (a *AmountPrice5) SetPriceValue(value, currency string) {
	a.PriceValue = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}
