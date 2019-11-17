package model

// Choice of format for the securities financing unmatched reason.
type UnmatchedReason24Choice struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason13Code `xml:"Cd"`

	// Specifies the reason why the instruction has an unmatched status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (u *UnmatchedReason24Choice) SetCode(value string) {
	u.Code = (*UnmatchedReason13Code)(&value)
}

func (u *UnmatchedReason24Choice) AddProprietary() *GenericIdentification30 {
	u.Proprietary = new(GenericIdentification30)
	return u.Proprietary
}
