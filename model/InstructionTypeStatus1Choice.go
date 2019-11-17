package model

// Choice of instruction type status.
type InstructionTypeStatus1Choice struct {

	// Status applying to the instruction request received. The instruction is identified by the InstructionIdentification.
	InstructionStatus *InstructionStatus3Choice `xml:"InstrSts"`

	// Status applying to the instruction cancellation request received. The instruction cancellation is identified by the InstructionCancellationIdentification.
	CancellationStatus *CancellationStatus2Choice `xml:"CxlSts"`
}

func (i *InstructionTypeStatus1Choice) AddInstructionStatus() *InstructionStatus3Choice {
	i.InstructionStatus = new(InstructionStatus3Choice)
	return i.InstructionStatus
}

func (i *InstructionTypeStatus1Choice) AddCancellationStatus() *CancellationStatus2Choice {
	i.CancellationStatus = new(CancellationStatus2Choice)
	return i.CancellationStatus
}
