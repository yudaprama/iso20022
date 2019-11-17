package model

// Choice between a fixed amount and an unlimited amount.
type FixedAmountOrUnlimited1Choice struct {

	// Fixed amount.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Unlimited amount.
	NotLimited *Unlimited9Text `xml:"NotLtd"`
}

func (f *FixedAmountOrUnlimited1Choice) SetAmount(value, currency string) {
	f.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FixedAmountOrUnlimited1Choice) SetNotLimited(value string) {
	f.NotLimited = (*Unlimited9Text)(&value)
}
