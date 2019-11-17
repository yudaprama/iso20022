package model

// Set of characteristics related to the protocol.
type Header1 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction1Code `xml:"MsgFctn"`

	// Version of the acquirer protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification32 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification32 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability1 `xml:"Tracblt,omitempty"`
}

func (h *Header1) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction1Code)(&value)
}

func (h *Header1) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header1) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header1) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header1) AddInitiatingParty() *GenericIdentification32 {
	h.InitiatingParty = new(GenericIdentification32)
	return h.InitiatingParty
}

func (h *Header1) AddRecipientParty() *GenericIdentification32 {
	h.RecipientParty = new(GenericIdentification32)
	return h.RecipientParty
}

func (h *Header1) AddTraceability() *Traceability1 {
	newValue := new(Traceability1)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}
