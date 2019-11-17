package model

// Choice of format for the communication channel method.
type CommunicationMethod1Choice struct {

	// Communication channel method.
	//
	Code *ExternalChannel1Code `xml:"Cd"`

	// Communication channel method expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (c *CommunicationMethod1Choice) SetCode(value string) {
	c.Code = (*ExternalChannel1Code)(&value)
}

func (c *CommunicationMethod1Choice) AddProprietary() *GenericIdentification1 {
	c.Proprietary = new(GenericIdentification1)
	return c.Proprietary
}
