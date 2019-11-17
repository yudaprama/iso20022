package model

// Choice between a percentage price or an amount price.
type PriceFormat52Choice struct {

	// Price expressed as a percentage.
	PercentagePrice *PercentagePrice1 `xml:"PctgPric"`

	// Price expressed as a currency and amount.
	AmountPrice *AmountPrice5 `xml:"AmtPric"`
}

func (p *PriceFormat52Choice) AddPercentagePrice() *PercentagePrice1 {
	p.PercentagePrice = new(PercentagePrice1)
	return p.PercentagePrice
}

func (p *PriceFormat52Choice) AddAmountPrice() *AmountPrice5 {
	p.AmountPrice = new(AmountPrice5)
	return p.AmountPrice
}
