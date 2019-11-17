package model

// Choice between a percentage price or an amount price.
type PriceFormat5Choice struct {

	// Price expressed as a percentage.
	PercentagePrice *PercentagePrice1 `xml:"PctgPric"`

	// Price expressed as a currency and amount.
	AmountPrice *AmountPrice3 `xml:"AmtPric"`
}

func (p *PriceFormat5Choice) AddPercentagePrice() *PercentagePrice1 {
	p.PercentagePrice = new(PercentagePrice1)
	return p.PercentagePrice
}

func (p *PriceFormat5Choice) AddAmountPrice() *AmountPrice3 {
	p.AmountPrice = new(AmountPrice3)
	return p.AmountPrice
}
