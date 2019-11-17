package model

// Information about a transfer in transaction.
type TransferIn7 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer22 `xml:"TrfDtls"`

	// Information related to the account into which the financial instrument is to be received.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation9 `xml:"SttlmDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn7) AddTransferDetails() *Transfer22 {
	newValue := new(Transfer22)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferIn7) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferIn7) AddSettlementDetails() *DeliverInformation9 {
	t.SettlementDetails = new(DeliverInformation9)
	return t.SettlementDetails
}

func (t *TransferIn7) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
