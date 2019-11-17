package model

// Presence condition of a message item.
type MessageItemCondition1 struct {

	// Unique identification of the message and the message item.
	ItemIdentification *Max140Text `xml:"ItmId"`

	// Condition of presence of the message item.
	Condition *MessageItemCondition1Code `xml:"Cond"`

	// Value to be used for the message item.
	Value []*Max140Text `xml:"Val,omitempty"`
}

func (m *MessageItemCondition1) SetItemIdentification(value string) {
	m.ItemIdentification = (*Max140Text)(&value)
}

func (m *MessageItemCondition1) SetCondition(value string) {
	m.Condition = (*MessageItemCondition1Code)(&value)
}

func (m *MessageItemCondition1) AddValue(value string) {
	m.Value = append(m.Value, (*Max140Text)(&value))
}
