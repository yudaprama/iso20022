package model

// Choice between a code or a proprietary code for a date code.
type DateCode27Choice struct {

	// Standard code to indicate the date is unknown.
	Code *DateType1Code `xml:"Cd"`

	// Proprietary identification of the type of date.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (d *DateCode27Choice) SetCode(value string) {
	d.Code = (*DateType1Code)(&value)
}

func (d *DateCode27Choice) AddProprietary() *GenericIdentification47 {
	d.Proprietary = new(GenericIdentification47)
	return d.Proprietary
}
