package model

// Choice of format for the settlement transaction unmatched reason.
type UnmatchedReason21Choice struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason11Code `xml:"Cd"`

	// Specifies the reason why the instruction has an unmatched status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (u *UnmatchedReason21Choice) SetCode(value string) {
	u.Code = (*UnmatchedReason11Code)(&value)
}

func (u *UnmatchedReason21Choice) AddProprietary() *GenericIdentification30 {
	u.Proprietary = new(GenericIdentification30)
	return u.Proprietary
}
