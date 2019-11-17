package model

// Specifies a ratio: Amount price per amount.
// Example:
// ISIN US629377AS17. Repurchase USD1087.17 cash for every USD1000 stock (NRG Energy Inc 8% Senior Notes 15/12/13).
type AmountPricePerAmount3 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType1Code `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`

	// The amount on which the price is based.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`
}

func (a *AmountPricePerAmount3) SetAmountPriceType(value string) {
	a.AmountPriceType = (*AmountPriceType1Code)(&value)
}

func (a *AmountPricePerAmount3) SetPriceValue(value, currency string) {
	a.PriceValue = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountPricePerAmount3) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}
