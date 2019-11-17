package model

// Further information related to the processing of the payment instruction that may need to be acted upon by an other agent. The instruction may relate to a level of service, or may be an instruction that has to be executed by the creditor's agent, or may be information required by the other agent.
type InstructionForNextAgent1 struct {

	// Coded information related to the processing of the payment instruction, provided by the initiating party, and intended for the next agent in the payment chain.
	Code *Instruction4Code `xml:"Cd,omitempty"`

	// Further information complementing the coded instruction or instruction to the next agent that is bilaterally agreed or specific to a user community.
	InstructionInformation *Max140Text `xml:"InstrInf,omitempty"`
}

func (i *InstructionForNextAgent1) SetCode(value string) {
	i.Code = (*Instruction4Code)(&value)
}

func (i *InstructionForNextAgent1) SetInstructionInformation(value string) {
	i.InstructionInformation = (*Max140Text)(&value)
}
