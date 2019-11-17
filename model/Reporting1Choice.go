package model

// Choice of format for the reporting type.
type Reporting1Choice struct {

	// Third party reporting information expressed as an ISO 20022 code.
	Code *Reporting1Code `xml:"Cd"`

	// Third party reporting information expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *Reporting1Choice) SetCode(value string) {
	r.Code = (*Reporting1Code)(&value)
}

func (r *Reporting1Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
