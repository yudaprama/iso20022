package model

// Provides detailed information on the suspension reason.
type MandateSuspensionReason1 struct {

	// Party that issues the suspension request.
	Originator *PartyIdentification43 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the suspension request.
	Reason *MandateSuspensionReason1Choice `xml:"Rsn"`

	// Further details on the suspension request reason.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (m *MandateSuspensionReason1) AddOriginator() *PartyIdentification43 {
	m.Originator = new(PartyIdentification43)
	return m.Originator
}

func (m *MandateSuspensionReason1) AddReason() *MandateSuspensionReason1Choice {
	m.Reason = new(MandateSuspensionReason1Choice)
	return m.Reason
}

func (m *MandateSuspensionReason1) AddAdditionalInformation(value string) {
	m.AdditionalInformation = append(m.AdditionalInformation, (*Max105Text)(&value))
}
