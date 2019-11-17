package model

// Specifies the high level purpose of the instruction based on a set of pre-defined categories.
// Usage: This is used by the initiating party to provide information concerning the processing of the payment. It is likely to trigger special processing by any of the agents involved in the payment chain.
type MandateClassification1Choice struct {

	// Category purpose, as published in an external category purpose code list.
	Code *MandateClassification1Code `xml:"Cd"`

	// Category purpose, in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (m *MandateClassification1Choice) SetCode(value string) {
	m.Code = (*MandateClassification1Code)(&value)
}

func (m *MandateClassification1Choice) SetProprietary(value string) {
	m.Proprietary = (*Max35Text)(&value)
}
