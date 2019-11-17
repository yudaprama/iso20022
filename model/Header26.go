package model

// Set of characteristics related to the protocol after a rejection.
type Header26 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction9Code `xml:"MsgFctn"`

	// Version of the acquirer protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Number `xml:"XchgId,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification53 `xml:"InitgPty,omitempty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification94 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability5 `xml:"Tracblt,omitempty"`
}

func (h *Header26) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction9Code)(&value)
}

func (h *Header26) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header26) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Number)(&value)
}

func (h *Header26) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header26) AddInitiatingParty() *GenericIdentification53 {
	h.InitiatingParty = new(GenericIdentification53)
	return h.InitiatingParty
}

func (h *Header26) AddRecipientParty() *GenericIdentification94 {
	h.RecipientParty = new(GenericIdentification94)
	return h.RecipientParty
}

func (h *Header26) AddTraceability() *Traceability5 {
	newValue := new(Traceability5)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}
