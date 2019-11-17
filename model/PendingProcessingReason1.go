package model

// The status of an instruction, advice or request.
type PendingProcessingReason1 struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason1Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingProcessingReason1) AddCode() *PendingProcessingReason1Choice {
	p.Code = new(PendingProcessingReason1Choice)
	return p.Code
}

func (p *PendingProcessingReason1) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
