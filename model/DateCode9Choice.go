package model

// Choice between a code or a proprietary code for a date code.
type DateCode9Choice struct {

	// Standard code to specify the type of date.
	Code *DateType2Code `xml:"Cd"`

	// Proprietary identification of the type of date.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (d *DateCode9Choice) SetCode(value string) {
	d.Code = (*DateType2Code)(&value)
}

func (d *DateCode9Choice) AddProprietary() *GenericIdentification30 {
	d.Proprietary = new(GenericIdentification30)
	return d.Proprietary
}
