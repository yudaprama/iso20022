package model

// Choice of format for the securities financing unmatched reason.
type UnmatchedReason15Choice struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason10Code `xml:"Cd"`

	// Specifies the reason why the instruction has an unmatched status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (u *UnmatchedReason15Choice) SetCode(value string) {
	u.Code = (*UnmatchedReason10Code)(&value)
}

func (u *UnmatchedReason15Choice) AddProprietary() *GenericIdentification20 {
	u.Proprietary = new(GenericIdentification20)
	return u.Proprietary
}
