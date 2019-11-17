package model

// Status applying globally to the instruction received.
type InstructionStatus6Choice struct {

	// Status advising on the processing of the instruction.
	ProcessingStatus *InstructionProcessingStatus3 `xml:"PrcgSts"`

	// Status advising on the rejection of the instruction.
	RejectionStatus *AdditionalStatus1 `xml:"RjctnSts"`
}

func (i *InstructionStatus6Choice) AddProcessingStatus() *InstructionProcessingStatus3 {
	i.ProcessingStatus = new(InstructionProcessingStatus3)
	return i.ProcessingStatus
}

func (i *InstructionStatus6Choice) AddRejectionStatus() *AdditionalStatus1 {
	i.RejectionStatus = new(AdditionalStatus1)
	return i.RejectionStatus
}
