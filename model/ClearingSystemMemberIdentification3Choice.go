package model

// Choice of identifiers for a clearing system member, as assigned by the clearing system. In some clearing systems, the accounts of the clearing system members are also assigned an identifier.
type ClearingSystemMemberIdentification3Choice struct {

	// Identification for a clearing system member, identified in the list of clearing system member identifications published externally.
	Identification *ExternalClearingSystemMemberCode `xml:"Id"`

	// Identification Code for a clearing system, that has not yet been identified in the list of clearing systems.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (c *ClearingSystemMemberIdentification3Choice) SetIdentification(value string) {
	c.Identification = (*ExternalClearingSystemMemberCode)(&value)
}

func (c *ClearingSystemMemberIdentification3Choice) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
