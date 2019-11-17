package model

// Information about the confirmation of a transfer out transaction.
type TransferOut12 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer28 `xml:"TrfDtls"`

	// Information related to the account from which the financial instrument was withdrawn.
	AccountDetails *InvestmentAccount40 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation12 `xml:"SttlmDtls,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut12) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferOut12) AddTransferDetails() *Transfer28 {
	newValue := new(Transfer28)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferOut12) AddAccountDetails() *InvestmentAccount40 {
	t.AccountDetails = new(InvestmentAccount40)
	return t.AccountDetails
}

func (t *TransferOut12) AddSettlementDetails() *ReceiveInformation12 {
	t.SettlementDetails = new(ReceiveInformation12)
	return t.SettlementDetails
}

func (t *TransferOut12) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
