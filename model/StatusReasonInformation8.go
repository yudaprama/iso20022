package model

// Set of elements used to provide information on the status reason of the transaction.
type StatusReasonInformation8 struct {

	// Party that issues the status.
	Originator *PartyIdentification32 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the status report.
	Reason *StatusReason6Choice `xml:"Rsn,omitempty"`

	// Further details on the status reason.
	//
	// Usage: Additional information can be used for several purposes such as the reporting of repaired information.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (s *StatusReasonInformation8) AddOriginator() *PartyIdentification32 {
	s.Originator = new(PartyIdentification32)
	return s.Originator
}

func (s *StatusReasonInformation8) AddReason() *StatusReason6Choice {
	s.Reason = new(StatusReason6Choice)
	return s.Reason
}

func (s *StatusReasonInformation8) AddAdditionalInformation(value string) {
	s.AdditionalInformation = append(s.AdditionalInformation, (*Max105Text)(&value))
}
