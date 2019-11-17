package model

// Information about the confirmation of a transfer out transaction.
type TransferOut10 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer24 `xml:"TrfDtls"`

	// Information related to the account from which the financial instrument was withdrawn.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation11 `xml:"SttlmDtls,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut10) AddTransferDetails() *Transfer24 {
	newValue := new(Transfer24)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferOut10) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferOut10) AddSettlementDetails() *ReceiveInformation11 {
	t.SettlementDetails = new(ReceiveInformation11)
	return t.SettlementDetails
}

func (t *TransferOut10) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
