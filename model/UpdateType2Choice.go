package model

// Choice of format for the update information.
type UpdateType2Choice struct {

	// Indicates whether the report is complete or contains changes only.
	Code *StatementUpdateType1Code `xml:"Cd"`

	// Indicates whether the report is complete or contains changes only.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (u *UpdateType2Choice) SetCode(value string) {
	u.Code = (*StatementUpdateType1Code)(&value)
}

func (u *UpdateType2Choice) AddProprietary() *GenericIdentification20 {
	u.Proprietary = new(GenericIdentification20)
	return u.Proprietary
}
