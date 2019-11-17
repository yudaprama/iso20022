package model

// Choice of format for the registration information.
type Registration3Choice struct {

	// Registration information expressed as an ISO 20022 code.
	Code *Registration2Code `xml:"Cd"`

	// Registration information expressed as a proprietary code.
	Proprietary *GenericIdentification19 `xml:"Prtry"`
}

func (r *Registration3Choice) SetCode(value string) {
	r.Code = (*Registration2Code)(&value)
}

func (r *Registration3Choice) AddProprietary() *GenericIdentification19 {
	r.Proprietary = new(GenericIdentification19)
	return r.Proprietary
}
