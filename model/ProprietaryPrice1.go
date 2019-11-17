package model

// Set of elements to identify a proprietary price.
type ProprietaryPrice1 struct {

	// Identifies the type of price reported.
	Type *Max35Text `xml:"Tp"`

	// Proprietary price specification related to the underlying transaction.
	Price *CurrencyAndAmount `xml:"Pric"`
}

func (p *ProprietaryPrice1) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryPrice1) SetPrice(value, currency string) {
	p.Price = NewCurrencyAndAmount(value, currency)
}
