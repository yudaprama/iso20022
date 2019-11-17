package model

// Set of elements providing further details on the message.
type GroupHeader23 struct {

	// Point to point reference assigned by the account servicing institution and sent to the account owner to unambiguously identify the message.
	//
	// Usage: The account servicing institution has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created by the account servicer.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Party that is entitled by the account owner to receive information about movements in the account.
	//
	// Guideline : MessageRecipient should only be identified when different from the account owner.
	MessageRecipient *PartyIdentification8 `xml:"MsgRcpt,omitempty"`

	// Pagination of the message.
	//
	// Usage: the pagination of the message is only allowed when agreed between the parties.
	MessagePagination *Pagination `xml:"MsgPgntn,omitempty"`

	// Further details on the message.
	AdditionalInformation *Max500Text `xml:"AddtlInf,omitempty"`
}

func (g *GroupHeader23) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader23) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader23) AddMessageRecipient() *PartyIdentification8 {
	g.MessageRecipient = new(PartyIdentification8)
	return g.MessageRecipient
}

func (g *GroupHeader23) AddMessagePagination() *Pagination {
	g.MessagePagination = new(Pagination)
	return g.MessagePagination
}

func (g *GroupHeader23) SetAdditionalInformation(value string) {
	g.AdditionalInformation = (*Max500Text)(&value)
}
