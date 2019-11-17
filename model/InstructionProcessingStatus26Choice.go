package model

// Choice of format for the instruction processing status.
type InstructionProcessingStatus26Choice struct {

	// Provides the status of an instruction.
	Code *InstructionProcessingStatus1Code `xml:"Cd"`

	// Provides the status of an instruction.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (i *InstructionProcessingStatus26Choice) SetCode(value string) {
	i.Code = (*InstructionProcessingStatus1Code)(&value)
}

func (i *InstructionProcessingStatus26Choice) AddProprietary() *GenericIdentification47 {
	i.Proprietary = new(GenericIdentification47)
	return i.Proprietary
}
