package model

// Specifies the reason why the instruction or request has a pending status.
type PendingReason19 struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason36Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason19) AddCode() *PendingReason36Choice {
	p.Code = new(PendingReason36Choice)
	return p.Code
}

func (p *PendingReason19) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
