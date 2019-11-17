package model

// Set of elements used to provide further details on the message.
type GroupHeader43 struct {

	// Point to point reference, as assigned by the account owner or the party authorised to send the message, and sent to the account servicing institution, to unambiguously identify the message.
	// Usage: The sender has to make sure that 'MessageIdentification' is unique per account servicing institution for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Identification of the party that is sending the message, when different from the account owner.
	MessageSender *Party7Choice `xml:"MsgSndr,omitempty"`
}

func (g *GroupHeader43) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader43) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader43) AddMessageSender() *Party7Choice {
	g.MessageSender = new(Party7Choice)
	return g.MessageSender
}
