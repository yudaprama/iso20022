package model

// Provides details on the account notification.
type AccountNotification4 struct {

	// Unique identification, as assigned by the account owner, to unambiguously identify the account notification.
	Identification *Max35Text `xml:"Id"`

	// Identifies the account to be credited with the incoming amount of money.
	Account *CashAccount16 `xml:"Acct,omitempty"`

	// Party that legally owns the account.
	AccountOwner *Party12Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcr,omitempty"`

	// Identifies the parent account of the account to be credited with the incoming amount of money.
	RelatedAccount *CashAccount16 `xml:"RltdAcct,omitempty"`

	// Sum of the amounts in all the Item entries.
	TotalAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	ExpectedValueDate *ISODate `xml:"XpctdValDt,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *Party12Choice `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	IntermediaryAgent *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt,omitempty"`

	// Provides details of the expected amount on the account serviced by the account servicer.
	Item []*NotificationItem3 `xml:"Itm"`
}

func (a *AccountNotification4) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountNotification4) AddAccount() *CashAccount16 {
	a.Account = new(CashAccount16)
	return a.Account
}

func (a *AccountNotification4) AddAccountOwner() *Party12Choice {
	a.AccountOwner = new(Party12Choice)
	return a.AccountOwner
}

func (a *AccountNotification4) AddAccountServicer() *BranchAndFinancialInstitutionIdentification5 {
	a.AccountServicer = new(BranchAndFinancialInstitutionIdentification5)
	return a.AccountServicer
}

func (a *AccountNotification4) AddRelatedAccount() *CashAccount16 {
	a.RelatedAccount = new(CashAccount16)
	return a.RelatedAccount
}

func (a *AccountNotification4) SetTotalAmount(value, currency string) {
	a.TotalAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AccountNotification4) SetExpectedValueDate(value string) {
	a.ExpectedValueDate = (*ISODate)(&value)
}

func (a *AccountNotification4) AddDebtor() *Party12Choice {
	a.Debtor = new(Party12Choice)
	return a.Debtor
}

func (a *AccountNotification4) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.DebtorAgent
}

func (a *AccountNotification4) AddIntermediaryAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.IntermediaryAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.IntermediaryAgent
}

func (a *AccountNotification4) AddItem() *NotificationItem3 {
	newValue := new(NotificationItem3)
	a.Item = append(a.Item, newValue)
	return newValue
}
