package model

// Choice of format for a frequency, for example, the frequency of delivery of a statement.
type Frequency25Choice struct {

	// Frequency expressed as an ISO 20022 code.
	Code *EventFrequency4Code `xml:"Cd"`

	// Frequency expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (f *Frequency25Choice) SetCode(value string) {
	f.Code = (*EventFrequency4Code)(&value)
}

func (f *Frequency25Choice) AddProprietary() *GenericIdentification30 {
	f.Proprietary = new(GenericIdentification30)
	return f.Proprietary
}
