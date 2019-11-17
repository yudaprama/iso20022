package model

// Provides further details on the message.
type GroupHeader59 struct {

	// Point to point reference, as assigned by the account owner or the party authorised to send the message, and sent to the account servicing institution, to unambiguously identify the message.
	// Usage: The sender has to make sure that 'MessageIdentification' is unique per account servicing institution for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Identification of the party that is sending the message, when different from the account owner.
	MessageSender *Party12Choice `xml:"MsgSndr,omitempty"`
}

func (g *GroupHeader59) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader59) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader59) AddMessageSender() *Party12Choice {
	g.MessageSender = new(Party12Choice)
	return g.MessageSender
}
