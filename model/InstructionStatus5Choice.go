package model

// Status applying to the instruction request received.
type InstructionStatus5Choice struct {

	// Status applying globally to the instruction received. The instruction is identified by the InstructionIdentification.
	GlobalInstructionStatus *InstructionStatus6Choice `xml:"GblInstrSts"`

	// Status applying to individual instructions of a MeetingInstruction.
	DetailedInstructionStatus []*DetailedInstructionStatus11 `xml:"DtldInstrSts"`
}

func (i *InstructionStatus5Choice) AddGlobalInstructionStatus() *InstructionStatus6Choice {
	i.GlobalInstructionStatus = new(InstructionStatus6Choice)
	return i.GlobalInstructionStatus
}

func (i *InstructionStatus5Choice) AddDetailedInstructionStatus() *DetailedInstructionStatus11 {
	newValue := new(DetailedInstructionStatus11)
	i.DetailedInstructionStatus = append(i.DetailedInstructionStatus, newValue)
	return newValue
}
