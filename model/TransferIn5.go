package model

// Information about a transfer in transaction.
type TransferIn5 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer16 `xml:"TrfDtls"`

	// Information related to the account into which the financial instrument is to be received.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation8 `xml:"SttlmDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn5) AddTransferDetails() *Transfer16 {
	newValue := new(Transfer16)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferIn5) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferIn5) AddSettlementDetails() *DeliverInformation8 {
	t.SettlementDetails = new(DeliverInformation8)
	return t.SettlementDetails
}

func (t *TransferIn5) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
