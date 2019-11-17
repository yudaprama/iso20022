package model

// Choice between a code or a proprietary code for a date code.
type DateCode20Choice struct {

	// Standard code to indicate the date is unknown.
	Code *DateType1Code `xml:"Cd"`

	// Proprietary identification of the type of date.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (d *DateCode20Choice) SetCode(value string) {
	d.Code = (*DateType1Code)(&value)
}

func (d *DateCode20Choice) AddProprietary() *GenericIdentification30 {
	d.Proprietary = new(GenericIdentification30)
	return d.Proprietary
}
