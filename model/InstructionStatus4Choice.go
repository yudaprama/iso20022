package model

// Status applying globally to the instruction received.
type InstructionStatus4Choice struct {

	// Status advising on the processing of the instruction.
	ProcessingStatus *InstructionProcessingStatus1 `xml:"PrcgSts"`

	// Status advising on the rejection of the instruction.
	RejectionStatus *AdditionalStatus1 `xml:"RjctnSts"`
}

func (i *InstructionStatus4Choice) AddProcessingStatus() *InstructionProcessingStatus1 {
	i.ProcessingStatus = new(InstructionProcessingStatus1)
	return i.ProcessingStatus
}

func (i *InstructionStatus4Choice) AddRejectionStatus() *AdditionalStatus1 {
	i.RejectionStatus = new(AdditionalStatus1)
	return i.RejectionStatus
}
