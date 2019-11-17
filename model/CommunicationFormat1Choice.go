package model

// Specifies a choice of format for statements.
type CommunicationFormat1Choice struct {

	// Unique and unambiguous identification of communication format using a code list.
	Code *ExternalCommunicationFormat1Code `xml:"Cd"`

	// Unique and unambiguous identification of communication format using a proprietary identification scheme.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (c *CommunicationFormat1Choice) SetCode(value string) {
	c.Code = (*ExternalCommunicationFormat1Code)(&value)
}

func (c *CommunicationFormat1Choice) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
