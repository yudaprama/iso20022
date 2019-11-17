package model

// Unique and unambiguous identifier for a message, as assigned by the Sender. This unique identifier can be used for cross-referencing purposes in subsequent messages.
type MessageIdentification struct {

	// String of characters that uniquely identifies a message.
	Identification *Max35Text `xml:"Id"`
}

func (m *MessageIdentification) SetIdentification(value string) {
	m.Identification = (*Max35Text)(&value)
}
