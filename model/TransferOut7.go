package model

// Information about a transfer out transaction.
type TransferOut7 struct {

	// Requested date at which the instructing party places the transfer instruction.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer12 `xml:"TrfDtls"`

	// Information related to the account from which the financial instrument is to be withdrawn.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation7 `xml:"SttlmDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut7) AddRequestedTransferDate() *DateFormat1Choice {
	t.RequestedTransferDate = new(DateFormat1Choice)
	return t.RequestedTransferDate
}

func (t *TransferOut7) AddTransferDetails() *Transfer12 {
	newValue := new(Transfer12)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferOut7) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferOut7) AddSettlementDetails() *ReceiveInformation7 {
	t.SettlementDetails = new(ReceiveInformation7)
	return t.SettlementDetails
}

func (t *TransferOut7) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
