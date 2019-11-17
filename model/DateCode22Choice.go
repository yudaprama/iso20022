package model

// Choice between a code or a proprietary code for a date code.
type DateCode22Choice struct {

	// Standard code to specify the type of date.
	Code *DateType8Code `xml:"Cd"`

	// Proprietary identification of the type of date.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (d *DateCode22Choice) SetCode(value string) {
	d.Code = (*DateType8Code)(&value)
}

func (d *DateCode22Choice) AddProprietary() *GenericIdentification47 {
	d.Proprietary = new(GenericIdentification47)
	return d.Proprietary
}
