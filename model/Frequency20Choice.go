package model

// Choice of formats for the frequency.
type Frequency20Choice struct {

	// Frequency expressed as a code.
	Code *EventFrequency8Code `xml:"Cd"`

	// Frequency expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (f *Frequency20Choice) SetCode(value string) {
	f.Code = (*EventFrequency8Code)(&value)
}

func (f *Frequency20Choice) AddProprietary() *GenericIdentification47 {
	f.Proprietary = new(GenericIdentification47)
	return f.Proprietary
}
