package model

// Amount of money for which goods or services are offered, sold, or bought.
type PriceInformation13 struct {

	// Specifies the type of price and information about the price.
	Type *TypeOfPrice30Choice `xml:"Tp"`

	// Type of value in which the price is expressed.
	ValueType *YieldedOrValueType1Choice `xml:"ValTp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceRateOrAmountOrUnknownChoice `xml:"Val"`

	// Place from which the price was obtained.
	SourceOfPrice *MarketIdentification89 `xml:"SrcOfPric,omitempty"`

	// Date on which the price is obtained. With an investment fund, this is as stated in the prospectus.
	QuotationDate *DateAndDateTimeChoice `xml:"QtnDt,omitempty"`
}

func (p *PriceInformation13) AddType() *TypeOfPrice30Choice {
	p.Type = new(TypeOfPrice30Choice)
	return p.Type
}

func (p *PriceInformation13) AddValueType() *YieldedOrValueType1Choice {
	p.ValueType = new(YieldedOrValueType1Choice)
	return p.ValueType
}

func (p *PriceInformation13) AddValue() *PriceRateOrAmountOrUnknownChoice {
	p.Value = new(PriceRateOrAmountOrUnknownChoice)
	return p.Value
}

func (p *PriceInformation13) AddSourceOfPrice() *MarketIdentification89 {
	p.SourceOfPrice = new(MarketIdentification89)
	return p.SourceOfPrice
}

func (p *PriceInformation13) AddQuotationDate() *DateAndDateTimeChoice {
	p.QuotationDate = new(DateAndDateTimeChoice)
	return p.QuotationDate
}
