package model

// Information about the type of request or instruction.
type AccountManagementConfirmation2 struct {

	// Specifies if the confirmation message applies to an account opening,  an account modification request or to a get account details.
	ConfirmationType *AccountManagementType2Code `xml:"ConfTp"`

	// Unique and unambiguous identifier of the account opening or modification instruction at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Unique and unambiguous investor's identification of a transfer.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer as allocated by the counterparty.
	CounterpartyReference *AdditionalReference2 `xml:"CtrPtyRef,omitempty"`
}

func (a *AccountManagementConfirmation2) SetConfirmationType(value string) {
	a.ConfirmationType = (*AccountManagementType2Code)(&value)
}

func (a *AccountManagementConfirmation2) SetAccountApplicationIdentification(value string) {
	a.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (a *AccountManagementConfirmation2) SetClientReference(value string) {
	a.ClientReference = (*Max35Text)(&value)
}

func (a *AccountManagementConfirmation2) AddCounterpartyReference() *AdditionalReference2 {
	a.CounterpartyReference = new(AdditionalReference2)
	return a.CounterpartyReference
}
