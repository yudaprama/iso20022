package model

// Choice of format for the securities financing unmatched reason.
type UnmatchedReason25Choice struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason13Code `xml:"Cd"`

	// Specifies the reason why the instruction has an unmatched status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (u *UnmatchedReason25Choice) SetCode(value string) {
	u.Code = (*UnmatchedReason13Code)(&value)
}

func (u *UnmatchedReason25Choice) AddProprietary() *GenericIdentification47 {
	u.Proprietary = new(GenericIdentification47)
	return u.Proprietary
}
