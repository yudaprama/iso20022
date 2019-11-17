package model

// Reason for a suspended status.
type SuspendedStatusReason4 struct {

	// Reason for the conditionally accepted status expressed as a code.
	Reason *SuspendedStatusReason5Choice `xml:"Rsn"`

	// Additional information about the suspended reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (s *SuspendedStatusReason4) AddReason() *SuspendedStatusReason5Choice {
	s.Reason = new(SuspendedStatusReason5Choice)
	return s.Reason
}

func (s *SuspendedStatusReason4) SetAdditionalInformation(value string) {
	s.AdditionalInformation = (*Max350Text)(&value)
}
