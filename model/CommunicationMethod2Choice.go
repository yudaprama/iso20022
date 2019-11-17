package model

// Specifies a choice for the method of communication for documents.
type CommunicationMethod2Choice struct {

	// Unique and unambiguous identification of communication method using a code list.
	Code *CommunicationMethod2Code `xml:"Cd"`

	// Unique and unambiguous identification of communication method using a bilaterally or multilaterally agreed description.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (c *CommunicationMethod2Choice) SetCode(value string) {
	c.Code = (*CommunicationMethod2Code)(&value)
}

func (c *CommunicationMethod2Choice) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
