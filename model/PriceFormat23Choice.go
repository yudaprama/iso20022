package model

// Choice between a percentage price or an amount price or an unspecified price or index points.
type PriceFormat23Choice struct {

	// Price expressed as a percentage.
	PercentagePrice *PercentagePrice1 `xml:"PctgPric"`

	// Price expressed as a currency and amount.
	AmountPrice *AmountPrice3 `xml:"AmtPric"`

	// Value of the price not specified.
	NotSpecifiedPrice *PriceValueType10Code `xml:"NotSpcfdPric"`

	// Price expressed as an index points.
	IndexPoints *DecimalNumber `xml:"IndxPts"`
}

func (p *PriceFormat23Choice) AddPercentagePrice() *PercentagePrice1 {
	p.PercentagePrice = new(PercentagePrice1)
	return p.PercentagePrice
}

func (p *PriceFormat23Choice) AddAmountPrice() *AmountPrice3 {
	p.AmountPrice = new(AmountPrice3)
	return p.AmountPrice
}

func (p *PriceFormat23Choice) SetNotSpecifiedPrice(value string) {
	p.NotSpecifiedPrice = (*PriceValueType10Code)(&value)
}

func (p *PriceFormat23Choice) SetIndexPoints(value string) {
	p.IndexPoints = (*DecimalNumber)(&value)
}
