package model

// Choice of format for the settlement transaction unmatched reason.
type UnmatchedReason9Choice struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason5Code `xml:"Cd"`

	// Specifies the reason why the instruction has an unmatched status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (u *UnmatchedReason9Choice) SetCode(value string) {
	u.Code = (*UnmatchedReason5Code)(&value)
}

func (u *UnmatchedReason9Choice) AddProprietary() *GenericIdentification20 {
	u.Proprietary = new(GenericIdentification20)
	return u.Proprietary
}
