package model

// Choice of format for a frequency, for example, a payment frequency.
type Frequency3Choice struct {

	// Frequency expressed as an ISO 20022 code.
	Code *EventFrequency3Code `xml:"Cd"`

	// Frequency expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (f *Frequency3Choice) SetCode(value string) {
	f.Code = (*EventFrequency3Code)(&value)
}

func (f *Frequency3Choice) AddProprietary() *GenericIdentification20 {
	f.Proprietary = new(GenericIdentification20)
	return f.Proprietary
}
