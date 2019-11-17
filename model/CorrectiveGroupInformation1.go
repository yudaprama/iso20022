package model

// Set of elements used to provide information on the group of the corrective transaction, to which the resolution message refers.
type CorrectiveGroupInformation1 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Specifies the message name identifier to which the message refers.
	MessageNameIdentification *Max35Text `xml:"MsgNmId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm,omitempty"`
}

func (c *CorrectiveGroupInformation1) SetMessageIdentification(value string) {
	c.MessageIdentification = (*Max35Text)(&value)
}

func (c *CorrectiveGroupInformation1) SetMessageNameIdentification(value string) {
	c.MessageNameIdentification = (*Max35Text)(&value)
}

func (c *CorrectiveGroupInformation1) SetCreationDateTime(value string) {
	c.CreationDateTime = (*ISODateTime)(&value)
}
