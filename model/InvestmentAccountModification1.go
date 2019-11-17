package model

// Provide information about the reason for the modification and about the application request which triggered this modification.
type InvestmentAccountModification1 struct {

	// Reason for the modification brought to the investment account information.
	ModificationReason *Max350Text `xml:"ModRsn,omitempty"`

	// Unique and unambiguous identifier of the account modification request.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Unique and unambiguous investor's identification of a transfer.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer as allocated by the counterparty.
	CounterpartyReference *AdditionalReference2 `xml:"CtrPtyRef,omitempty"`
}

func (i *InvestmentAccountModification1) SetModificationReason(value string) {
	i.ModificationReason = (*Max350Text)(&value)
}

func (i *InvestmentAccountModification1) SetAccountApplicationIdentification(value string) {
	i.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountModification1) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentAccountModification1) AddCounterpartyReference() *AdditionalReference2 {
	i.CounterpartyReference = new(AdditionalReference2)
	return i.CounterpartyReference
}
