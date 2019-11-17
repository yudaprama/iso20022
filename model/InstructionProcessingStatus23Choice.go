package model

// Choice of format for the instruction processing status.
type InstructionProcessingStatus23Choice struct {

	// Provides the status of an instruction.
	Code *InstructionProcessingStatus1Code `xml:"Cd"`

	// Provides the status of an instruction.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (i *InstructionProcessingStatus23Choice) SetCode(value string) {
	i.Code = (*InstructionProcessingStatus1Code)(&value)
}

func (i *InstructionProcessingStatus23Choice) AddProprietary() *GenericIdentification30 {
	i.Proprietary = new(GenericIdentification30)
	return i.Proprietary
}
