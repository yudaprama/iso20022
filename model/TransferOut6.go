package model

// Information about the confirmation of a transfer out transaction.
type TransferOut6 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails *Transfer10 `xml:"TrfDtls"`

	// Information related to the financial instrument withdrawn.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Information related to the account from which the financial instrument was withdrawn.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation4 `xml:"SttlmDtls"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut6) AddTransferDetails() *Transfer10 {
	t.TransferDetails = new(Transfer10)
	return t.TransferDetails
}

func (t *TransferOut6) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *TransferOut6) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferOut6) AddSettlementDetails() *ReceiveInformation4 {
	t.SettlementDetails = new(ReceiveInformation4)
	return t.SettlementDetails
}

func (t *TransferOut6) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
