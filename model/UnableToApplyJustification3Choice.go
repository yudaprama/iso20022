package model

// Specifies the details of missing or incorrect information or the complete set of available information.
type UnableToApplyJustification3Choice struct {

	// Indicates whether or not all available information on the underlying payment instruction is requested.
	AnyInformation *YesNoIndicator `xml:"AnyInf"`

	// Set of elements used to indicate which information is missing or incorrect.
	MissingOrIncorrectInformation *MissingOrIncorrectInformation3 `xml:"MssngOrIncrrctInf"`

	// Indicates whether or not the referred item is a possible duplicate of a previous instruction or entry.
	PossibleDuplicateInstruction *TrueFalseIndicator `xml:"PssblDplctInstr"`
}

func (u *UnableToApplyJustification3Choice) SetAnyInformation(value string) {
	u.AnyInformation = (*YesNoIndicator)(&value)
}

func (u *UnableToApplyJustification3Choice) AddMissingOrIncorrectInformation() *MissingOrIncorrectInformation3 {
	u.MissingOrIncorrectInformation = new(MissingOrIncorrectInformation3)
	return u.MissingOrIncorrectInformation
}

func (u *UnableToApplyJustification3Choice) SetPossibleDuplicateInstruction(value string) {
	u.PossibleDuplicateInstruction = (*TrueFalseIndicator)(&value)
}
