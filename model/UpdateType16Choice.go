package model

// Choice of format for the update information.
type UpdateType16Choice struct {

	// Indicates whether the report is complete or contains changes only.
	Code *StatementUpdateType1Code `xml:"Cd"`

	// Indicates whether the report is complete or contains changes only.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (u *UpdateType16Choice) SetCode(value string) {
	u.Code = (*StatementUpdateType1Code)(&value)
}

func (u *UpdateType16Choice) AddProprietary() *GenericIdentification47 {
	u.Proprietary = new(GenericIdentification47)
	return u.Proprietary
}
