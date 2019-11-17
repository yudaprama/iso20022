package model

// Specifies the details of missing information or the complete set of available information.
type UnableToApplyJustification2Choice struct {

	// Indicates whether or not all available information on the underlying payment instruction is requested.
	AnyInformation *YesNoIndicator `xml:"AnyInf"`

	// Set of elements used to indicate which information is missing or incorrect.
	MissingOrIncorrectInformation *MissingOrIncorrectInformation2 `xml:"MssngOrIncrrctInf"`

	// Indicates whether or not the referred item is a possible duplicate of a previous instruction or entry.
	PossibleDuplicateInstruction *TrueFalseIndicator `xml:"PssblDplctInstr"`
}

func (u *UnableToApplyJustification2Choice) SetAnyInformation(value string) {
	u.AnyInformation = (*YesNoIndicator)(&value)
}

func (u *UnableToApplyJustification2Choice) AddMissingOrIncorrectInformation() *MissingOrIncorrectInformation2 {
	u.MissingOrIncorrectInformation = new(MissingOrIncorrectInformation2)
	return u.MissingOrIncorrectInformation
}

func (u *UnableToApplyJustification2Choice) SetPossibleDuplicateInstruction(value string) {
	u.PossibleDuplicateInstruction = (*TrueFalseIndicator)(&value)
}
