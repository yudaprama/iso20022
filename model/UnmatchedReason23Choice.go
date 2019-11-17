package model

// Choice of format for the settlement transaction unmatched reason.
type UnmatchedReason23Choice struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason12Code `xml:"Cd"`

	// Specifies the reason why the instruction has an unmatched status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (u *UnmatchedReason23Choice) SetCode(value string) {
	u.Code = (*UnmatchedReason12Code)(&value)
}

func (u *UnmatchedReason23Choice) AddProprietary() *GenericIdentification30 {
	u.Proprietary = new(GenericIdentification30)
	return u.Proprietary
}
