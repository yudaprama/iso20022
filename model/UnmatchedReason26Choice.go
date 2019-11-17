package model

// Choice of format for the settlement transaction unmatched reason.
type UnmatchedReason26Choice struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason12Code `xml:"Cd"`

	// Specifies the reason why the instruction has an unmatched status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (u *UnmatchedReason26Choice) SetCode(value string) {
	u.Code = (*UnmatchedReason12Code)(&value)
}

func (u *UnmatchedReason26Choice) AddProprietary() *GenericIdentification47 {
	u.Proprietary = new(GenericIdentification47)
	return u.Proprietary
}
