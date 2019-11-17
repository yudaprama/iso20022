package model

// Specifies reasons for the accepted status.
type AcceptedStatusReason9 struct {

	// Specifies the reason why the instruction or instruction cancellation has been accepted.
	ReasonCode *AcceptedReason10Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcceptedStatusReason9) AddReasonCode() *AcceptedReason10Choice {
	a.ReasonCode = new(AcceptedReason10Choice)
	return a.ReasonCode
}

func (a *AcceptedStatusReason9) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}
