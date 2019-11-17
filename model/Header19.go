package model

// Set of characteristics related to the protocol.
type Header19 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction6Code `xml:"MsgFctn"`

	// Identifies the type of process related to the message which has to be reversed.
	OriginalMessageFunction *MessageFunction6Code `xml:"OrgnlMsgFctn,omitempty"`

	// Version of the acquirer to issuer protocol specifications
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	// Number of retransmission of the message. Incremented by 1 for each retransmission.
	ReTransmissionCounter *Max3NumericText `xml:"ReTrnsmssnCntr,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification73 `xml:"InitgPty,omitempty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification73 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability3 `xml:"Tracblt,omitempty"`
}

func (h *Header19) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction6Code)(&value)
}

func (h *Header19) SetOriginalMessageFunction(value string) {
	h.OriginalMessageFunction = (*MessageFunction6Code)(&value)
}

func (h *Header19) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header19) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header19) SetReTransmissionCounter(value string) {
	h.ReTransmissionCounter = (*Max3NumericText)(&value)
}

func (h *Header19) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header19) AddInitiatingParty() *GenericIdentification73 {
	h.InitiatingParty = new(GenericIdentification73)
	return h.InitiatingParty
}

func (h *Header19) AddRecipientParty() *GenericIdentification73 {
	h.RecipientParty = new(GenericIdentification73)
	return h.RecipientParty
}

func (h *Header19) AddTraceability() *Traceability3 {
	newValue := new(Traceability3)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}
