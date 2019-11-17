package model

// Specifies the reason why the instruction or request has a pending status.
type PendingReason24 struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason41Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason24) AddCode() *PendingReason41Choice {
	p.Code = new(PendingReason41Choice)
	return p.Code
}

func (p *PendingReason24) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
