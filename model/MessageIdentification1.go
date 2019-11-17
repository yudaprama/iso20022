package model

// Identifies a message by a unique identifier and the date and time when the message was created by the sender.
type MessageIdentification1 struct {

	// Identification of the message.
	Identification *Max35Text `xml:"Id"`

	// Date of creation of the message.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`
}

func (m *MessageIdentification1) SetIdentification(value string) {
	m.Identification = (*Max35Text)(&value)
}

func (m *MessageIdentification1) SetCreationDateTime(value string) {
	m.CreationDateTime = (*ISODateTime)(&value)
}
