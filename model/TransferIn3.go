package model

// Information about a transfer in transaction.
type TransferIn3 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails *Transfer6 `xml:"TrfDtls"`

	// Information related to the financial instrument to be received.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Information related to the account into which the financial instrument is to be received.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation3 `xml:"SttlmDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn3) AddTransferDetails() *Transfer6 {
	t.TransferDetails = new(Transfer6)
	return t.TransferDetails
}

func (t *TransferIn3) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *TransferIn3) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferIn3) AddSettlementDetails() *DeliverInformation3 {
	t.SettlementDetails = new(DeliverInformation3)
	return t.SettlementDetails
}

func (t *TransferIn3) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
