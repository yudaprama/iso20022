package model

// Choice of format for the reporting type.
type Reporting6Choice struct {

	// Third party reporting information expressed as an ISO 20022 code.
	Code *Reporting2Code `xml:"Cd"`

	// Third party reporting information expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *Reporting6Choice) SetCode(value string) {
	r.Code = (*Reporting2Code)(&value)
}

func (r *Reporting6Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
