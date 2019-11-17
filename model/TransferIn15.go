package model

// Information about a transfer in transaction.
type TransferIn15 struct {

	// Requested date at which the instructing party places the transfer instruction.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Details of the transfer and cancellation.
	TransferAndReferences []*TransferIn16 `xml:"TrfAndRefs"`

	// Information related to the account into which the financial instrument is to be received.
	AccountDetails *InvestmentAccount56 `xml:"AcctDtls"`

	// Information related to the delivering side of the transfer.
	SettlementDetails *DeliverInformation16 `xml:"SttlmDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferIn15) AddRequestedTransferDate() *DateFormat1Choice {
	t.RequestedTransferDate = new(DateFormat1Choice)
	return t.RequestedTransferDate
}

func (t *TransferIn15) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferIn15) AddTransferAndReferences() *TransferIn16 {
	newValue := new(TransferIn16)
	t.TransferAndReferences = append(t.TransferAndReferences, newValue)
	return newValue
}

func (t *TransferIn15) AddAccountDetails() *InvestmentAccount56 {
	t.AccountDetails = new(InvestmentAccount56)
	return t.AccountDetails
}

func (t *TransferIn15) AddSettlementDetails() *DeliverInformation16 {
	t.SettlementDetails = new(DeliverInformation16)
	return t.SettlementDetails
}

func (t *TransferIn15) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
