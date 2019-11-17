package model

// Choice between a code or a data source scheme to determine a date format.
type Date2Choice struct {

	// Date is defined using a code.
	Code *DateType2Code `xml:"Cd"`

	// Date is determined using a data source scheme and a code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (d *Date2Choice) SetCode(value string) {
	d.Code = (*DateType2Code)(&value)
}

func (d *Date2Choice) AddProprietary() *GenericIdentification38 {
	d.Proprietary = new(GenericIdentification38)
	return d.Proprietary
}
