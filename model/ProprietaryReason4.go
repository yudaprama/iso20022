package model

// Proprietary identification of the reason related to a status.
type ProprietaryReason4 struct {

	// Proprietary identification of the reason related to a status.
	Reason *GenericIdentification30 `xml:"Rsn,omitempty"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *ProprietaryReason4) AddReason() *GenericIdentification30 {
	p.Reason = new(GenericIdentification30)
	return p.Reason
}

func (p *ProprietaryReason4) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
