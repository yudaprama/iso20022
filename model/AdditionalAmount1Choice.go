package model

// Choice between additional cash in or resulting cash out.
type AdditionalAmount1Choice struct {

	// Additional amount of money paid by the investor in addition to the switch redemption amount.
	AdditionalCashIn *ActiveOrHistoricCurrencyAndAmount `xml:"AddtlCshIn"`

	// Amount of money that results from a switch-out, that is not reinvested in another investment fund, and is repaid to the investor.
	ResultingCashOut *ActiveOrHistoricCurrencyAndAmount `xml:"RsltgCshOut"`
}

func (a *AdditionalAmount1Choice) SetAdditionalCashIn(value, currency string) {
	a.AdditionalCashIn = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AdditionalAmount1Choice) SetResultingCashOut(value, currency string) {
	a.ResultingCashOut = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}
