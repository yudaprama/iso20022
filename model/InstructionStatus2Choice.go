package model

// Status applying globally to the instruction received.
type InstructionStatus2Choice struct {

	// Status advising on the processing of the instruction.
	ProcessingStatus *InstructionProcessingStatus1 `xml:"PrcgSts"`

	// Status advising on the rejection of the instruction.
	RejectionStatus *InstructionRejectionStatus1 `xml:"RjctnSts"`
}

func (i *InstructionStatus2Choice) AddProcessingStatus() *InstructionProcessingStatus1 {
	i.ProcessingStatus = new(InstructionProcessingStatus1)
	return i.ProcessingStatus
}

func (i *InstructionStatus2Choice) AddRejectionStatus() *InstructionRejectionStatus1 {
	i.RejectionStatus = new(InstructionRejectionStatus1)
	return i.RejectionStatus
}
