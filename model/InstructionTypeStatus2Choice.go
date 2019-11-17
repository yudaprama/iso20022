package model

// Choice of instruction type status.
type InstructionTypeStatus2Choice struct {

	// Status applying to the instruction request received. The instruction is identified by the InstructionIdentification.
	InstructionStatus *InstructionStatus5Choice `xml:"InstrSts"`

	// Status applying to the instruction cancellation request received. The instruction cancellation is identified by the InstructionCancellationIdentification.
	CancellationStatus *CancellationStatus2Choice `xml:"CxlSts"`
}

func (i *InstructionTypeStatus2Choice) AddInstructionStatus() *InstructionStatus5Choice {
	i.InstructionStatus = new(InstructionStatus5Choice)
	return i.InstructionStatus
}

func (i *InstructionTypeStatus2Choice) AddCancellationStatus() *CancellationStatus2Choice {
	i.CancellationStatus = new(CancellationStatus2Choice)
	return i.CancellationStatus
}
