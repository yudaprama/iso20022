package model

// Information about the message reference of the message for which the status is requested and the business reference of the transfer instruction.
type MessageAndBusinessReference8 struct {

	// Reference to the message or communication that was previously sent.
	Reference *References48Choice `xml:"Ref,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identification of a transfer, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Investment account information of the transfer message for which the status is requested.
	InvestmentAccountDetails *InvestmentAccount57 `xml:"InvstmtAcctDtls,omitempty"`
}

func (m *MessageAndBusinessReference8) AddReference() *References48Choice {
	m.Reference = new(References48Choice)
	return m.Reference
}

func (m *MessageAndBusinessReference8) SetMasterReference(value string) {
	m.MasterReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference8) SetTransferReference(value string) {
	m.TransferReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference8) AddClientReference() *AdditionalReference7 {
	m.ClientReference = new(AdditionalReference7)
	return m.ClientReference
}

func (m *MessageAndBusinessReference8) SetCancellationReference(value string) {
	m.CancellationReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference8) AddInvestmentAccountDetails() *InvestmentAccount57 {
	m.InvestmentAccountDetails = new(InvestmentAccount57)
	return m.InvestmentAccountDetails
}
