package model

// Choice of format for the statement query unmatched reason.
type UnmatchedReason29Choice struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason14Code `xml:"Cd"`

	// Specifies the reason why the instruction has an unmatched status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (u *UnmatchedReason29Choice) SetCode(value string) {
	u.Code = (*UnmatchedReason14Code)(&value)
}

func (u *UnmatchedReason29Choice) AddProprietary() *GenericIdentification47 {
	u.Proprietary = new(GenericIdentification47)
	return u.Proprietary
}
