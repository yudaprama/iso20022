package model

// Further information related to the processing of the payment instruction that may need to be acted upon by the creditor's agent. The instruction may relate to a level of service, or may be an instruction that has to be executed by the creditor's agent, or may be information required by the creditor's agent.
type InstructionForCreditorAgent2 struct {

	// Coded information related to the processing of the payment instruction, provided by the initiating party, and intended for the creditor's agent.
	Code *Instruction5Code `xml:"Cd,omitempty"`

	// Further information complementing the coded instruction or instruction to the creditor's agent that is bilaterally agreed or specific to a user community.
	InstructionInformation *Max140Text `xml:"InstrInf,omitempty"`
}

func (i *InstructionForCreditorAgent2) SetCode(value string) {
	i.Code = (*Instruction5Code)(&value)
}

func (i *InstructionForCreditorAgent2) SetInstructionInformation(value string) {
	i.InstructionInformation = (*Max140Text)(&value)
}
