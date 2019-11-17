package model

// Information about the message reference of the message for which the status is requested and the business reference of the transfer instruction.
type MessageAndBusinessReference7 struct {

	// Reference to the message or communication that was previously sent.
	Reference *References39Choice `xml:"Ref,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identification of a transfer, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Investment account information of the transfer message for which the status is requested.
	InvestmentAccountDetails *InvestmentAccount40 `xml:"InvstmtAcctDtls,omitempty"`
}

func (m *MessageAndBusinessReference7) AddReference() *References39Choice {
	m.Reference = new(References39Choice)
	return m.Reference
}

func (m *MessageAndBusinessReference7) SetMasterReference(value string) {
	m.MasterReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference7) SetTransferReference(value string) {
	m.TransferReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference7) SetClientReference(value string) {
	m.ClientReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference7) SetCancellationReference(value string) {
	m.CancellationReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference7) AddInvestmentAccountDetails() *InvestmentAccount40 {
	m.InvestmentAccountDetails = new(InvestmentAccount40)
	return m.InvestmentAccountDetails
}
