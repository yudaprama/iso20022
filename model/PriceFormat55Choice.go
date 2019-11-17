package model

// Choice between a percentage price or an amount price or index points.
type PriceFormat55Choice struct {

	// Price expressed as a percentage.
	PercentagePrice *PercentagePrice1 `xml:"PctgPric"`

	// Price expressed as a currency and amount.
	AmountPrice *AmountPrice5 `xml:"AmtPric"`

	// Price expressed as an index points
	IndexPoints *RestrictedFINDecimalNumber `xml:"IndxPts"`
}

func (p *PriceFormat55Choice) AddPercentagePrice() *PercentagePrice1 {
	p.PercentagePrice = new(PercentagePrice1)
	return p.PercentagePrice
}

func (p *PriceFormat55Choice) AddAmountPrice() *AmountPrice5 {
	p.AmountPrice = new(AmountPrice5)
	return p.AmountPrice
}

func (p *PriceFormat55Choice) SetIndexPoints(value string) {
	p.IndexPoints = (*RestrictedFINDecimalNumber)(&value)
}
