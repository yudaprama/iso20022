package model

// Provides further details on the message.
type GroupHeader58 struct {

	// Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message.
	// Usage: The account servicing institution has to make sure that MessageIdentification is unique per account owner for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Party authorised by the account owner to receive information about movements on the account.
	// Usage: MessageRecipient should only be identified when different from the account owner.
	MessageRecipient *PartyIdentification43 `xml:"MsgRcpt,omitempty"`

	// Provides details on the page number of the message.
	//
	// Usage: The pagination of the message is only allowed when agreed between the parties.
	MessagePagination *Pagination `xml:"MsgPgntn,omitempty"`

	// Unique identification, as assigned by the original requestor, to unambiguously identify the business query message.
	OriginalBusinessQuery *OriginalBusinessQuery1 `xml:"OrgnlBizQry,omitempty"`

	// Further details of the message.
	AdditionalInformation *Max500Text `xml:"AddtlInf,omitempty"`
}

func (g *GroupHeader58) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader58) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader58) AddMessageRecipient() *PartyIdentification43 {
	g.MessageRecipient = new(PartyIdentification43)
	return g.MessageRecipient
}

func (g *GroupHeader58) AddMessagePagination() *Pagination {
	g.MessagePagination = new(Pagination)
	return g.MessagePagination
}

func (g *GroupHeader58) AddOriginalBusinessQuery() *OriginalBusinessQuery1 {
	g.OriginalBusinessQuery = new(OriginalBusinessQuery1)
	return g.OriginalBusinessQuery
}

func (g *GroupHeader58) SetAdditionalInformation(value string) {
	g.AdditionalInformation = (*Max500Text)(&value)
}
