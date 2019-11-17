package model

// Status report of an account opening instruction or account modification instruction that was previously received.
type AccountManagementStatusAndReason3 struct {

	// Status of the account opening instruction or account modification instruction.
	Status *Status14Choice `xml:"Sts"`

	// Unique and unambiguous identifier of the account opening or modification instruction at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`
}

func (a *AccountManagementStatusAndReason3) AddStatus() *Status14Choice {
	a.Status = new(Status14Choice)
	return a.Status
}

func (a *AccountManagementStatusAndReason3) SetAccountApplicationIdentification(value string) {
	a.AccountApplicationIdentification = (*Max35Text)(&value)
}
