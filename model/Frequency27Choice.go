package model

// Choice of format for a frequency, for example, a payment frequency.
type Frequency27Choice struct {

	// Frequency expressed as an ISO 20022 code.
	Code *EventFrequency3Code `xml:"Cd"`

	// Frequency expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (f *Frequency27Choice) SetCode(value string) {
	f.Code = (*EventFrequency3Code)(&value)
}

func (f *Frequency27Choice) AddProprietary() *GenericIdentification47 {
	f.Proprietary = new(GenericIdentification47)
	return f.Proprietary
}
