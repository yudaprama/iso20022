package model

// Information about the modification of an account.
type InvestmentAccountModification2 struct {

	// Reason for the modification to the investment account information.
	ModificationReason *Max350Text `xml:"ModRsn,omitempty"`

	// Unique and unambiguous identifier of the account modification request.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Unique and unambiguous investor's identification of a transfer.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transaction, for example, a transfer, as allocated by the counterparty.
	CounterpartyReference *AdditionalReference6 `xml:"CtrPtyRef,omitempty"`

	// Account to which the account opening is related.
	ExistingAccountIdentification *Max35Text `xml:"ExstgAcctId,omitempty"`
}

func (i *InvestmentAccountModification2) SetModificationReason(value string) {
	i.ModificationReason = (*Max350Text)(&value)
}

func (i *InvestmentAccountModification2) SetAccountApplicationIdentification(value string) {
	i.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountModification2) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentAccountModification2) AddCounterpartyReference() *AdditionalReference6 {
	i.CounterpartyReference = new(AdditionalReference6)
	return i.CounterpartyReference
}

func (i *InvestmentAccountModification2) SetExistingAccountIdentification(value string) {
	i.ExistingAccountIdentification = (*Max35Text)(&value)
}
