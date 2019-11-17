package model

// Specifies the type of identification of a person.
type PersonIdentificationType3Choice struct {

	// Person identification expressed as a code.
	Code *PersonIdentificationType3Code `xml:"Cd"`

	// Person identification expressed as a proprietary code.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (p *PersonIdentificationType3Choice) SetCode(value string) {
	p.Code = (*PersonIdentificationType3Code)(&value)
}

func (p *PersonIdentificationType3Choice) AddProprietary() *GenericIdentification13 {
	p.Proprietary = new(GenericIdentification13)
	return p.Proprietary
}
