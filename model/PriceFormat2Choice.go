package model

// Choice of formats to express a price.
type PriceFormat2Choice struct {

	// Price expressed as a currency and amount.
	Amount *AmountPrice1 `xml:"Amt"`

	// Price expressed as a rate, ie, percentage.
	Rate *PriceRate1 `xml:"Rate"`

	// The value of the price is not specified.
	NotSpecified *PriceValueType5FormatChoice `xml:"NotSpcfd"`
}

func (p *PriceFormat2Choice) AddAmount() *AmountPrice1 {
	p.Amount = new(AmountPrice1)
	return p.Amount
}

func (p *PriceFormat2Choice) AddRate() *PriceRate1 {
	p.Rate = new(PriceRate1)
	return p.Rate
}

func (p *PriceFormat2Choice) AddNotSpecified() *PriceValueType5FormatChoice {
	p.NotSpecified = new(PriceValueType5FormatChoice)
	return p.NotSpecified
}
