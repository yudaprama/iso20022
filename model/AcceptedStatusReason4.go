package model

// Specifies reasons for the accepted status.
type AcceptedStatusReason4 struct {

	// Specifies the reason why the instruction has been accepted.
	ReasonCode *AcceptedReason4Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcceptedStatusReason4) AddReasonCode() *AcceptedReason4Choice {
	a.ReasonCode = new(AcceptedReason4Choice)
	return a.ReasonCode
}

func (a *AcceptedStatusReason4) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}
