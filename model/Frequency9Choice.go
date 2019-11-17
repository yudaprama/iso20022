package model

// Choice of format for a frequency, for example, the frequency of delivery of a statement.
type Frequency9Choice struct {

	// Frequency expressed as an ISO 20022 code.
	Code *EventFrequency7Code `xml:"Cd"`

	// Frequency expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (f *Frequency9Choice) SetCode(value string) {
	f.Code = (*EventFrequency7Code)(&value)
}

func (f *Frequency9Choice) AddProprietary() *GenericIdentification20 {
	f.Proprietary = new(GenericIdentification20)
	return f.Proprietary
}
