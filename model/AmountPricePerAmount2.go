package model

// Specifies a ratio: Amount price per amount.
// Example:
// ISIN US629377AS17. Repurchase USD1087.17 cash for every USD1000 stock (NRG Energy Inc 8% Senior Notes 15/12/13).
type AmountPricePerAmount2 struct {

	// Type of amount price.
	AmountPriceType *AmountPriceType1Code `xml:"AmtPricTp"`

	// Value of the price.
	PriceValue *ActiveCurrencyAnd13DecimalAmount `xml:"PricVal"`

	// The amount on which the price is based.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`
}

func (a *AmountPricePerAmount2) SetAmountPriceType(value string) {
	a.AmountPriceType = (*AmountPriceType1Code)(&value)
}

func (a *AmountPricePerAmount2) SetPriceValue(value, currency string) {
	a.PriceValue = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountPricePerAmount2) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}
