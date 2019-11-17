package model

// Choice of format for the settlement transaction unmatched reason.
type UnmatchedReason14Choice struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason8Code `xml:"Cd"`

	// Specifies the reason why the instruction has an unmatched status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (u *UnmatchedReason14Choice) SetCode(value string) {
	u.Code = (*UnmatchedReason8Code)(&value)
}

func (u *UnmatchedReason14Choice) AddProprietary() *GenericIdentification20 {
	u.Proprietary = new(GenericIdentification20)
	return u.Proprietary
}
