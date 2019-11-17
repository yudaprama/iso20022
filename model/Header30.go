package model

// Set of characteristics related to the protocol.
type Header30 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction10Code `xml:"MsgFctn"`

	// Version of the acquirer protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Number `xml:"XchgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification53 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification94 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability5 `xml:"Tracblt,omitempty"`
}

func (h *Header30) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction10Code)(&value)
}

func (h *Header30) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header30) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Number)(&value)
}

func (h *Header30) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header30) AddInitiatingParty() *GenericIdentification53 {
	h.InitiatingParty = new(GenericIdentification53)
	return h.InitiatingParty
}

func (h *Header30) AddRecipientParty() *GenericIdentification94 {
	h.RecipientParty = new(GenericIdentification94)
	return h.RecipientParty
}

func (h *Header30) AddTraceability() *Traceability5 {
	newValue := new(Traceability5)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}
