package model

// Choice of format for a frequency, for example, a payment frequency.
type Frequency23Choice struct {

	// Frequency expressed as an ISO 20022 code.
	Code *EventFrequency3Code `xml:"Cd"`

	// Frequency expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (f *Frequency23Choice) SetCode(value string) {
	f.Code = (*EventFrequency3Code)(&value)
}

func (f *Frequency23Choice) AddProprietary() *GenericIdentification30 {
	f.Proprietary = new(GenericIdentification30)
	return f.Proprietary
}
