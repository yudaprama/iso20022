package model

// Amount of money for which goods or services are offered, sold, or bought.
type PriceInformation5 struct {

	// Specifies the type of price and information about the price.
	Type *TypeOfPrice4Choice `xml:"Tp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceRateOrAmountOrUnknownChoice `xml:"Val"`

	// Type of value in which the price is expressed.
	ValueType *YieldedOrValueType1Choice `xml:"ValTp"`

	// Place from which the price was obtained.
	SourceOfPrice *MarketIdentification6 `xml:"SrcOfPric,omitempty"`

	// Date on which the price is obtained. With an investment fund, this is as stated in the prospectus.
	QuotationDate *DateAndDateTimeChoice `xml:"QtnDt,omitempty"`
}

func (p *PriceInformation5) AddType() *TypeOfPrice4Choice {
	p.Type = new(TypeOfPrice4Choice)
	return p.Type
}

func (p *PriceInformation5) AddValue() *PriceRateOrAmountOrUnknownChoice {
	p.Value = new(PriceRateOrAmountOrUnknownChoice)
	return p.Value
}

func (p *PriceInformation5) AddValueType() *YieldedOrValueType1Choice {
	p.ValueType = new(YieldedOrValueType1Choice)
	return p.ValueType
}

func (p *PriceInformation5) AddSourceOfPrice() *MarketIdentification6 {
	p.SourceOfPrice = new(MarketIdentification6)
	return p.SourceOfPrice
}

func (p *PriceInformation5) AddQuotationDate() *DateAndDateTimeChoice {
	p.QuotationDate = new(DateAndDateTimeChoice)
	return p.QuotationDate
}
