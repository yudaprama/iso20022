package model

// Information about a transfer out transaction.
type TransferOut5 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails *Transfer9 `xml:"TrfDtls"`

	// Information related to the financial instrument to be withdrawn.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Information related to the account from which the financial instrument is to be withdrawn.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation3 `xml:"SttlmDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut5) AddTransferDetails() *Transfer9 {
	t.TransferDetails = new(Transfer9)
	return t.TransferDetails
}

func (t *TransferOut5) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *TransferOut5) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferOut5) AddSettlementDetails() *ReceiveInformation3 {
	t.SettlementDetails = new(ReceiveInformation3)
	return t.SettlementDetails
}

func (t *TransferOut5) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
