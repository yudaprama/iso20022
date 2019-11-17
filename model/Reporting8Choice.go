package model

// Choice of format for the reporting type.
type Reporting8Choice struct {

	// Third party reporting information expressed as an ISO 20022 code.
	Code *Reporting1Code `xml:"Cd"`

	// Third party reporting information expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *Reporting8Choice) SetCode(value string) {
	r.Code = (*Reporting1Code)(&value)
}

func (r *Reporting8Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
