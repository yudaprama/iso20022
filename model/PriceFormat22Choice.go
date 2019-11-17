package model

// Choice between a percentage price or an amount price or an unspecified price or an amount price per amount or an amount price per financial instrument quantity.
type PriceFormat22Choice struct {

	// Price expressed as a percentage.
	PercentagePrice *PercentagePrice1 `xml:"PctgPric"`

	// Price expressed as a currency and amount.
	AmountPrice *AmountPrice3 `xml:"AmtPric"`

	// Value of the price not specified.
	NotSpecifiedPrice *PriceValueType10Code `xml:"NotSpcfdPric"`

	// Price expressed as a ratio: amount price per financial instrument quantity.
	AmountPricePerFinancialInstrumentQuantity *AmountPricePerFinancialInstrumentQuantity3 `xml:"AmtPricPerFinInstrmQty"`

	// Price expressed as a ratio: amount price per amount
	AmountPricePerAmount *AmountPricePerAmount2 `xml:"AmtPricPerAmt"`
}

func (p *PriceFormat22Choice) AddPercentagePrice() *PercentagePrice1 {
	p.PercentagePrice = new(PercentagePrice1)
	return p.PercentagePrice
}

func (p *PriceFormat22Choice) AddAmountPrice() *AmountPrice3 {
	p.AmountPrice = new(AmountPrice3)
	return p.AmountPrice
}

func (p *PriceFormat22Choice) SetNotSpecifiedPrice(value string) {
	p.NotSpecifiedPrice = (*PriceValueType10Code)(&value)
}

func (p *PriceFormat22Choice) AddAmountPricePerFinancialInstrumentQuantity() *AmountPricePerFinancialInstrumentQuantity3 {
	p.AmountPricePerFinancialInstrumentQuantity = new(AmountPricePerFinancialInstrumentQuantity3)
	return p.AmountPricePerFinancialInstrumentQuantity
}

func (p *PriceFormat22Choice) AddAmountPricePerAmount() *AmountPricePerAmount2 {
	p.AmountPricePerAmount = new(AmountPricePerAmount2)
	return p.AmountPricePerAmount
}
