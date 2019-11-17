package model

// Set of elements used to provide further details of the message.
type GroupHeader42 struct {

	// Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message.
	// Usage: The account servicing institution has to make sure that MessageIdentification is unique per account owner for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Party authorised by the account owner to receive information about movements on the account.
	// Usage: MessageRecipient should only be identified when different from the account owner.
	MessageRecipient *PartyIdentification32 `xml:"MsgRcpt,omitempty"`

	// Set of elements used to provide details on the page number of the message.
	//
	// Usage: The pagination of the message is only allowed when agreed between the parties.
	MessagePagination *Pagination `xml:"MsgPgntn,omitempty"`

	// Further details of the message.
	AdditionalInformation *Max500Text `xml:"AddtlInf,omitempty"`
}

func (g *GroupHeader42) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader42) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader42) AddMessageRecipient() *PartyIdentification32 {
	g.MessageRecipient = new(PartyIdentification32)
	return g.MessageRecipient
}

func (g *GroupHeader42) AddMessagePagination() *Pagination {
	g.MessagePagination = new(Pagination)
	return g.MessagePagination
}

func (g *GroupHeader42) SetAdditionalInformation(value string) {
	g.AdditionalInformation = (*Max500Text)(&value)
}
