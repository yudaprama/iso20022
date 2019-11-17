package model

// Specifies the type of instruction requested by the submitter by means of a code.
type InstructionType3 struct {

	// Specifies whether the data set has to be matched or pre-matched.
	Type *InstructionType3Code `xml:"Tp"`
}

func (i *InstructionType3) SetType(value string) {
	i.Type = (*InstructionType3Code)(&value)
}
