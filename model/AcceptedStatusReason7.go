package model

// Reason for an accepted status.
type AcceptedStatusReason7 struct {

	// Reason for the accepted status.
	Reason *AcceptedReason8Choice `xml:"Rsn"`

	// Additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcceptedStatusReason7) AddReason() *AcceptedReason8Choice {
	a.Reason = new(AcceptedReason8Choice)
	return a.Reason
}

func (a *AcceptedStatusReason7) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}
