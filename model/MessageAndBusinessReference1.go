package model

// Information about the message reference of the message for which the status is requested and the business reference of the transfer instruction.
type MessageAndBusinessReference1 struct {

	// Reference to a linked message that was previously sent.
	PreviousReference *AdditionalReference2 `xml:"PrvsRef"`

	// Reference to a linked message sent in a proprietary way or the reference of a system.
	OtherReference *AdditionalReference2 `xml:"OthrRef"`

	// Investment account information of the transfer message for which the status is requested.
	InvestmentAccountDetails *InvestmentAccount10 `xml:"InvstmtAcctDtls,omitempty"`

	// Business reference of the transfer instruction message.
	TransferReference *Max35Text `xml:"TrfRef,omitempty"`
}

func (m *MessageAndBusinessReference1) AddPreviousReference() *AdditionalReference2 {
	m.PreviousReference = new(AdditionalReference2)
	return m.PreviousReference
}

func (m *MessageAndBusinessReference1) AddOtherReference() *AdditionalReference2 {
	m.OtherReference = new(AdditionalReference2)
	return m.OtherReference
}

func (m *MessageAndBusinessReference1) AddInvestmentAccountDetails() *InvestmentAccount10 {
	m.InvestmentAccountDetails = new(InvestmentAccount10)
	return m.InvestmentAccountDetails
}

func (m *MessageAndBusinessReference1) SetTransferReference(value string) {
	m.TransferReference = (*Max35Text)(&value)
}
