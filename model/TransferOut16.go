package model

// Information about the confirmation of a transfer out transaction.
type TransferOut16 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// General information related to the transfer of a financial instrument.
	TransferDetails []*Transfer31 `xml:"TrfDtls"`

	// Information related to the account from which the financial instrument was withdrawn.
	AccountDetails *InvestmentAccount54 `xml:"AcctDtls"`

	// Information related to the receiving side of the transfer.
	SettlementDetails *ReceiveInformation17 `xml:"SttlmDtls,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (t *TransferOut16) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferOut16) AddTransferDetails() *Transfer31 {
	newValue := new(Transfer31)
	t.TransferDetails = append(t.TransferDetails, newValue)
	return newValue
}

func (t *TransferOut16) AddAccountDetails() *InvestmentAccount54 {
	t.AccountDetails = new(InvestmentAccount54)
	return t.AccountDetails
}

func (t *TransferOut16) AddSettlementDetails() *ReceiveInformation17 {
	t.SettlementDetails = new(ReceiveInformation17)
	return t.SettlementDetails
}

func (t *TransferOut16) AddExtension() *Extension1 {
	newValue := new(Extension1)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
