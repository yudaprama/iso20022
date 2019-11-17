package model

// Information about the confirmation of a transfer in transaction.
type TransferIn8 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer25 `xml:"TrfDtls"`

	// Information related to the account into which the financial instrument was received.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation11 `xml:"SttlmDtls,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn8) AddTransferDetails() *Transfer25 {
	newValue := new(Transfer25)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferIn8) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferIn8) AddSettlementDetails() *DeliverInformation11 {
	t.SettlementDetails = new(DeliverInformation11)
	return t.SettlementDetails
}

func (t *TransferIn8) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
