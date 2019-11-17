package model

// Choice of format for the update information.
type UpdateType4Choice struct {

	// Update type expressed in coded form.
	Code *StatementUpdateType1Code `xml:"Cd"`

	// Update type expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (u *UpdateType4Choice) SetCode(value string) {
	u.Code = (*StatementUpdateType1Code)(&value)
}

func (u *UpdateType4Choice) AddProprietary() *GenericIdentification30 {
	u.Proprietary = new(GenericIdentification30)
	return u.Proprietary
}
