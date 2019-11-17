package model

// Specifies the reason why the verified identification information is incorrect.
type VerificationReason1Choice struct {

	// Reason why the verified identification information is incorrect, as published in an external reason code list.
	Code *ExternalVerificationReason1Code `xml:"Cd"`

	// Reason why the verified identification information is incorrect, in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (v *VerificationReason1Choice) SetCode(value string) {
	v.Code = (*ExternalVerificationReason1Code)(&value)
}

func (v *VerificationReason1Choice) SetProprietary(value string) {
	v.Proprietary = (*Max35Text)(&value)
}
