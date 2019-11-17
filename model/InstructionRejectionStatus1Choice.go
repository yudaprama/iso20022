package model

// Choice of instruction rejection status.
type InstructionRejectionStatus1Choice struct {

	// Reason advising the rejection of the instruction in the form of a code.
	Code *RejectionReason1Code `xml:"Cd"`

	// This code can be used in case another reason is required.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (i *InstructionRejectionStatus1Choice) SetCode(value string) {
	i.Code = (*RejectionReason1Code)(&value)
}

func (i *InstructionRejectionStatus1Choice) AddProprietary() *GenericIdentification13 {
	i.Proprietary = new(GenericIdentification13)
	return i.Proprietary
}
