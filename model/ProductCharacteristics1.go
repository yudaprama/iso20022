package model

// Identifies the characteristic of a product.
type ProductCharacteristics1 struct {

	// Specifies the type of product characteristic by means of a code.
	Type *ProductCharacteristics1Code `xml:"Tp"`

	// Specifies the characteristic of a product.
	Characteristics *Max35Text `xml:"Chrtcs"`
}

func (p *ProductCharacteristics1) SetType(value string) {
	p.Type = (*ProductCharacteristics1Code)(&value)
}

func (p *ProductCharacteristics1) SetCharacteristics(value string) {
	p.Characteristics = (*Max35Text)(&value)
}
