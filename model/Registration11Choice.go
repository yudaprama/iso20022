package model

// Choice of format for the registration information.
type Registration11Choice struct {

	// Registration information expressed as an ISO 20022 code.
	Code *Registration1Code `xml:"Cd"`

	// Registration information expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *Registration11Choice) SetCode(value string) {
	r.Code = (*Registration1Code)(&value)
}

func (r *Registration11Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
