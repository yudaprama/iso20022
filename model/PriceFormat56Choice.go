package model

// Choice between a percentage price or an amount price or an amount price per amount or an amount price per financial instrument quantity.
type PriceFormat56Choice struct {

	// Price expressed as a percentage.
	PercentagePrice *PercentagePrice1 `xml:"PctgPric"`

	// Price expressed as a currency and amount.
	AmountPrice *AmountPrice5 `xml:"AmtPric"`

	// Price expressed as a ratio: amount price per financial instrument quantity.
	AmountPricePerFinancialInstrumentQuantity *AmountPricePerFinancialInstrumentQuantity7 `xml:"AmtPricPerFinInstrmQty"`

	// Price expressed as a ratio: amount price per amount
	AmountPricePerAmount *AmountPricePerAmount3 `xml:"AmtPricPerAmt"`

	// Price expressed in index points.
	IndexPoints *RestrictedFINDecimalNumber `xml:"IndxPts"`
}

func (p *PriceFormat56Choice) AddPercentagePrice() *PercentagePrice1 {
	p.PercentagePrice = new(PercentagePrice1)
	return p.PercentagePrice
}

func (p *PriceFormat56Choice) AddAmountPrice() *AmountPrice5 {
	p.AmountPrice = new(AmountPrice5)
	return p.AmountPrice
}

func (p *PriceFormat56Choice) AddAmountPricePerFinancialInstrumentQuantity() *AmountPricePerFinancialInstrumentQuantity7 {
	p.AmountPricePerFinancialInstrumentQuantity = new(AmountPricePerFinancialInstrumentQuantity7)
	return p.AmountPricePerFinancialInstrumentQuantity
}

func (p *PriceFormat56Choice) AddAmountPricePerAmount() *AmountPricePerAmount3 {
	p.AmountPricePerAmount = new(AmountPricePerAmount3)
	return p.AmountPricePerAmount
}

func (p *PriceFormat56Choice) SetIndexPoints(value string) {
	p.IndexPoints = (*RestrictedFINDecimalNumber)(&value)
}
