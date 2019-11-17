package model

// Identifies a message by a unique identifier and the date and time when the message was created by the sender.
type MessageIdentification5 struct {

	// Point to point reference, as assigned by the assigner, and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The assigner has to make sure that MessageIdentification is unique per assignee for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm,omitempty"`

	// Identifies the first agent in the identification chain, following the payment initiating party.
	FirstAgent *BranchAndFinancialInstitutionIdentification5 `xml:"FrstAgt,omitempty"`
}

func (m *MessageIdentification5) SetMessageIdentification(value string) {
	m.MessageIdentification = (*Max35Text)(&value)
}

func (m *MessageIdentification5) SetCreationDateTime(value string) {
	m.CreationDateTime = (*ISODateTime)(&value)
}

func (m *MessageIdentification5) AddFirstAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.FirstAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.FirstAgent
}
