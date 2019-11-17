package model

// Specifies reasons for the pending status.
type PendingStatusReason11 struct {

	// Specifies the reason why the instruction's processing is pending.
	ReasonCode *PendingReason34Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingStatusReason11) AddReasonCode() *PendingReason34Choice {
	p.ReasonCode = new(PendingReason34Choice)
	return p.ReasonCode
}

func (p *PendingStatusReason11) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
