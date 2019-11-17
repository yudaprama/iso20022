package model

// Information about the confirmation of a transfer in transaction.
type TransferIn12 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer29 `xml:"TrfDtls"`

	// Information related to the account into which the financial instrument was received.
	AccountDetails *InvestmentAccount40 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation14 `xml:"SttlmDtls,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn12) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferIn12) AddTransferDetails() *Transfer29 {
	newValue := new(Transfer29)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferIn12) AddAccountDetails() *InvestmentAccount40 {
	t.AccountDetails = new(InvestmentAccount40)
	return t.AccountDetails
}

func (t *TransferIn12) AddSettlementDetails() *DeliverInformation14 {
	t.SettlementDetails = new(DeliverInformation14)
	return t.SettlementDetails
}

func (t *TransferIn12) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
