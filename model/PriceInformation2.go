package model

// Amount of money for which goods or services are offered, sold, or bought.
type PriceInformation2 struct {

	// Value of the price, eg, as a currency and value.
	Value *PriceRateOrAmountOrUnknownChoice `xml:"Val"`

	// Type of value in which the price is expressed.
	ValueType *PriceValueType2Code `xml:"ValTp,omitempty"`

	// Type and information about a price.
	Type *TypeOfPrice11Code `xml:"Tp"`

	// Type and information about a price.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Place from which the price was obtained.
	SourceOfPrice *PriceSourceFormatChoice `xml:"SrcOfPric,omitempty"`

	// Date on which the price is obtained. With an investment fund, this is as stated in the prospectus.
	QuotationDate *DateAndDateTimeChoice `xml:"QtnDt,omitempty"`

	// Indicates whether the price is expressed as a yield. The absence of Yielded means it is not applicable.
	Yielded *YesNoIndicator `xml:"Yldd,omitempty"`
}

func (p *PriceInformation2) AddValue() *PriceRateOrAmountOrUnknownChoice {
	p.Value = new(PriceRateOrAmountOrUnknownChoice)
	return p.Value
}

func (p *PriceInformation2) SetValueType(value string) {
	p.ValueType = (*PriceValueType2Code)(&value)
}

func (p *PriceInformation2) SetType(value string) {
	p.Type = (*TypeOfPrice11Code)(&value)
}

func (p *PriceInformation2) SetExtendedType(value string) {
	p.ExtendedType = (*Extended350Code)(&value)
}

func (p *PriceInformation2) AddSourceOfPrice() *PriceSourceFormatChoice {
	p.SourceOfPrice = new(PriceSourceFormatChoice)
	return p.SourceOfPrice
}

func (p *PriceInformation2) AddQuotationDate() *DateAndDateTimeChoice {
	p.QuotationDate = new(DateAndDateTimeChoice)
	return p.QuotationDate
}

func (p *PriceInformation2) SetYielded(value string) {
	p.Yielded = (*YesNoIndicator)(&value)
}
