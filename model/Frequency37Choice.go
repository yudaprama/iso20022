package model

// Choice of format for a frequency, for example, the frequency of payment.
type Frequency37Choice struct {

	// Frequency expressed as an ISO 20022 code.
	Code *Frequency10Code `xml:"Cd"`

	// Frequency expressed as a proprietary code.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (f *Frequency37Choice) SetCode(value string) {
	f.Code = (*Frequency10Code)(&value)
}

func (f *Frequency37Choice) SetProprietary(value string) {
	f.Proprietary = (*Max35Text)(&value)
}
