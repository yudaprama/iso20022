package model

// Status information.
type AccountManagementStatusAndReason4 struct {

	// Status of the account opening instruction or account modification instruction.
	Status *Status20Choice `xml:"Sts"`

	// Unique and unambiguous identifier of the account opening or modification instruction at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Account to which the account opening is related.
	ExistingAccountIdentification *Max35Text `xml:"ExstgAcctId,omitempty"`
}

func (a *AccountManagementStatusAndReason4) AddStatus() *Status20Choice {
	a.Status = new(Status20Choice)
	return a.Status
}

func (a *AccountManagementStatusAndReason4) SetAccountApplicationIdentification(value string) {
	a.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (a *AccountManagementStatusAndReason4) SetExistingAccountIdentification(value string) {
	a.ExistingAccountIdentification = (*Max35Text)(&value)
}
