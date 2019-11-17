package model

// Specifies reasons for the pending status.
type PendingStatusReason15 struct {

	// Specifies the reason why the instruction's processing is pending.
	ReasonCode *PendingReason49Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingStatusReason15) AddReasonCode() *PendingReason49Choice {
	p.ReasonCode = new(PendingReason49Choice)
	return p.ReasonCode
}

func (p *PendingStatusReason15) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
