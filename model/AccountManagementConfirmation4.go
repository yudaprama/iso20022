package model

// Information about the type of request or instruction.
type AccountManagementConfirmation4 struct {

	// Specifies the confirmation type.
	ConfirmationType *ConfirmationType1Choice `xml:"ConfTp"`

	// Unique and unambiguous identifier of the account opening or modification instruction at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Unique and unambiguous identification of a transaction, for example, a transfer, as assigned by the investor or account owner.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transaction, for example, a transfer, as allocated by the counterparty.
	CounterpartyReference *AdditionalReference6 `xml:"CtrPtyRef,omitempty"`

	// Account to which the account opening is related.
	ExistingAccountIdentification []*Account23 `xml:"ExstgAcctId,omitempty"`
}

func (a *AccountManagementConfirmation4) AddConfirmationType() *ConfirmationType1Choice {
	a.ConfirmationType = new(ConfirmationType1Choice)
	return a.ConfirmationType
}

func (a *AccountManagementConfirmation4) SetAccountApplicationIdentification(value string) {
	a.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (a *AccountManagementConfirmation4) SetClientReference(value string) {
	a.ClientReference = (*Max35Text)(&value)
}

func (a *AccountManagementConfirmation4) AddCounterpartyReference() *AdditionalReference6 {
	a.CounterpartyReference = new(AdditionalReference6)
	return a.CounterpartyReference
}

func (a *AccountManagementConfirmation4) AddExistingAccountIdentification() *Account23 {
	newValue := new(Account23)
	a.ExistingAccountIdentification = append(a.ExistingAccountIdentification, newValue)
	return newValue
}
