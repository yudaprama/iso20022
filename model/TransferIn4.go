package model

// Information about the confirmation of a transfer in transaction.
type TransferIn4 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails *Transfer7 `xml:"TrfDtls"`

	// Information related to the financial instrument received.
	FinancialInstrumentDetails *FinancialInstrument13 `xml:"FinInstrmDtls"`

	// Information related to the account into which the financial instrument was received.
	AccountDetails *InvestmentAccount22 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation4 `xml:"SttlmDtls"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn4) AddTransferDetails() *Transfer7 {
	t.TransferDetails = new(Transfer7)
	return t.TransferDetails
}

func (t *TransferIn4) AddFinancialInstrumentDetails() *FinancialInstrument13 {
	t.FinancialInstrumentDetails = new(FinancialInstrument13)
	return t.FinancialInstrumentDetails
}

func (t *TransferIn4) AddAccountDetails() *InvestmentAccount22 {
	t.AccountDetails = new(InvestmentAccount22)
	return t.AccountDetails
}

func (t *TransferIn4) AddSettlementDetails() *DeliverInformation4 {
	t.SettlementDetails = new(DeliverInformation4)
	return t.SettlementDetails
}

func (t *TransferIn4) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
