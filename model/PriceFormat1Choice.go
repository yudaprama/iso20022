package model

// Choice of formats to express a price.
type PriceFormat1Choice struct {

	// Price expressed as a currency and amount.
	Amount *AmountPrice1 `xml:"Amt"`

	// Price expressed as a rate, ie, percentage.
	Rate *PercentageRate `xml:"Rate"`

	// Price expressed as an amount per a quantity of financial instruments.
	AmountPricePerFinancialInstrumentQuantity *AmountPricePerFinancialInstrumentQuantity1 `xml:"AmtPricPerFinInstrmQty"`

	// Price expressed as an amount per another amount.
	AmountPricePerAmount *AmountPricePerAmount1 `xml:"AmtPricPerAmt"`

	// The value of the price is not specified.
	NotSpecified *PriceValueType6FormatChoice `xml:"NotSpcfd"`
}

func (p *PriceFormat1Choice) AddAmount() *AmountPrice1 {
	p.Amount = new(AmountPrice1)
	return p.Amount
}

func (p *PriceFormat1Choice) SetRate(value string) {
	p.Rate = (*PercentageRate)(&value)
}

func (p *PriceFormat1Choice) AddAmountPricePerFinancialInstrumentQuantity() *AmountPricePerFinancialInstrumentQuantity1 {
	p.AmountPricePerFinancialInstrumentQuantity = new(AmountPricePerFinancialInstrumentQuantity1)
	return p.AmountPricePerFinancialInstrumentQuantity
}

func (p *PriceFormat1Choice) AddAmountPricePerAmount() *AmountPricePerAmount1 {
	p.AmountPricePerAmount = new(AmountPricePerAmount1)
	return p.AmountPricePerAmount
}

func (p *PriceFormat1Choice) AddNotSpecified() *PriceValueType6FormatChoice {
	p.NotSpecified = new(PriceValueType6FormatChoice)
	return p.NotSpecified
}
