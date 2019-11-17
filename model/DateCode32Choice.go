package model

// Choice of format for the date.
type DateCode32Choice struct {

	// Date expressed as an ISO 20022 code.
	Code *DateType5Code `xml:"Cd"`

	// Date expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (d *DateCode32Choice) SetCode(value string) {
	d.Code = (*DateType5Code)(&value)
}

func (d *DateCode32Choice) AddProprietary() *GenericIdentification47 {
	d.Proprietary = new(GenericIdentification47)
	return d.Proprietary
}
