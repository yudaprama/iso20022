package model

// Amount of money for which goods or services are offered, sold, or bought.
type PriceInformation14 struct {

	// Specifies the type of price and information about the price.
	Type *TypeOfPrice33Choice `xml:"Tp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceRateOrAmountOrUnknown1Choice `xml:"Val"`

	// Type of value in which the price is expressed.
	ValueType *YieldedOrValueType1Choice `xml:"ValTp"`

	// Place from which the price was obtained.
	SourceOfPrice *MarketIdentification91 `xml:"SrcOfPric,omitempty"`

	// Date on which the price is obtained. With an investment fund, this is as stated in the prospectus.
	QuotationDate *DateAndDateTimeChoice `xml:"QtnDt,omitempty"`
}

func (p *PriceInformation14) AddType() *TypeOfPrice33Choice {
	p.Type = new(TypeOfPrice33Choice)
	return p.Type
}

func (p *PriceInformation14) AddValue() *PriceRateOrAmountOrUnknown1Choice {
	p.Value = new(PriceRateOrAmountOrUnknown1Choice)
	return p.Value
}

func (p *PriceInformation14) AddValueType() *YieldedOrValueType1Choice {
	p.ValueType = new(YieldedOrValueType1Choice)
	return p.ValueType
}

func (p *PriceInformation14) AddSourceOfPrice() *MarketIdentification91 {
	p.SourceOfPrice = new(MarketIdentification91)
	return p.SourceOfPrice
}

func (p *PriceInformation14) AddQuotationDate() *DateAndDateTimeChoice {
	p.QuotationDate = new(DateAndDateTimeChoice)
	return p.QuotationDate
}
