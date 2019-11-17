package model

// Message reference of relevance to the present message.
type MessageReference struct {

	// Business reference of the present message assigned by the party issuing the message. This reference must be unique amongst all messages of the same name sent by the same party.
	Reference *Max35Text `xml:"Ref"`
}

func (m *MessageReference) SetReference(value string) {
	m.Reference = (*Max35Text)(&value)
}
