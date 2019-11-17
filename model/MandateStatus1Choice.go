package model

// Sets of elements to identify the status of a mandate.
type MandateStatus1Choice struct {

	// Status of the mandate, in a coded form as published in an external list.
	Code *ExternalMandateStatus1Code `xml:"Cd"`

	// Name of the identification scheme, in a free text form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (m *MandateStatus1Choice) SetCode(value string) {
	m.Code = (*ExternalMandateStatus1Code)(&value)
}

func (m *MandateStatus1Choice) SetProprietary(value string) {
	m.Proprietary = (*Max35Text)(&value)
}
