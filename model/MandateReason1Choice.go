package model

// Specifies the reason for the request or status of a mandate.
type MandateReason1Choice struct {

	// Reason, as published in an external reason code list.
	Code *ExternalMandateReason1Code `xml:"Cd"`

	// Reason, in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (m *MandateReason1Choice) SetCode(value string) {
	m.Code = (*ExternalMandateReason1Code)(&value)
}

func (m *MandateReason1Choice) SetProprietary(value string) {
	m.Proprietary = (*Max35Text)(&value)
}
