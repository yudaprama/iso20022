package model

// Choice of format for the unaffirmed reason.
type UnaffirmedReason2Choice struct {

	// Specifies the reason why the instruction/request has an unaffirmed status.
	Code *UnaffirmedReason1Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a unaffirmed status.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (u *UnaffirmedReason2Choice) SetCode(value string) {
	u.Code = (*UnaffirmedReason1Code)(&value)
}

func (u *UnaffirmedReason2Choice) AddProprietary() *GenericIdentification38 {
	u.Proprietary = new(GenericIdentification38)
	return u.Proprietary
}
