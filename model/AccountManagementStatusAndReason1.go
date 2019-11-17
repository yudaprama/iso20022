package model

// Status report of a account opening instruction or account modification instruction that was previously received.
type AccountManagementStatusAndReason1 struct {

	// Status of the account opening instruction or account modification instruction.
	Status *AccountManagementStatus1Code `xml:"Sts"`

	// Status of the order is rejected.
	Rejected *RejectedStatus5 `xml:"Rjctd"`

	// Unique and unambiguous identifier of the account opening or modification instruction at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`
}

func (a *AccountManagementStatusAndReason1) SetStatus(value string) {
	a.Status = (*AccountManagementStatus1Code)(&value)
}

func (a *AccountManagementStatusAndReason1) AddRejected() *RejectedStatus5 {
	a.Rejected = new(RejectedStatus5)
	return a.Rejected
}

func (a *AccountManagementStatusAndReason1) SetAccountApplicationIdentification(value string) {
	a.AccountApplicationIdentification = (*Max35Text)(&value)
}
