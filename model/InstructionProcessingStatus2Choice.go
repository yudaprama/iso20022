package model

// Choice of format for the instruction processing status.
type InstructionProcessingStatus2Choice struct {

	// Provides the status of an instruction.
	Code *InstructionProcessingStatus1Code `xml:"Cd"`

	// Provides the status of an instruction.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (i *InstructionProcessingStatus2Choice) SetCode(value string) {
	i.Code = (*InstructionProcessingStatus1Code)(&value)
}

func (i *InstructionProcessingStatus2Choice) AddProprietary() *GenericIdentification20 {
	i.Proprietary = new(GenericIdentification20)
	return i.Proprietary
}
