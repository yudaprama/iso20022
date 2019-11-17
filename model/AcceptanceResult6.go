package model

// Set of elements used to provide detailed information on the acceptance result.
type AcceptanceResult6 struct {

	// Indicates whether the mandate request was accepted or rejected.
	Accepted *YesNoIndicator `xml:"Accptd"`

	// Specifies the reason for the rejection of a mandate request.
	RejectReason *MandateReason1Choice `xml:"RjctRsn,omitempty"`

	// Further details on the reject reason.
	AdditionalRejectReasonInformation []*Max105Text `xml:"AddtlRjctRsnInf,omitempty"`
}

func (a *AcceptanceResult6) SetAccepted(value string) {
	a.Accepted = (*YesNoIndicator)(&value)
}

func (a *AcceptanceResult6) AddRejectReason() *MandateReason1Choice {
	a.RejectReason = new(MandateReason1Choice)
	return a.RejectReason
}

func (a *AcceptanceResult6) AddAdditionalRejectReasonInformation(value string) {
	a.AdditionalRejectReasonInformation = append(a.AdditionalRejectReasonInformation, (*Max105Text)(&value))
}
