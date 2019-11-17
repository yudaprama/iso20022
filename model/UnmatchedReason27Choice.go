package model

// Choice of format for the settlement transaction unmatched reason.
type UnmatchedReason27Choice struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason11Code `xml:"Cd"`

	// Specifies the reason why the instruction has an unmatched status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (u *UnmatchedReason27Choice) SetCode(value string) {
	u.Code = (*UnmatchedReason11Code)(&value)
}

func (u *UnmatchedReason27Choice) AddProprietary() *GenericIdentification47 {
	u.Proprietary = new(GenericIdentification47)
	return u.Proprietary
}
