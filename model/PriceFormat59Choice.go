package model

// Choice between a percentage price or an amount price or an unspecified price or index points.
type PriceFormat59Choice struct {

	// Price expressed as a percentage.
	PercentagePrice *PercentagePrice1 `xml:"PctgPric"`

	// Price expressed as a currency and amount.
	AmountPrice *AmountPrice5 `xml:"AmtPric"`

	// Value of the price not specified.
	NotSpecifiedPrice *PriceValueType10Code `xml:"NotSpcfdPric"`

	// Price expressed as an index points.
	IndexPoints *RestrictedFINDecimalNumber `xml:"IndxPts"`
}

func (p *PriceFormat59Choice) AddPercentagePrice() *PercentagePrice1 {
	p.PercentagePrice = new(PercentagePrice1)
	return p.PercentagePrice
}

func (p *PriceFormat59Choice) AddAmountPrice() *AmountPrice5 {
	p.AmountPrice = new(AmountPrice5)
	return p.AmountPrice
}

func (p *PriceFormat59Choice) SetNotSpecifiedPrice(value string) {
	p.NotSpecifiedPrice = (*PriceValueType10Code)(&value)
}

func (p *PriceFormat59Choice) SetIndexPoints(value string) {
	p.IndexPoints = (*RestrictedFINDecimalNumber)(&value)
}
