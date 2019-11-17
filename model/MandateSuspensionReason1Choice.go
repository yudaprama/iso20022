package model

// Specifies the reason for the suspension request of a mandate.
type MandateSuspensionReason1Choice struct {

	// Reason, as published in an external reason code list.
	Code *ExternalMandateSuspensionReason1Code `xml:"Cd"`

	// Reason, in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (m *MandateSuspensionReason1Choice) SetCode(value string) {
	m.Code = (*ExternalMandateSuspensionReason1Code)(&value)
}

func (m *MandateSuspensionReason1Choice) SetProprietary(value string) {
	m.Proprietary = (*Max35Text)(&value)
}
