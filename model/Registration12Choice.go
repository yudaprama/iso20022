package model

// Choice of format for the registration information.
type Registration12Choice struct {

	// Registration information expressed as an ISO 20022 code.
	Code *Registration2Code `xml:"Cd"`

	// Registration information expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *Registration12Choice) SetCode(value string) {
	r.Code = (*Registration2Code)(&value)
}

func (r *Registration12Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
