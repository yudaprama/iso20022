package model

// Information used to identify a product.
type ProductIdentifier2 struct {

	// Specifies the type of product identifier by means of a code.
	Type *ProductIdentifier2Code `xml:"Tp"`

	// Specifies the product identifier.
	Identifier *Max35Text `xml:"Idr"`
}

func (p *ProductIdentifier2) SetType(value string) {
	p.Type = (*ProductIdentifier2Code)(&value)
}

func (p *ProductIdentifier2) SetIdentifier(value string) {
	p.Identifier = (*Max35Text)(&value)
}
