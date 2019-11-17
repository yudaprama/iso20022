package model

// Information about the references of an account management instruction message.
type AccountManagementMessageReference3 struct {

	// Reference to a linked message.
	LinkedReference *LinkedMessage3Choice `xml:"LkdRef,omitempty"`

	// Specifies if the status request refers to an earlier account opening or modification instruction message.
	StatusRequestType *AccountManagementType1Code `xml:"StsReqTp"`

	// Unique and unambiguous identifier of the account opening or account modification instruction at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Account to which the account opening is related.
	ExistingAccountIdentification *Max35Text `xml:"ExstgAcctId,omitempty"`

	// Account information for which the status of an account management instruction is requested.
	InvestmentAccount *InvestmentAccount53 `xml:"InvstmtAcct,omitempty"`
}

func (a *AccountManagementMessageReference3) AddLinkedReference() *LinkedMessage3Choice {
	a.LinkedReference = new(LinkedMessage3Choice)
	return a.LinkedReference
}

func (a *AccountManagementMessageReference3) SetStatusRequestType(value string) {
	a.StatusRequestType = (*AccountManagementType1Code)(&value)
}

func (a *AccountManagementMessageReference3) SetAccountApplicationIdentification(value string) {
	a.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (a *AccountManagementMessageReference3) SetExistingAccountIdentification(value string) {
	a.ExistingAccountIdentification = (*Max35Text)(&value)
}

func (a *AccountManagementMessageReference3) AddInvestmentAccount() *InvestmentAccount53 {
	a.InvestmentAccount = new(InvestmentAccount53)
	return a.InvestmentAccount
}
