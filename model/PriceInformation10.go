package model

// Amount of money for which goods or services are offered, sold, or bought.
type PriceInformation10 struct {

	// Current value of the price, eg, as a currency and value.
	CurrentPrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"CurPric"`

	// Specifies the type of price and information about the price.
	Type *TypeOfPrice27Choice `xml:"Tp"`

	// Previous value of the price, eg, as a currency and value.
	PreviousPrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"PrvsPric,omitempty"`

	// Difference or change between the previous price value and the current price value.
	AmountOfChange *PriceValueAndRate4 `xml:"AmtOfChng,omitempty"`
}

func (p *PriceInformation10) SetCurrentPrice(value, currency string) {
	p.CurrentPrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (p *PriceInformation10) AddType() *TypeOfPrice27Choice {
	p.Type = new(TypeOfPrice27Choice)
	return p.Type
}

func (p *PriceInformation10) SetPreviousPrice(value, currency string) {
	p.PreviousPrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (p *PriceInformation10) AddAmountOfChange() *PriceValueAndRate4 {
	p.AmountOfChange = new(PriceValueAndRate4)
	return p.AmountOfChange
}
