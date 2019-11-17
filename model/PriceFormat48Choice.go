package model

// Choice between a percentage price or an amount price or an amount price per amount or an amount price per financial instrument quantity.
type PriceFormat48Choice struct {

	// Price expressed as a percentage.
	PercentagePrice *PercentagePrice1 `xml:"PctgPric"`

	// Price expressed as a currency and amount.
	AmountPrice *AmountPrice3 `xml:"AmtPric"`

	// Price expressed as a ratio: amount price per financial instrument quantity.
	AmountPricePerFinancialInstrumentQuantity *AmountPricePerFinancialInstrumentQuantity6 `xml:"AmtPricPerFinInstrmQty"`

	// Price expressed as a ratio: amount price per amount
	AmountPricePerAmount *AmountPricePerAmount2 `xml:"AmtPricPerAmt"`

	// Price expressed in index points.
	IndexPoints *DecimalNumber `xml:"IndxPts"`
}

func (p *PriceFormat48Choice) AddPercentagePrice() *PercentagePrice1 {
	p.PercentagePrice = new(PercentagePrice1)
	return p.PercentagePrice
}

func (p *PriceFormat48Choice) AddAmountPrice() *AmountPrice3 {
	p.AmountPrice = new(AmountPrice3)
	return p.AmountPrice
}

func (p *PriceFormat48Choice) AddAmountPricePerFinancialInstrumentQuantity() *AmountPricePerFinancialInstrumentQuantity6 {
	p.AmountPricePerFinancialInstrumentQuantity = new(AmountPricePerFinancialInstrumentQuantity6)
	return p.AmountPricePerFinancialInstrumentQuantity
}

func (p *PriceFormat48Choice) AddAmountPricePerAmount() *AmountPricePerAmount2 {
	p.AmountPricePerAmount = new(AmountPricePerAmount2)
	return p.AmountPricePerAmount
}

func (p *PriceFormat48Choice) SetIndexPoints(value string) {
	p.IndexPoints = (*DecimalNumber)(&value)
}
