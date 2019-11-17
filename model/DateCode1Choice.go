package model

// Choice of format for the date.
type DateCode1Choice struct {

	// Date expressed as an ISO 20022 code.
	Code *DateType5Code `xml:"Cd"`

	// Date expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (d *DateCode1Choice) SetCode(value string) {
	d.Code = (*DateType5Code)(&value)
}

func (d *DateCode1Choice) AddProprietary() *GenericIdentification20 {
	d.Proprietary = new(GenericIdentification20)
	return d.Proprietary
}
