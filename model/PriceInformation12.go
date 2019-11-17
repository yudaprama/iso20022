package model

// Amount of money for which goods or services are offered, sold, or bought.
type PriceInformation12 struct {

	// Specifies the type of price and information about the price.
	Type *TypeOfPrice28Choice `xml:"Tp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceRateOrAmountOrUnknownChoice `xml:"Val"`

	// Type of value in which the price is expressed.
	ValueType *YieldedOrValueType1Choice `xml:"ValTp"`

	// Place from which the price was obtained.
	SourceOfPrice *MarketIdentification89 `xml:"SrcOfPric,omitempty"`

	// Date on which the price is obtained. With an investment fund, this is as stated in the prospectus.
	QuotationDate *DateAndDateTimeChoice `xml:"QtnDt,omitempty"`
}

func (p *PriceInformation12) AddType() *TypeOfPrice28Choice {
	p.Type = new(TypeOfPrice28Choice)
	return p.Type
}

func (p *PriceInformation12) AddValue() *PriceRateOrAmountOrUnknownChoice {
	p.Value = new(PriceRateOrAmountOrUnknownChoice)
	return p.Value
}

func (p *PriceInformation12) AddValueType() *YieldedOrValueType1Choice {
	p.ValueType = new(YieldedOrValueType1Choice)
	return p.ValueType
}

func (p *PriceInformation12) AddSourceOfPrice() *MarketIdentification89 {
	p.SourceOfPrice = new(MarketIdentification89)
	return p.SourceOfPrice
}

func (p *PriceInformation12) AddQuotationDate() *DateAndDateTimeChoice {
	p.QuotationDate = new(DateAndDateTimeChoice)
	return p.QuotationDate
}
