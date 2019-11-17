package model

// Choice of format for the statement query unmatched reason.
type UnmatchedReason16Choice struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason9Code `xml:"Cd"`

	// Specifies the reason why the instruction has an unmatched status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (u *UnmatchedReason16Choice) SetCode(value string) {
	u.Code = (*UnmatchedReason9Code)(&value)
}

func (u *UnmatchedReason16Choice) AddProprietary() *GenericIdentification20 {
	u.Proprietary = new(GenericIdentification20)
	return u.Proprietary
}
