package model

// Specifies the reason why the instruction or request has a pending status.
type PendingReason21 struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason38Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason21) AddCode() *PendingReason38Choice {
	p.Code = new(PendingReason38Choice)
	return p.Code
}

func (p *PendingReason21) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
