package model

// Information about a transfer in transaction.
type TransferIn2 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails *Transfer3 `xml:"TrfDtls"`

	// Information related to the financial instrument to be received.
	FinancialInstrumentDetails *FinancialInstrument3 `xml:"FinInstrmDtls"`

	// Information related to the account into which the financial instrument is to be received.
	AccountDetails *InvestmentAccount10 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation1 `xml:"SttlmDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn2) AddTransferDetails() *Transfer3 {
	t.TransferDetails = new(Transfer3)
	return t.TransferDetails
}

func (t *TransferIn2) AddFinancialInstrumentDetails() *FinancialInstrument3 {
	t.FinancialInstrumentDetails = new(FinancialInstrument3)
	return t.FinancialInstrumentDetails
}

func (t *TransferIn2) AddAccountDetails() *InvestmentAccount10 {
	t.AccountDetails = new(InvestmentAccount10)
	return t.AccountDetails
}

func (t *TransferIn2) AddSettlementDetails() *DeliverInformation1 {
	t.SettlementDetails = new(DeliverInformation1)
	return t.SettlementDetails
}

func (t *TransferIn2) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
