package model

// The status of an instruction, advice or request.
type PendingProcessingReason3 struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason3Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingProcessingReason3) AddCode() *PendingProcessingReason3Choice {
	p.Code = new(PendingProcessingReason3Choice)
	return p.Code
}

func (p *PendingProcessingReason3) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
