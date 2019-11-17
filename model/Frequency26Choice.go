package model

// Choice of format for a frequency, for example, the frequency of delivery of a statement.
type Frequency26Choice struct {

	// Frequency expressed as an ISO 20022 code.
	Code *EventFrequency4Code `xml:"Cd"`

	// Frequency expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (f *Frequency26Choice) SetCode(value string) {
	f.Code = (*EventFrequency4Code)(&value)
}

func (f *Frequency26Choice) AddProprietary() *GenericIdentification47 {
	f.Proprietary = new(GenericIdentification47)
	return f.Proprietary
}
