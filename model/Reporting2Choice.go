package model

// Choice of format for the reporting type.
type Reporting2Choice struct {

	// Third party reporting information expressed as an ISO 20022 code.
	Code *Reporting2Code `xml:"Cd"`

	// Third party reporting information expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *Reporting2Choice) SetCode(value string) {
	r.Code = (*Reporting2Code)(&value)
}

func (r *Reporting2Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
