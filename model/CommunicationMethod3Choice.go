package model

// Choice of formats for the method of transmission.
type CommunicationMethod3Choice struct {

	// Method of transmission expressed as a code.
	Code *CommunicationMethod1Code `xml:"Cd"`

	// Method of transmission expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CommunicationMethod3Choice) SetCode(value string) {
	c.Code = (*CommunicationMethod1Code)(&value)
}

func (c *CommunicationMethod3Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
