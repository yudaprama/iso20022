package model

// Information about the confirmation of a transfer out transaction.
type TransferOut1 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails *Transfer2 `xml:"TrfDtls"`

	// Information related to the financial instrument withdrawn.
	FinancialInstrumentDetails *FinancialInstrument3 `xml:"FinInstrmDtls"`

	// Information related to the account from which the financial instrument was withdrawn.
	AccountDetails *InvestmentAccount10 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation2 `xml:"SttlmDtls"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut1) AddTransferDetails() *Transfer2 {
	t.TransferDetails = new(Transfer2)
	return t.TransferDetails
}

func (t *TransferOut1) AddFinancialInstrumentDetails() *FinancialInstrument3 {
	t.FinancialInstrumentDetails = new(FinancialInstrument3)
	return t.FinancialInstrumentDetails
}

func (t *TransferOut1) AddAccountDetails() *InvestmentAccount10 {
	t.AccountDetails = new(InvestmentAccount10)
	return t.AccountDetails
}

func (t *TransferOut1) AddSettlementDetails() *ReceiveInformation2 {
	t.SettlementDetails = new(ReceiveInformation2)
	return t.SettlementDetails
}

func (t *TransferOut1) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
