package model

// Choice of format for the registration information.
type Registration9Choice struct {

	// Registration information expressed as an ISO 20022 code.
	Code *Registration1Code `xml:"Cd"`

	// Registration information expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *Registration9Choice) SetCode(value string) {
	r.Code = (*Registration1Code)(&value)
}

func (r *Registration9Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
