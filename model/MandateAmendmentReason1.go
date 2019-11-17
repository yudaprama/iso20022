package model

// Provides further details on the reason of the amendment of the mandate.
type MandateAmendmentReason1 struct {

	// Party that issues the amendment request.
	Originator *PartyIdentification43 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the amendment request.
	Reason *MandateReason1Choice `xml:"Rsn"`

	// Further details on the amendment request reason.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (m *MandateAmendmentReason1) AddOriginator() *PartyIdentification43 {
	m.Originator = new(PartyIdentification43)
	return m.Originator
}

func (m *MandateAmendmentReason1) AddReason() *MandateReason1Choice {
	m.Reason = new(MandateReason1Choice)
	return m.Reason
}

func (m *MandateAmendmentReason1) AddAdditionalInformation(value string) {
	m.AdditionalInformation = append(m.AdditionalInformation, (*Max105Text)(&value))
}
