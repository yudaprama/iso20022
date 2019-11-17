package model

// Choice between a percentage price or an amount price or an unspecified price.
type PriceFormat57Choice struct {

	// Price expressed as a percentage.
	PercentagePrice *PercentagePrice1 `xml:"PctgPric"`

	// Price expressed as a currency and amount.
	AmountPrice *AmountPrice5 `xml:"AmtPric"`

	// Value of the price not specified.
	NotSpecifiedPrice *PriceValueType10Code `xml:"NotSpcfdPric"`
}

func (p *PriceFormat57Choice) AddPercentagePrice() *PercentagePrice1 {
	p.PercentagePrice = new(PercentagePrice1)
	return p.PercentagePrice
}

func (p *PriceFormat57Choice) AddAmountPrice() *AmountPrice5 {
	p.AmountPrice = new(AmountPrice5)
	return p.AmountPrice
}

func (p *PriceFormat57Choice) SetNotSpecifiedPrice(value string) {
	p.NotSpecifiedPrice = (*PriceValueType10Code)(&value)
}
