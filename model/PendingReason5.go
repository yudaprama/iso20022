package model

// The status of an instruction, advice or request.
type PendingReason5 struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason13Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason5) AddCode() *PendingReason13Choice {
	p.Code = new(PendingReason13Choice)
	return p.Code
}

func (p *PendingReason5) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
