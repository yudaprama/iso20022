package model

// Choice between profit and loss.
type ProfitAndLoss2Choice struct {

	// Value of the positive amount.
	Profit *ActiveCurrencyAndAmount `xml:"Prft"`

	// Value of the negative amount.
	Loss *ActiveCurrencyAndAmount `xml:"Loss"`
}

func (p *ProfitAndLoss2Choice) SetProfit(value, currency string) {
	p.Profit = NewActiveCurrencyAndAmount(value, currency)
}

func (p *ProfitAndLoss2Choice) SetLoss(value, currency string) {
	p.Loss = NewActiveCurrencyAndAmount(value, currency)
}
