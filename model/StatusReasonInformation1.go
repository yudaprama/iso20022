package model

// Further information on the status reason of the transaction.
type StatusReasonInformation1 struct {

	// Party issuing the status.
	StatusOriginator *PartyIdentification8 `xml:"StsOrgtr,omitempty"`

	// Specifies the reason for the status report.
	StatusReason *StatusReason1Choice `xml:"StsRsn,omitempty"`

	// Further details on the status reason.
	//
	// Usage: Additional information can be used for several purposes, e.g. to report repaired information, or to further detail the status reason.
	AdditionalStatusReasonInformation []*Max105Text `xml:"AddtlStsRsnInf,omitempty"`
}

func (s *StatusReasonInformation1) AddStatusOriginator() *PartyIdentification8 {
	s.StatusOriginator = new(PartyIdentification8)
	return s.StatusOriginator
}

func (s *StatusReasonInformation1) AddStatusReason() *StatusReason1Choice {
	s.StatusReason = new(StatusReason1Choice)
	return s.StatusReason
}

func (s *StatusReasonInformation1) AddAdditionalStatusReasonInformation(value string) {
	s.AdditionalStatusReasonInformation = append(s.AdditionalStatusReasonInformation, (*Max105Text)(&value))
}
