package model

// Price expressed as a percentage price.
type PercentagePrice1 struct {

	// Specifies the type of percentage price.
	PercentagePriceType *PriceRateType3Code `xml:"PctgPricTp"`

	// Specifies the value of price.
	PriceValue *PercentageRate `xml:"PricVal"`
}

func (p *PercentagePrice1) SetPercentagePriceType(value string) {
	p.PercentagePriceType = (*PriceRateType3Code)(&value)
}

func (p *PercentagePrice1) SetPriceValue(value string) {
	p.PriceValue = (*PercentageRate)(&value)
}
