package model

// Status report of an account opening instruction or account modification instruction that was previously received.
type AccountManagementStatusAndReason2 struct {

	// Status of the account opening instruction or account modification instruction.
	Status *Status12Choice `xml:"Sts"`

	// Unique and unambiguous identifier of the account opening or modification instruction at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`
}

func (a *AccountManagementStatusAndReason2) AddStatus() *Status12Choice {
	a.Status = new(Status12Choice)
	return a.Status
}

func (a *AccountManagementStatusAndReason2) SetAccountApplicationIdentification(value string) {
	a.AccountApplicationIdentification = (*Max35Text)(&value)
}
