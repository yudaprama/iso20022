package model

// Choice between a code or a proprietary code for a date code.
type DateCode26Choice struct {

	// Standard code to specify the type of date.
	Code *DateType7Code `xml:"Cd"`

	// Proprietary identification of the type of date.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (d *DateCode26Choice) SetCode(value string) {
	d.Code = (*DateType7Code)(&value)
}

func (d *DateCode26Choice) AddProprietary() *GenericIdentification47 {
	d.Proprietary = new(GenericIdentification47)
	return d.Proprietary
}
