package model

// Proprietary identification of the reason related to a status.
type ProprietaryReason5 struct {

	// Proprietary identification of the reason related to a status.
	Reason *GenericIdentification47 `xml:"Rsn,omitempty"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *ProprietaryReason5) AddReason() *GenericIdentification47 {
	p.Reason = new(GenericIdentification47)
	return p.Reason
}

func (p *ProprietaryReason5) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
