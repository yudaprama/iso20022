package model

// Set of characteristics related to the protocol.
type Header17 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction6Code `xml:"MsgFctn"`

	// Version of the acquirer to issuer protocol specifications
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	// Number of retransmission of the message. Incremented by one for each retransmission.
	ReTransmissionCounter *Max3NumericText `xml:"ReTrnsmssnCntr,omitempty"`

	// Date and time at which the message was sent.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification73 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification73 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability3 `xml:"Tracblt,omitempty"`
}

func (h *Header17) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction6Code)(&value)
}

func (h *Header17) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header17) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header17) SetReTransmissionCounter(value string) {
	h.ReTransmissionCounter = (*Max3NumericText)(&value)
}

func (h *Header17) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header17) AddInitiatingParty() *GenericIdentification73 {
	h.InitiatingParty = new(GenericIdentification73)
	return h.InitiatingParty
}

func (h *Header17) AddRecipientParty() *GenericIdentification73 {
	h.RecipientParty = new(GenericIdentification73)
	return h.RecipientParty
}

func (h *Header17) AddTraceability() *Traceability3 {
	newValue := new(Traceability3)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}
