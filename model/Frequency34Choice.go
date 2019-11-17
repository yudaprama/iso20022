package model

// Choice of format for a frequency, for example, the frequency of delivery of a statement.
type Frequency34Choice struct {

	// Frequency expressed as an ISO 20022 code.
	Code *EventFrequency7Code `xml:"Cd"`

	// Frequency expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (f *Frequency34Choice) SetCode(value string) {
	f.Code = (*EventFrequency7Code)(&value)
}

func (f *Frequency34Choice) AddProprietary() *GenericIdentification47 {
	f.Proprietary = new(GenericIdentification47)
	return f.Proprietary
}
