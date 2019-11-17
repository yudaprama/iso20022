package model

// Set of elements used to provide further details on the message.
type GroupHeader44 struct {

	// Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message.
	//
	// Usage: The account servicing institution has to make sure that 'MessageIdentification' is unique per account owner for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Identification of the party that is receiving the message, when different from the account owner.
	MessageRecipient *Party7Choice `xml:"MsgRcpt,omitempty"`
}

func (g *GroupHeader44) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader44) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader44) AddMessageRecipient() *Party7Choice {
	g.MessageRecipient = new(Party7Choice)
	return g.MessageRecipient
}
