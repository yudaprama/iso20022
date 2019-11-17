package model

// Information about the references of an account management instruction message.
type AccountManagementMessageReference4 struct {

	// Reference to a linked message.
	LinkedReference *LinkedMessage4Choice `xml:"LkdRef,omitempty"`

	// Type of account management instruction for which the status  is requested or a request to know the status of the account.
	StatusRequestType *AccountManagementType3Code `xml:"StsReqTp"`

	// Unique and unambiguous identifier of the account opening or account modification instruction at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Account to which the account opening is related.
	ExistingAccountIdentification *Account23 `xml:"ExstgAcctId,omitempty"`

	// Account information for which the status of an account management instruction is requested.
	InvestmentAccount *InvestmentAccount53 `xml:"InvstmtAcct,omitempty"`
}

func (a *AccountManagementMessageReference4) AddLinkedReference() *LinkedMessage4Choice {
	a.LinkedReference = new(LinkedMessage4Choice)
	return a.LinkedReference
}

func (a *AccountManagementMessageReference4) SetStatusRequestType(value string) {
	a.StatusRequestType = (*AccountManagementType3Code)(&value)
}

func (a *AccountManagementMessageReference4) SetAccountApplicationIdentification(value string) {
	a.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (a *AccountManagementMessageReference4) AddExistingAccountIdentification() *Account23 {
	a.ExistingAccountIdentification = new(Account23)
	return a.ExistingAccountIdentification
}

func (a *AccountManagementMessageReference4) AddInvestmentAccount() *InvestmentAccount53 {
	a.InvestmentAccount = new(InvestmentAccount53)
	return a.InvestmentAccount
}
