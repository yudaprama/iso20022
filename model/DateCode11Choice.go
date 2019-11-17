package model

// Choice between a code or a proprietary code for a date code.
type DateCode11Choice struct {

	// Standard code to specify the type of date.
	Code *DateType8Code `xml:"Cd"`

	// Proprietary identification of the type of date.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (d *DateCode11Choice) SetCode(value string) {
	d.Code = (*DateType8Code)(&value)
}

func (d *DateCode11Choice) AddProprietary() *GenericIdentification20 {
	d.Proprietary = new(GenericIdentification20)
	return d.Proprietary
}
