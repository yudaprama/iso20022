package model

// Choice of format for the statement query unmatched reason.
type UnmatchedReason22Choice struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason14Code `xml:"Cd"`

	// Specifies the reason why the instruction has an unmatched status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (u *UnmatchedReason22Choice) SetCode(value string) {
	u.Code = (*UnmatchedReason14Code)(&value)
}

func (u *UnmatchedReason22Choice) AddProprietary() *GenericIdentification30 {
	u.Proprietary = new(GenericIdentification30)
	return u.Proprietary
}
