package model

// Information about a transfer out transaction.
type TransferOut2 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails *Transfer1 `xml:"TrfDtls"`

	// Information related to the financial instrument to be withdrawn.
	FinancialInstrumentDetails *FinancialInstrument3 `xml:"FinInstrmDtls"`

	// Information related to the account from which the financial instrument is to be withdrawn.
	AccountDetails *InvestmentAccount10 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation1 `xml:"SttlmDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut2) AddTransferDetails() *Transfer1 {
	t.TransferDetails = new(Transfer1)
	return t.TransferDetails
}

func (t *TransferOut2) AddFinancialInstrumentDetails() *FinancialInstrument3 {
	t.FinancialInstrumentDetails = new(FinancialInstrument3)
	return t.FinancialInstrumentDetails
}

func (t *TransferOut2) AddAccountDetails() *InvestmentAccount10 {
	t.AccountDetails = new(InvestmentAccount10)
	return t.AccountDetails
}

func (t *TransferOut2) AddSettlementDetails() *ReceiveInformation1 {
	t.SettlementDetails = new(ReceiveInformation1)
	return t.SettlementDetails
}

func (t *TransferOut2) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
