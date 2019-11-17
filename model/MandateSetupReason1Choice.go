package model

// Specifies the reason for the setup of the mandate.
type MandateSetupReason1Choice struct {

	// Reason for the return, as published in an external reason code list.
	Code *ExternalMandateSetupReason1Code `xml:"Cd"`

	// Reason for the return, in a proprietary form.
	Proprietary *Max70Text `xml:"Prtry"`
}

func (m *MandateSetupReason1Choice) SetCode(value string) {
	m.Code = (*ExternalMandateSetupReason1Code)(&value)
}

func (m *MandateSetupReason1Choice) SetProprietary(value string) {
	m.Proprietary = (*Max70Text)(&value)
}
