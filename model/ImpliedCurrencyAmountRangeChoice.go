package model

// Choice between ranges of values in which an amount is considered valid or a specified amount value which has to be matched or unmatched to be valid.
type ImpliedCurrencyAmountRangeChoice struct {

	// Lower boundary of a range of amount values.
	FromAmount *AmountRangeBoundary1 `xml:"FrAmt"`

	// Upper boundary of a range of amount values.
	ToAmount *AmountRangeBoundary1 `xml:"ToAmt"`

	// Range of valid amount values.
	FromToAmount *FromToAmountRange `xml:"FrToAmt"`

	// Exact value an amount must match to be considered valid.
	EqualAmount *ImpliedCurrencyAndAmount `xml:"EQAmt"`

	// Value that an amount must not match to be considered valid.
	NotEqualAmount *ImpliedCurrencyAndAmount `xml:"NEQAmt"`
}

func (i *ImpliedCurrencyAmountRangeChoice) AddFromAmount() *AmountRangeBoundary1 {
	i.FromAmount = new(AmountRangeBoundary1)
	return i.FromAmount
}

func (i *ImpliedCurrencyAmountRangeChoice) AddToAmount() *AmountRangeBoundary1 {
	i.ToAmount = new(AmountRangeBoundary1)
	return i.ToAmount
}

func (i *ImpliedCurrencyAmountRangeChoice) AddFromToAmount() *FromToAmountRange {
	i.FromToAmount = new(FromToAmountRange)
	return i.FromToAmount
}

func (i *ImpliedCurrencyAmountRangeChoice) SetEqualAmount(value, currency string) {
	i.EqualAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (i *ImpliedCurrencyAmountRangeChoice) SetNotEqualAmount(value, currency string) {
	i.NotEqualAmount = NewImpliedCurrencyAndAmount(value, currency)
}
