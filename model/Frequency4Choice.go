package model

// Choice of format for a frequency, for example, the frequency of delivery of a statement.
type Frequency4Choice struct {

	// Frequency expressed as an ISO 20022 code.
	Code *EventFrequency4Code `xml:"Cd"`

	// Frequency expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (f *Frequency4Choice) SetCode(value string) {
	f.Code = (*EventFrequency4Code)(&value)
}

func (f *Frequency4Choice) AddProprietary() *GenericIdentification20 {
	f.Proprietary = new(GenericIdentification20)
	return f.Proprietary
}
