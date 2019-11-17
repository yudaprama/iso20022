package model

// Status applying to the instruction request received.
type InstructionStatus3Choice struct {

	// Status applying globally to the instruction received. The instruction is identified by the InstructionIdentification.
	GlobalInstructionStatus *InstructionStatus4Choice `xml:"GblInstrSts"`

	// Status applying to individual instructions of a MeetingInstruction.
	DetailedInstructionStatus []*DetailedInstructionStatus8 `xml:"DtldInstrSts"`
}

func (i *InstructionStatus3Choice) AddGlobalInstructionStatus() *InstructionStatus4Choice {
	i.GlobalInstructionStatus = new(InstructionStatus4Choice)
	return i.GlobalInstructionStatus
}

func (i *InstructionStatus3Choice) AddDetailedInstructionStatus() *DetailedInstructionStatus8 {
	newValue := new(DetailedInstructionStatus8)
	i.DetailedInstructionStatus = append(i.DetailedInstructionStatus, newValue)
	return newValue
}
