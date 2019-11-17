package model

// Provide information about the type of request or instruction which triggered this confirmation.
type AccountManagementConfirmation1 struct {

	// Specifies if the confirmation message applies to an account opening,  an account modification request or to a get account details.
	ConfirmationType *AccountManagementType2Code `xml:"ConfTp"`

	// Unique and unambiguous identifier of the account opening or modification instruction at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`
}

func (a *AccountManagementConfirmation1) SetConfirmationType(value string) {
	a.ConfirmationType = (*AccountManagementType2Code)(&value)
}

func (a *AccountManagementConfirmation1) SetAccountApplicationIdentification(value string) {
	a.AccountApplicationIdentification = (*Max35Text)(&value)
}
