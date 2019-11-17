package model

// Information about the modification of an account.
type InvestmentAccountModification3 struct {

	// Reason for the modification to the investment account information.
	ModificationReason *Max350Text `xml:"ModRsn,omitempty"`

	// Unique and unambiguous identifier of the account modification request.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Unique and unambiguous investor's identification of a transfer.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transaction, for example, a transfer, as allocated by the counterparty.
	CounterpartyReference *AdditionalReference6 `xml:"CtrPtyRef,omitempty"`

	// Account to which the account opening is related.
	ExistingAccountIdentification []*Account23 `xml:"ExstgAcctId,omitempty"`
}

func (i *InvestmentAccountModification3) SetModificationReason(value string) {
	i.ModificationReason = (*Max350Text)(&value)
}

func (i *InvestmentAccountModification3) SetAccountApplicationIdentification(value string) {
	i.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccountModification3) SetClientReference(value string) {
	i.ClientReference = (*Max35Text)(&value)
}

func (i *InvestmentAccountModification3) AddCounterpartyReference() *AdditionalReference6 {
	i.CounterpartyReference = new(AdditionalReference6)
	return i.CounterpartyReference
}

func (i *InvestmentAccountModification3) AddExistingAccountIdentification() *Account23 {
	newValue := new(Account23)
	i.ExistingAccountIdentification = append(i.ExistingAccountIdentification, newValue)
	return newValue
}
