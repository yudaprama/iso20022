package model

// Choice of format for the frequency.
type Frequency8Choice struct {

	// Frequency expressed in coded form.
	Code *EventFrequency1Code `xml:"Cd"`

	// Frequency expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (f *Frequency8Choice) SetCode(value string) {
	f.Code = (*EventFrequency1Code)(&value)
}

func (f *Frequency8Choice) AddProprietary() *GenericIdentification30 {
	f.Proprietary = new(GenericIdentification30)
	return f.Proprietary
}
