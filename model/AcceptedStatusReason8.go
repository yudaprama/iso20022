package model

// Specifies reasons for the accepted status.
type AcceptedStatusReason8 struct {

	// Specifies the reason why the instruction has been accepted.
	ReasonCode *AcceptedReason9Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcceptedStatusReason8) AddReasonCode() *AcceptedReason9Choice {
	a.ReasonCode = new(AcceptedReason9Choice)
	return a.ReasonCode
}

func (a *AcceptedStatusReason8) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}
