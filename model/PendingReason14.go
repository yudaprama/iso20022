package model

// Specifies the reason why the instruction or request has a pending status.
type PendingReason14 struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason26Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason14) AddCode() *PendingReason26Choice {
	p.Code = new(PendingReason26Choice)
	return p.Code
}

func (p *PendingReason14) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
