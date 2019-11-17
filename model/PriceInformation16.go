package model

// Amount of money for which goods or services are offered, sold, or bought.
type PriceInformation16 struct {

	// Specifies the type of price and information about the price.
	Type *TypeOfPrice45Choice `xml:"Tp"`

	// Type of value in which the price is expressed.
	ValueType *YieldedOrValueType1Choice `xml:"ValTp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceRateOrAmountOrUnknown1Choice `xml:"Val"`

	// Place from which the price was obtained.
	SourceOfPrice *MarketIdentification91 `xml:"SrcOfPric,omitempty"`

	// Date on which the price is obtained. With an investment fund, this is as stated in the prospectus.
	QuotationDate *DateAndDateTimeChoice `xml:"QtnDt,omitempty"`
}

func (p *PriceInformation16) AddType() *TypeOfPrice45Choice {
	p.Type = new(TypeOfPrice45Choice)
	return p.Type
}

func (p *PriceInformation16) AddValueType() *YieldedOrValueType1Choice {
	p.ValueType = new(YieldedOrValueType1Choice)
	return p.ValueType
}

func (p *PriceInformation16) AddValue() *PriceRateOrAmountOrUnknown1Choice {
	p.Value = new(PriceRateOrAmountOrUnknown1Choice)
	return p.Value
}

func (p *PriceInformation16) AddSourceOfPrice() *MarketIdentification91 {
	p.SourceOfPrice = new(MarketIdentification91)
	return p.SourceOfPrice
}

func (p *PriceInformation16) AddQuotationDate() *DateAndDateTimeChoice {
	p.QuotationDate = new(DateAndDateTimeChoice)
	return p.QuotationDate
}
