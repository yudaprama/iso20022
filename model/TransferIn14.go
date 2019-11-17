package model

// Information about the confirmation of a transfer in transaction.
type TransferIn14 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer33 `xml:"TrfDtls"`

	// Information related to the account into which the financial instrument was received.
	AccountDetails *InvestmentAccount56 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation17 `xml:"SttlmDtls,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn14) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferIn14) AddTransferDetails() *Transfer33 {
	newValue := new(Transfer33)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferIn14) AddAccountDetails() *InvestmentAccount56 {
	t.AccountDetails = new(InvestmentAccount56)
	return t.AccountDetails
}

func (t *TransferIn14) AddSettlementDetails() *DeliverInformation17 {
	t.SettlementDetails = new(DeliverInformation17)
	return t.SettlementDetails
}

func (t *TransferIn14) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
