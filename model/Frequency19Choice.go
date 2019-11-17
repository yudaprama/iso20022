package model

// Choice of formats for the frequency.
type Frequency19Choice struct {

	// Frequency expressed as a code.
	Code *EventFrequency1Code `xml:"Cd"`

	// Frequency expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (f *Frequency19Choice) SetCode(value string) {
	f.Code = (*EventFrequency1Code)(&value)
}

func (f *Frequency19Choice) AddProprietary() *GenericIdentification47 {
	f.Proprietary = new(GenericIdentification47)
	return f.Proprietary
}
