package model

// Choice between profit and loss.
type ProfitAndLoss1Choice struct {

	// Value of the positive amount.
	Profit *ActiveCurrencyAnd13DecimalAmount `xml:"Prft"`

	// Value of the negative amount.
	Loss *ActiveCurrencyAnd13DecimalAmount `xml:"Loss"`
}

func (p *ProfitAndLoss1Choice) SetProfit(value, currency string) {
	p.Profit = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (p *ProfitAndLoss1Choice) SetLoss(value, currency string) {
	p.Loss = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}
