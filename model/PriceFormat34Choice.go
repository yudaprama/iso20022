package model

// Choice between a percentage price or an amount price or an amount price per amount or an amount price per financial instrument quantity.
type PriceFormat34Choice struct {

	// Price expressed as a percentage.
	PercentagePrice *PercentagePrice1 `xml:"PctgPric"`

	// Price expressed as a currency and amount.
	AmountPrice *AmountPrice3 `xml:"AmtPric"`

	// Price expressed as a ratio: amount price per financial instrument quantity.
	AmountPricePerFinancialInstrumentQuantity *AmountPricePerFinancialInstrumentQuantity3 `xml:"AmtPricPerFinInstrmQty"`

	// Price expressed as a ratio: amount price per amount
	AmountPricePerAmount *AmountPricePerAmount2 `xml:"AmtPricPerAmt"`

	// Price expressed in index points.
	IndexPoints *DecimalNumber `xml:"IndxPts"`
}

func (p *PriceFormat34Choice) AddPercentagePrice() *PercentagePrice1 {
	p.PercentagePrice = new(PercentagePrice1)
	return p.PercentagePrice
}

func (p *PriceFormat34Choice) AddAmountPrice() *AmountPrice3 {
	p.AmountPrice = new(AmountPrice3)
	return p.AmountPrice
}

func (p *PriceFormat34Choice) AddAmountPricePerFinancialInstrumentQuantity() *AmountPricePerFinancialInstrumentQuantity3 {
	p.AmountPricePerFinancialInstrumentQuantity = new(AmountPricePerFinancialInstrumentQuantity3)
	return p.AmountPricePerFinancialInstrumentQuantity
}

func (p *PriceFormat34Choice) AddAmountPricePerAmount() *AmountPricePerAmount2 {
	p.AmountPricePerAmount = new(AmountPricePerAmount2)
	return p.AmountPricePerAmount
}

func (p *PriceFormat34Choice) SetIndexPoints(value string) {
	p.IndexPoints = (*DecimalNumber)(&value)
}
