package model

// Choice of format for a frequency, for example, a payment frequency.
type Frequency7Choice struct {

	// Frequency expressed as an ISO 20022 code.
	Code *EventFrequency3Code `xml:"Cd"`

	// Frequency expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (f *Frequency7Choice) SetCode(value string) {
	f.Code = (*EventFrequency3Code)(&value)
}

func (f *Frequency7Choice) AddProprietary() *GenericIdentification38 {
	f.Proprietary = new(GenericIdentification38)
	return f.Proprietary
}
