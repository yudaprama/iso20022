package model

// Choice of formats to express a price.
type PriceFormat3Choice struct {

	// Price expressed as a currency and amount.
	Amount *AmountPrice1 `xml:"Amt"`

	// Price expressed as a rate, ie, percentage.
	Rate *PriceRate1 `xml:"Rate"`
}

func (p *PriceFormat3Choice) AddAmount() *AmountPrice1 {
	p.Amount = new(AmountPrice1)
	return p.Amount
}

func (p *PriceFormat3Choice) AddRate() *PriceRate1 {
	p.Rate = new(PriceRate1)
	return p.Rate
}
