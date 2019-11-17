package model

// Information about a transfer out transaction.
type TransferOut9 struct {

	// Requested date at which the instructing party places the transfer instruction.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer20 `xml:"TrfDtls"`

	// Information related to the account from which the financial instrument is to be withdrawn.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation9 `xml:"SttlmDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut9) AddRequestedTransferDate() *DateFormat1Choice {
	t.RequestedTransferDate = new(DateFormat1Choice)
	return t.RequestedTransferDate
}

func (t *TransferOut9) AddTransferDetails() *Transfer20 {
	newValue := new(Transfer20)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferOut9) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferOut9) AddSettlementDetails() *ReceiveInformation9 {
	t.SettlementDetails = new(ReceiveInformation9)
	return t.SettlementDetails
}

func (t *TransferOut9) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
