package model

// Choice between details of missing information or the complete set of available information.
type UnableToApplyJustificationChoice struct {

	// When set to yes, indicates that all available information about the underlying payment instruction shall be sent.
	AnyInformation *YesNoIndicator `xml:"AnyInf"`

	// Missing or incorrect information.
	MissingOrIncorrectInformation *MissingOrIncorrectInformation `xml:"MssngOrIncrrctInf"`
}

func (u *UnableToApplyJustificationChoice) SetAnyInformation(value string) {
	u.AnyInformation = (*YesNoIndicator)(&value)
}

func (u *UnableToApplyJustificationChoice) AddMissingOrIncorrectInformation() *MissingOrIncorrectInformation {
	u.MissingOrIncorrectInformation = new(MissingOrIncorrectInformation)
	return u.MissingOrIncorrectInformation
}
