package model

// Choice of format for the registration information.
type Registration10Choice struct {

	// Registration information expressed as an ISO 20022 code.
	Code *Registration2Code `xml:"Cd"`

	// Registration information expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *Registration10Choice) SetCode(value string) {
	r.Code = (*Registration2Code)(&value)
}

func (r *Registration10Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
