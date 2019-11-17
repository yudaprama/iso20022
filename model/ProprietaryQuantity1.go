package model

// Set of elements to identify a proprietary quantity.
type ProprietaryQuantity1 struct {

	// Identifies the type of proprietary quantity reported.
	Type *Max35Text `xml:"Tp"`

	// Provides the proprietary quantity in free format.
	Quantity *Max35Text `xml:"Qty"`
}

func (p *ProprietaryQuantity1) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryQuantity1) SetQuantity(value string) {
	p.Quantity = (*Max35Text)(&value)
}
