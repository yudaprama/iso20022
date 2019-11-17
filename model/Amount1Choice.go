package model

// Choice of amounts.
type Amount1Choice struct {

	// Amount of increase, and currency.
	IncreaseAmount *ActiveCurrencyAndAmount `xml:"IncrAmt"`

	// Amount of decrease, and currency.
	DecreaseAmount *ActiveCurrencyAndAmount `xml:"DcrAmt"`
}

func (a *Amount1Choice) SetIncreaseAmount(value, currency string) {
	a.IncreaseAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *Amount1Choice) SetDecreaseAmount(value, currency string) {
	a.DecreaseAmount = NewActiveCurrencyAndAmount(value, currency)
}
