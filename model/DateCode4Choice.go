package model

// Choice between a code or a proprietary code for a date code.
type DateCode4Choice struct {

	// Standard code to specify the type of date.
	Code *DateType7Code `xml:"Cd"`

	// Proprietary identification of the type of date.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (d *DateCode4Choice) SetCode(value string) {
	d.Code = (*DateType7Code)(&value)
}

func (d *DateCode4Choice) AddProprietary() *GenericIdentification20 {
	d.Proprietary = new(GenericIdentification20)
	return d.Proprietary
}
