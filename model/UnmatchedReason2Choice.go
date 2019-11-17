package model

// Choice of format for the statement query unmatched reason.
type UnmatchedReason2Choice struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason3Code `xml:"Cd"`

	// Specifies the reason why the instruction has an unmatched status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (u *UnmatchedReason2Choice) SetCode(value string) {
	u.Code = (*UnmatchedReason3Code)(&value)
}

func (u *UnmatchedReason2Choice) AddProprietary() *GenericIdentification20 {
	u.Proprietary = new(GenericIdentification20)
	return u.Proprietary
}
