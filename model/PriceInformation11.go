package model

// Amount of money for which goods or services are offered, sold, or bought.
type PriceInformation11 struct {

	// Value of the price, for instance, as a currency and value.
	Value *Price4 `xml:"Val"`

	// Date on which the price is obtained. With an investment fund, this is as stated in the prospectus.
	QuotationDate *DateAndDateTime1Choice `xml:"QtnDt,omitempty"`

	// Period during which the price of a security is determined (For  outturn securities).
	PriceCalculationPeriod *DateTimePeriodChoice `xml:"PricClctnPrd,omitempty"`

	// Place from which the price was obtained.
	SourceOfPrice *MarketIdentification79 `xml:"SrcOfPric,omitempty"`
}

func (p *PriceInformation11) AddValue() *Price4 {
	p.Value = new(Price4)
	return p.Value
}

func (p *PriceInformation11) AddQuotationDate() *DateAndDateTime1Choice {
	p.QuotationDate = new(DateAndDateTime1Choice)
	return p.QuotationDate
}

func (p *PriceInformation11) AddPriceCalculationPeriod() *DateTimePeriodChoice {
	p.PriceCalculationPeriod = new(DateTimePeriodChoice)
	return p.PriceCalculationPeriod
}

func (p *PriceInformation11) AddSourceOfPrice() *MarketIdentification79 {
	p.SourceOfPrice = new(MarketIdentification79)
	return p.SourceOfPrice
}
