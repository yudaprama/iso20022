package model

// Choice of format for the registration information.
type Registration6Choice struct {

	// Registration information expressed as an ISO 20022 code.
	Code *Registration1Code `xml:"Cd"`

	// Registration information expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (r *Registration6Choice) SetCode(value string) {
	r.Code = (*Registration1Code)(&value)
}

func (r *Registration6Choice) AddProprietary() *GenericIdentification38 {
	r.Proprietary = new(GenericIdentification38)
	return r.Proprietary
}
