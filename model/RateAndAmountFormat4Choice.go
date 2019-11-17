package model

// Choice between an amount or an unspecified rate.
type RateAndAmountFormat4Choice struct {

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType6Code `xml:"NotSpcfdRate"`
}

func (r *RateAndAmountFormat4Choice) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateAndAmountFormat4Choice) SetNotSpecifiedRate(value string) {
	r.NotSpecifiedRate = (*RateValueType6Code)(&value)
}
