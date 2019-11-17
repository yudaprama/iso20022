package model

// Choice between a code or a proprietary code for a date code.
type DateCode21Choice struct {

	// Standard code to specify the type of date.
	Code *DateType7Code `xml:"Cd"`

	// Proprietary identification of the type of date.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (d *DateCode21Choice) SetCode(value string) {
	d.Code = (*DateType7Code)(&value)
}

func (d *DateCode21Choice) AddProprietary() *GenericIdentification30 {
	d.Proprietary = new(GenericIdentification30)
	return d.Proprietary
}
