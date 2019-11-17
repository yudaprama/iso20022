package model

// Proprietary status and reason of an instruction or an instruction cancellation.
type ProprietaryStatusAndReason5 struct {

	// Proprietary identification of the status of the instruction.
	Status *GenericIdentification36 `xml:"Sts"`

	// Proprietary identification of the reason for the status.
	Reason *ProprietaryReason1Choice `xml:"Rsn"`

	// Additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *ProprietaryStatusAndReason5) AddStatus() *GenericIdentification36 {
	p.Status = new(GenericIdentification36)
	return p.Status
}

func (p *ProprietaryStatusAndReason5) AddReason() *ProprietaryReason1Choice {
	p.Reason = new(ProprietaryReason1Choice)
	return p.Reason
}

func (p *ProprietaryStatusAndReason5) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
