package model

// Choice of a clearing system identifier.
type ClearingSystemIdentification2Choice struct {

	// Identification of a clearing system, in a coded form as published in an external list.
	Code *ExternalClearingSystemIdentification1Code `xml:"Cd"`

	// Identification code for a clearing system, that has not yet been identified in the list of clearing systems.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (c *ClearingSystemIdentification2Choice) SetCode(value string) {
	c.Code = (*ExternalClearingSystemIdentification1Code)(&value)
}

func (c *ClearingSystemIdentification2Choice) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
