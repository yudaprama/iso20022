package model

// Choice of formats to  express  the value of a price.
type PriceValueType5FormatChoice struct {

	// Standard code to  specify the value of a price.
	Code *PriceValueType5Code `xml:"Cd"`

	// Proprietary code to  express  the value of a price.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (p *PriceValueType5FormatChoice) SetCode(value string) {
	p.Code = (*PriceValueType5Code)(&value)
}

func (p *PriceValueType5FormatChoice) AddProprietary() *GenericIdentification13 {
	p.Proprietary = new(GenericIdentification13)
	return p.Proprietary
}
