package model

// Provides further means of referencing a payment transaction.
type NotificationItem6 struct {

	// Unique identification, as assigned by the account owner, to unambiguously identify the expected credit entry.
	Identification *Max35Text `xml:"Id"`

	// Unique identification, as assigned by the debtor, to unambiguously identify the underlying transaction to the creditor.
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction.
	EndToEndIdentification *Max35Text `xml:"EndToEndId,omitempty"`

	// Identifies the account to be credited with the incoming amount of money.
	Account *CashAccount24 `xml:"Acct,omitempty"`

	// Party that legally owns the account.
	AccountOwner *Party12Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcr,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Amount of money expected to be credited to the account.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Value date on which the account is expected to be credited.
	ExpectedValueDate *ISODate `xml:"XpctdValDt,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *Party12Choice `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Agent between the debtor agent and the account servicer.
	// Usage: This is the agent from which the account servicer will get the amount of money. If there is more than one intermediary agent, then IntermediaryAgent identifies the agent closest to the account servicer.
	// IntermediaryAgent must only be included when different from the debtor agent.
	IntermediaryAgent *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt,omitempty"`

	// Underlying reason for the payment transaction.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Provides information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation *RemittanceLocation4 `xml:"RltdRmtInf,omitempty"`

	// Structured information that enables the reconciliation of a payment with the items that the payment is intended to settle, such as commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation11 `xml:"RmtInf,omitempty"`
}

func (n *NotificationItem6) SetIdentification(value string) {
	n.Identification = (*Max35Text)(&value)
}

func (n *NotificationItem6) SetEndToEndIdentification(value string) {
	n.EndToEndIdentification = (*Max35Text)(&value)
}

func (n *NotificationItem6) AddAccount() *CashAccount24 {
	n.Account = new(CashAccount24)
	return n.Account
}

func (n *NotificationItem6) AddAccountOwner() *Party12Choice {
	n.AccountOwner = new(Party12Choice)
	return n.AccountOwner
}

func (n *NotificationItem6) AddAccountServicer() *BranchAndFinancialInstitutionIdentification5 {
	n.AccountServicer = new(BranchAndFinancialInstitutionIdentification5)
	return n.AccountServicer
}

func (n *NotificationItem6) AddRelatedAccount() *CashAccount24 {
	n.RelatedAccount = new(CashAccount24)
	return n.RelatedAccount
}

func (n *NotificationItem6) SetAmount(value, currency string) {
	n.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (n *NotificationItem6) SetExpectedValueDate(value string) {
	n.ExpectedValueDate = (*ISODate)(&value)
}

func (n *NotificationItem6) AddDebtor() *Party12Choice {
	n.Debtor = new(Party12Choice)
	return n.Debtor
}

func (n *NotificationItem6) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	n.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return n.DebtorAgent
}

func (n *NotificationItem6) AddIntermediaryAgent() *BranchAndFinancialInstitutionIdentification5 {
	n.IntermediaryAgent = new(BranchAndFinancialInstitutionIdentification5)
	return n.IntermediaryAgent
}

func (n *NotificationItem6) AddPurpose() *Purpose2Choice {
	n.Purpose = new(Purpose2Choice)
	return n.Purpose
}

func (n *NotificationItem6) AddRelatedRemittanceInformation() *RemittanceLocation4 {
	n.RelatedRemittanceInformation = new(RemittanceLocation4)
	return n.RelatedRemittanceInformation
}

func (n *NotificationItem6) AddRemittanceInformation() *RemittanceInformation11 {
	n.RemittanceInformation = new(RemittanceInformation11)
	return n.RemittanceInformation
}
