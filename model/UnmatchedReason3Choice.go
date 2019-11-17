package model

// Choice of format for the securities financing unmatched reason.
type UnmatchedReason3Choice struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason1Code `xml:"Cd"`

	// Specifies the reason why the instruction has an unmatched status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (u *UnmatchedReason3Choice) SetCode(value string) {
	u.Code = (*UnmatchedReason1Code)(&value)
}

func (u *UnmatchedReason3Choice) AddProprietary() *GenericIdentification20 {
	u.Proprietary = new(GenericIdentification20)
	return u.Proprietary
}
