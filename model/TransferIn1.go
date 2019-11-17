package model

// Information about the confirmation of a transfer in transaction.
type TransferIn1 struct {

	// General information related to the transfer of a financial instrument.
	TransferDetails *Transfer4 `xml:"TrfDtls"`

	// Information related to the financial instrument received.
	FinancialInstrumentDetails *FinancialInstrument3 `xml:"FinInstrmDtls"`

	// Information related to the account into which the financial instrument was received.
	AccountDetails *InvestmentAccount10 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation2 `xml:"SttlmDtls"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn1) AddTransferDetails() *Transfer4 {
	t.TransferDetails = new(Transfer4)
	return t.TransferDetails
}

func (t *TransferIn1) AddFinancialInstrumentDetails() *FinancialInstrument3 {
	t.FinancialInstrumentDetails = new(FinancialInstrument3)
	return t.FinancialInstrumentDetails
}

func (t *TransferIn1) AddAccountDetails() *InvestmentAccount10 {
	t.AccountDetails = new(InvestmentAccount10)
	return t.AccountDetails
}

func (t *TransferIn1) AddSettlementDetails() *DeliverInformation2 {
	t.SettlementDetails = new(DeliverInformation2)
	return t.SettlementDetails
}

func (t *TransferIn1) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
