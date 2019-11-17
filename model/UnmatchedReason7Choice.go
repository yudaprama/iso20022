package model

// Choice of format for the statement query unmatched reason.
type UnmatchedReason7Choice struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason4Code `xml:"Cd"`

	// Specifies the reason why the instruction has an unmatched status.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (u *UnmatchedReason7Choice) SetCode(value string) {
	u.Code = (*UnmatchedReason4Code)(&value)
}

func (u *UnmatchedReason7Choice) AddProprietary() *GenericIdentification38 {
	u.Proprietary = new(GenericIdentification38)
	return u.Proprietary
}
