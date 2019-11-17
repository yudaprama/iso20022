package model

// Choice between a code or a proprietary code for a date code.
type DateCode3Choice struct {

	// Standard code to indicate the date is unknown.
	Code *DateType1Code `xml:"Cd"`

	// Proprietary identification of the type of date.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (d *DateCode3Choice) SetCode(value string) {
	d.Code = (*DateType1Code)(&value)
}

func (d *DateCode3Choice) AddProprietary() *GenericIdentification20 {
	d.Proprietary = new(GenericIdentification20)
	return d.Proprietary
}
