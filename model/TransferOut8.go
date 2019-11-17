package model

// Information about the confirmation of a transfer out transaction.
type TransferOut8 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer14 `xml:"TrfDtls"`

	// Information related to the account from which the financial instrument was withdrawn.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation8 `xml:"SttlmDtls"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut8) AddTransferDetails() *Transfer14 {
	newValue := new(Transfer14)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferOut8) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferOut8) AddSettlementDetails() *ReceiveInformation8 {
	t.SettlementDetails = new(ReceiveInformation8)
	return t.SettlementDetails
}

func (t *TransferOut8) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
