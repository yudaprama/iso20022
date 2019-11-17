package model

// Proprietary identification of the reason related to a status.
type ProprietaryReason1 struct {

	// Proprietary identification of the reason related to a status.
	Reason *GenericIdentification20 `xml:"Rsn,omitempty"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *ProprietaryReason1) AddReason() *GenericIdentification20 {
	p.Reason = new(GenericIdentification20)
	return p.Reason
}

func (p *ProprietaryReason1) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
