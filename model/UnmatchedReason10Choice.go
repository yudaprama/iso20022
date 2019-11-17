package model

// Choice of format for the securities financing unmatched reason.
type UnmatchedReason10Choice struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason7Code `xml:"Cd"`

	// Specifies the reason why the instruction has an unmatched status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (u *UnmatchedReason10Choice) SetCode(value string) {
	u.Code = (*UnmatchedReason7Code)(&value)
}

func (u *UnmatchedReason10Choice) AddProprietary() *GenericIdentification20 {
	u.Proprietary = new(GenericIdentification20)
	return u.Proprietary
}
