package model

// Choice of format for the statement query unmatched reason.
type UnmatchedReason8Choice struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason6Code `xml:"Cd"`

	// Specifies the reason why the instruction has an unmatched status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (u *UnmatchedReason8Choice) SetCode(value string) {
	u.Code = (*UnmatchedReason6Code)(&value)
}

func (u *UnmatchedReason8Choice) AddProprietary() *GenericIdentification20 {
	u.Proprietary = new(GenericIdentification20)
	return u.Proprietary
}
