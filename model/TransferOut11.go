package model

// Information about a transfer out transaction.
type TransferOut11 struct {

	// Requested date at which the instructing party places the transfer instruction.
	RequestedTransferDate *DateFormat1Choice `xml:"ReqdTrfDt,omitempty"`

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Details of the transfer and cancellation.
	TransferAndReferences []*TransferOut13 `xml:"TrfAndRefs"`

	// Information related to the account from which the financial instrument is to be withdrawn.
	AccountDetails *InvestmentAccount40 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation13 `xml:"SttlmDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut11) AddRequestedTransferDate() *DateFormat1Choice {
	t.RequestedTransferDate = new(DateFormat1Choice)
	return t.RequestedTransferDate
}

func (t *TransferOut11) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferOut11) AddTransferAndReferences() *TransferOut13 {
	newValue := new(TransferOut13)
	t.TransferAndReferences = append(t.TransferAndReferences, newValue)
	return newValue
}

func (t *TransferOut11) AddAccountDetails() *InvestmentAccount40 {
	t.AccountDetails = new(InvestmentAccount40)
	return t.AccountDetails
}

func (t *TransferOut11) AddSettlementDetails() *ReceiveInformation13 {
	t.SettlementDetails = new(ReceiveInformation13)
	return t.SettlementDetails
}

func (t *TransferOut11) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
