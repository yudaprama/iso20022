package model

// Choice of formats for a confirmation type.
type ConfirmationType1Choice struct {

	// Confirmation type expressed as a code.
	Code *AccountManagementType2Code `xml:"Cd"`

	// Confirmation type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *ConfirmationType1Choice) SetCode(value string) {
	c.Code = (*AccountManagementType2Code)(&value)
}

func (c *ConfirmationType1Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
