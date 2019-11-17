package model

// Choice of format for the reporting type.
type Reporting5Choice struct {

	// Third party reporting information expressed as an ISO 20022 code.
	Code *Reporting2Code `xml:"Cd"`

	// Third party reporting information expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (r *Reporting5Choice) SetCode(value string) {
	r.Code = (*Reporting2Code)(&value)
}

func (r *Reporting5Choice) AddProprietary() *GenericIdentification38 {
	r.Proprietary = new(GenericIdentification38)
	return r.Proprietary
}
