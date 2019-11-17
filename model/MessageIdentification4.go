package model

// Set of element to identify a message by a unique identifier and the date and time when the message was created by the sender.
type MessageIdentification4 struct {

	// Point to point reference, as assigned by the assigner, and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The assigner has to make sure that MessageIdentification is unique per assignee for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm,omitempty"`
}

func (m *MessageIdentification4) SetMessageIdentification(value string) {
	m.MessageIdentification = (*Max35Text)(&value)
}

func (m *MessageIdentification4) SetCreationDateTime(value string) {
	m.CreationDateTime = (*ISODateTime)(&value)
}
