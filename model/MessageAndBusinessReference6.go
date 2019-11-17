package model

// Information about the message reference of the message for which the status is requested and the business reference of the transfer instruction.
type MessageAndBusinessReference6 struct {

	// Reference to a linked message that was previously sent.
	PreviousReference *AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message sent in a proprietary way or the reference of a system.
	OtherReference *AdditionalReference3 `xml:"OthrRef,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identification of a transfer, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Investment account information of the transfer message for which the status is requested.
	InvestmentAccountDetails *InvestmentAccount22 `xml:"InvstmtAcctDtls,omitempty"`
}

func (m *MessageAndBusinessReference6) AddPreviousReference() *AdditionalReference3 {
	m.PreviousReference = new(AdditionalReference3)
	return m.PreviousReference
}

func (m *MessageAndBusinessReference6) AddOtherReference() *AdditionalReference3 {
	m.OtherReference = new(AdditionalReference3)
	return m.OtherReference
}

func (m *MessageAndBusinessReference6) SetMasterReference(value string) {
	m.MasterReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference6) SetTransferReference(value string) {
	m.TransferReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference6) SetClientReference(value string) {
	m.ClientReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference6) SetCancellationReference(value string) {
	m.CancellationReference = (*Max35Text)(&value)
}

func (m *MessageAndBusinessReference6) AddInvestmentAccountDetails() *InvestmentAccount22 {
	m.InvestmentAccountDetails = new(InvestmentAccount22)
	return m.InvestmentAccountDetails
}
