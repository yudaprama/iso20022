package model

// Status applying to the instruction request received.
type InstructionStatus1Choice struct {

	// Status applying globally to the instruction received. The instruction is identified by the InstructionIdentification.
	GlobalInstructionStatus *InstructionStatus2Choice `xml:"GblInstrSts"`

	// Status applying to individual instructions of a MeetingInstruction.
	DetailedInstructionStatus []*DetailedInstructionStatus1 `xml:"DtldInstrSts"`
}

func (i *InstructionStatus1Choice) AddGlobalInstructionStatus() *InstructionStatus2Choice {
	i.GlobalInstructionStatus = new(InstructionStatus2Choice)
	return i.GlobalInstructionStatus
}

func (i *InstructionStatus1Choice) AddDetailedInstructionStatus() *DetailedInstructionStatus1 {
	newValue := new(DetailedInstructionStatus1)
	i.DetailedInstructionStatus = append(i.DetailedInstructionStatus, newValue)
	return newValue
}
