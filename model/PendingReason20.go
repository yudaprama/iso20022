package model

// Specifies the reason why the instruction or request has a pending status.
type PendingReason20 struct {

	// Specifies the reason why a cancellation request sent for the related instruction is pending.
	Code *PendingReason37Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason20) AddCode() *PendingReason37Choice {
	p.Code = new(PendingReason37Choice)
	return p.Code
}

func (p *PendingReason20) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
