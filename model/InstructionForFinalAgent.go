package model

// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the final agent.
type InstructionForFinalAgent struct {

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the final agent, in coded form.
	Code []*Instruction3Code `xml:"Cd,omitempty"`

	// Instruction to the final agent that is specific to a user community and is required for use within that user community.
	//
	// Usage : The proprietary element should only be used when the coded element does not provide sufficient codes or when the selected code in the coded element needs to be supplemented by additional information such as a passport number or telephone number.
	Proprietary *Max140Text `xml:"Prtry,omitempty"`
}

func (i *InstructionForFinalAgent) AddCode(value string) {
	i.Code = append(i.Code, (*Instruction3Code)(&value))
}

func (i *InstructionForFinalAgent) SetProprietary(value string) {
	i.Proprietary = (*Max140Text)(&value)
}
