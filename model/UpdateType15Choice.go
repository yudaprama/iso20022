package model

// Choice of format for the update information.
type UpdateType15Choice struct {

	// Indicates whether the report is complete or contains changes only.
	Code *StatementUpdateType1Code `xml:"Cd"`

	// Indicates whether the report is complete or contains changes only.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (u *UpdateType15Choice) SetCode(value string) {
	u.Code = (*StatementUpdateType1Code)(&value)
}

func (u *UpdateType15Choice) AddProprietary() *GenericIdentification30 {
	u.Proprietary = new(GenericIdentification30)
	return u.Proprietary
}
