package model

// Choice of format for the date.
type DateCode18Choice struct {

	// Date expressed as an ISO 20022 code.
	Code *DateType5Code `xml:"Cd"`

	// Date expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (d *DateCode18Choice) SetCode(value string) {
	d.Code = (*DateType5Code)(&value)
}

func (d *DateCode18Choice) AddProprietary() *GenericIdentification30 {
	d.Proprietary = new(GenericIdentification30)
	return d.Proprietary
}
