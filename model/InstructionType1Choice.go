package model

// Type of instruction.
type InstructionType1Choice struct {

	// Identifies the meeting instruction message for which the status is provided.
	InstructionIdentification *MessageIdentification `xml:"InstrId"`

	// Identifies the meeting instruction cancellation request message for which the status is provided.
	InstructionCancellationIdentification *MessageIdentification `xml:"InstrCxlId"`
}

func (i *InstructionType1Choice) AddInstructionIdentification() *MessageIdentification {
	i.InstructionIdentification = new(MessageIdentification)
	return i.InstructionIdentification
}

func (i *InstructionType1Choice) AddInstructionCancellationIdentification() *MessageIdentification {
	i.InstructionCancellationIdentification = new(MessageIdentification)
	return i.InstructionCancellationIdentification
}
