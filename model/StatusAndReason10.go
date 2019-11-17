package model

// Provides details related to the status of the order.
type StatusAndReason10 struct {

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus7Choice `xml:"AffirmSts"`

	// Specifies the reason why the instruction has an unaffirmed status.
	UnaffirmedReason *UnaffirmedReason2Choice `xml:"UaffrmdRsn,omitempty"`

	// Provides additional information about the reason in narrative form.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (s *StatusAndReason10) AddAffirmationStatus() *AffirmationStatus7Choice {
	s.AffirmationStatus = new(AffirmationStatus7Choice)
	return s.AffirmationStatus
}

func (s *StatusAndReason10) AddUnaffirmedReason() *UnaffirmedReason2Choice {
	s.UnaffirmedReason = new(UnaffirmedReason2Choice)
	return s.UnaffirmedReason
}

func (s *StatusAndReason10) SetAdditionalReasonInformation(value string) {
	s.AdditionalReasonInformation = (*Max210Text)(&value)
}
