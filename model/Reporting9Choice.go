package model

// Choice of format for the reporting type.
type Reporting9Choice struct {

	// Third party reporting information expressed as an ISO 20022 code.
	Code *Reporting2Code `xml:"Cd"`

	// Third party reporting information expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *Reporting9Choice) SetCode(value string) {
	r.Code = (*Reporting2Code)(&value)
}

func (r *Reporting9Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
