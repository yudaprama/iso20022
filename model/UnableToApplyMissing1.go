package model

// Specifies the details of missing information.
type UnableToApplyMissing1 struct {

	// Specifies the missing information in a coded form.
	Code *UnableToApplyMissingInformation3Code `xml:"Cd"`

	// Further details about the missing information.
	AdditionalMissingInformation *Max140Text `xml:"AddtlMssngInf,omitempty"`
}

func (u *UnableToApplyMissing1) SetCode(value string) {
	u.Code = (*UnableToApplyMissingInformation3Code)(&value)
}

func (u *UnableToApplyMissing1) SetAdditionalMissingInformation(value string) {
	u.AdditionalMissingInformation = (*Max140Text)(&value)
}
