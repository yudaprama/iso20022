package model

// Set of elements used to identify a proprietary price.
type ProprietaryPrice2 struct {

	// Specifies the type of price.
	Type *Max35Text `xml:"Tp"`

	// Proprietary price specification related to the underlying transaction.
	Price *ActiveOrHistoricCurrencyAndAmount `xml:"Pric"`
}

func (p *ProprietaryPrice2) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryPrice2) SetPrice(value, currency string) {
	p.Price = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}
