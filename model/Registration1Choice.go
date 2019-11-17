package model

// Choice of format for the registration information.
type Registration1Choice struct {

	// Registration information expressed as an ISO 20022 code.
	Code *Registration1Code `xml:"Cd"`

	// Registration information expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *Registration1Choice) SetCode(value string) {
	r.Code = (*Registration1Code)(&value)
}

func (r *Registration1Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
