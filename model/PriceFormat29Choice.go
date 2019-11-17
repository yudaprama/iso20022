package model

// Choice between a percentage price or an amount price or an unspecified price.
type PriceFormat29Choice struct {

	// Price expressed as a currency and amount.
	AmountPrice *AmountPrice2 `xml:"AmtPric"`

	// Value of the price not specified.
	NotSpecifiedPrice *PriceValueType10Code `xml:"NotSpcfdPric"`
}

func (p *PriceFormat29Choice) AddAmountPrice() *AmountPrice2 {
	p.AmountPrice = new(AmountPrice2)
	return p.AmountPrice
}

func (p *PriceFormat29Choice) SetNotSpecifiedPrice(value string) {
	p.NotSpecifiedPrice = (*PriceValueType10Code)(&value)
}
