package model

// The status of an instruction, advice or request.
type PendingProcessingReason5 struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason5Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingProcessingReason5) AddCode() *PendingProcessingReason5Choice {
	p.Code = new(PendingProcessingReason5Choice)
	return p.Code
}

func (p *PendingProcessingReason5) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
