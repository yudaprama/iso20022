package model

// The status of an instruction, advice or request.
type PendingReason8 struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason16Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason8) AddCode() *PendingReason16Choice {
	p.Code = new(PendingReason16Choice)
	return p.Code
}

func (p *PendingReason8) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
