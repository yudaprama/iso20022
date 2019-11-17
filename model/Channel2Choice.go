package model

// Specifies a channel type.
type Channel2Choice struct {

	// Specifies a channel by means of a code.
	Code *CommunicationMethod3Code `xml:"Cd"`

	// Specifies a channel by means of a text.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (c *Channel2Choice) SetCode(value string) {
	c.Code = (*CommunicationMethod3Code)(&value)
}

func (c *Channel2Choice) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
