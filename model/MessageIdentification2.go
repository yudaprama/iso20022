package model

// Set of elements providing the identification of a message.
type MessageIdentification2 struct {

	// Specifies the message name identifier of the message that will be used to provide additional details.
	MessageNameIdentification *Max35Text `xml:"MsgNmId,omitempty"`

	// Specifies the identification of the message that will be used to provide additional details.
	MessageIdentification *Max35Text `xml:"MsgId,omitempty"`
}

func (m *MessageIdentification2) SetMessageNameIdentification(value string) {
	m.MessageNameIdentification = (*Max35Text)(&value)
}

func (m *MessageIdentification2) SetMessageIdentification(value string) {
	m.MessageIdentification = (*Max35Text)(&value)
}
