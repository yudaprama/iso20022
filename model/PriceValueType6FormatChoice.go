package model

// Choice of formats to  express  the value of a price.
type PriceValueType6FormatChoice struct {

	// Standard code to  specify the value of a price.
	Code *PriceValueType6Code `xml:"Cd"`

	// Proprietary code to  express  the value of a price.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (p *PriceValueType6FormatChoice) SetCode(value string) {
	p.Code = (*PriceValueType6Code)(&value)
}

func (p *PriceValueType6FormatChoice) AddProprietary() *GenericIdentification13 {
	p.Proprietary = new(GenericIdentification13)
	return p.Proprietary
}
