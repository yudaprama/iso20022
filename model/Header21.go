package model

// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
type Header21 struct {

	// Identifies the type of process related to the message.
	MessageFunction *ATMMessageFunction1 `xml:"MsgFctn"`

	// Version of the ATM protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	//  Retransmission counter of this advice, 0 for the first transmission.
	ReTransmissionCounter *Number `xml:"ReTrnsmssnCntr,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *Max35Text `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *Max35Text `xml:"RcptPty,omitempty"`

	// State of the sender of the message inside the process flow.
	ProcessState *Max35Text `xml:"PrcStat,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability4 `xml:"Tracblt,omitempty"`
}

func (h *Header21) AddMessageFunction() *ATMMessageFunction1 {
	h.MessageFunction = new(ATMMessageFunction1)
	return h.MessageFunction
}

func (h *Header21) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header21) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header21) SetReTransmissionCounter(value string) {
	h.ReTransmissionCounter = (*Number)(&value)
}

func (h *Header21) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header21) SetInitiatingParty(value string) {
	h.InitiatingParty = (*Max35Text)(&value)
}

func (h *Header21) SetRecipientParty(value string) {
	h.RecipientParty = (*Max35Text)(&value)
}

func (h *Header21) SetProcessState(value string) {
	h.ProcessState = (*Max35Text)(&value)
}

func (h *Header21) AddTraceability() *Traceability4 {
	newValue := new(Traceability4)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}
