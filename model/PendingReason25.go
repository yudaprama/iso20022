package model

// Specifies the reason why the instruction or request has a pending status.
type PendingReason25 struct {

	// Specifies the reason why a cancellation request sent for the related instruction is pending.
	Code *PendingReason42Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason25) AddCode() *PendingReason42Choice {
	p.Code = new(PendingReason42Choice)
	return p.Code
}

func (p *PendingReason25) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
