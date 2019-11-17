package model

// Choice between a percentage price or an amount price or an unspecified price.
type PriceFormat58Choice struct {

	// Price expressed as a currency and amount.
	AmountPrice *AmountPrice4 `xml:"AmtPric"`

	// Value of the price not specified.
	NotSpecifiedPrice *PriceValueType10Code `xml:"NotSpcfdPric"`
}

func (p *PriceFormat58Choice) AddAmountPrice() *AmountPrice4 {
	p.AmountPrice = new(AmountPrice4)
	return p.AmountPrice
}

func (p *PriceFormat58Choice) SetNotSpecifiedPrice(value string) {
	p.NotSpecifiedPrice = (*PriceValueType10Code)(&value)
}
