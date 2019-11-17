package model

// Specifies the details of incorrect information.
type UnableToApplyIncorrect1 struct {

	// Specifies the missing information in a coded form.
	Code *UnableToApplyIncorrectInformation4Code `xml:"Cd"`

	// Further details about the incorrect information.
	AdditionalIncorrectInformation *Max140Text `xml:"AddtlIncrrctInf,omitempty"`
}

func (u *UnableToApplyIncorrect1) SetCode(value string) {
	u.Code = (*UnableToApplyIncorrectInformation4Code)(&value)
}

func (u *UnableToApplyIncorrect1) SetAdditionalIncorrectInformation(value string) {
	u.AdditionalIncorrectInformation = (*Max140Text)(&value)
}
