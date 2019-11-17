package model

// Information about the message reference of the account management instruction message for which the status is requested .
type AccountManagementMessageReference2 struct {

	// Reference to a linked message.
	LinkedReference *LinkedMessage2Choice `xml:"LkdRef,omitempty"`

	// Specifies if the status request refers to an earlier account opening or modification instruction message.
	StatusRequestType *AccountManagementType1Code `xml:"StsReqTp"`

	// Unique and unambiguous identifier of the account opening or account modification instruction at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Account information for which the status of an account management instruction is requested.
	InvestmentAccount *InvestmentAccount45 `xml:"InvstmtAcct,omitempty"`
}

func (a *AccountManagementMessageReference2) AddLinkedReference() *LinkedMessage2Choice {
	a.LinkedReference = new(LinkedMessage2Choice)
	return a.LinkedReference
}

func (a *AccountManagementMessageReference2) SetStatusRequestType(value string) {
	a.StatusRequestType = (*AccountManagementType1Code)(&value)
}

func (a *AccountManagementMessageReference2) SetAccountApplicationIdentification(value string) {
	a.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (a *AccountManagementMessageReference2) AddInvestmentAccount() *InvestmentAccount45 {
	a.InvestmentAccount = new(InvestmentAccount45)
	return a.InvestmentAccount
}
