package model

// Choice between a code or a proprietary code for a date code.
type DateCode2Choice struct {

	// Standard code to specify the type of date.
	Code *DateType6Code `xml:"Cd"`

	// Proprietary identification of the type of date.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (d *DateCode2Choice) SetCode(value string) {
	d.Code = (*DateType6Code)(&value)
}

func (d *DateCode2Choice) AddProprietary() *GenericIdentification20 {
	d.Proprietary = new(GenericIdentification20)
	return d.Proprietary
}
