package model

// Provides further means of referencing a payment transaction.
type OriginalItemReference4 struct {

	// Identifies the account to be credited with the incoming amount of money.
	Account *CashAccount24 `xml:"Acct,omitempty"`

	// Party that legally owns the account.
	AccountOwner *Party12Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcr,omitempty"`

	// Identifies the parent account of the account to be credited with the incoming amount of money.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *Party12Choice `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	IntermediaryAgent *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt,omitempty"`

	// Underlying reason for the payment transaction.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Provides information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation *RemittanceLocation4 `xml:"RltdRmtInf,omitempty"`

	// Structured information that enables the matching, that is reconciliation, of a payment with the items that the payment is intended to settle, such as commercial invoices in an account receivable system.
	RemittanceInformation *RemittanceInformation11 `xml:"RmtInf,omitempty"`
}

func (o *OriginalItemReference4) AddAccount() *CashAccount24 {
	o.Account = new(CashAccount24)
	return o.Account
}

func (o *OriginalItemReference4) AddAccountOwner() *Party12Choice {
	o.AccountOwner = new(Party12Choice)
	return o.AccountOwner
}

func (o *OriginalItemReference4) AddAccountServicer() *BranchAndFinancialInstitutionIdentification5 {
	o.AccountServicer = new(BranchAndFinancialInstitutionIdentification5)
	return o.AccountServicer
}

func (o *OriginalItemReference4) AddRelatedAccount() *CashAccount24 {
	o.RelatedAccount = new(CashAccount24)
	return o.RelatedAccount
}

func (o *OriginalItemReference4) AddDebtor() *Party12Choice {
	o.Debtor = new(Party12Choice)
	return o.Debtor
}

func (o *OriginalItemReference4) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.DebtorAgent
}

func (o *OriginalItemReference4) AddIntermediaryAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.IntermediaryAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.IntermediaryAgent
}

func (o *OriginalItemReference4) AddPurpose() *Purpose2Choice {
	o.Purpose = new(Purpose2Choice)
	return o.Purpose
}

func (o *OriginalItemReference4) AddRelatedRemittanceInformation() *RemittanceLocation4 {
	o.RelatedRemittanceInformation = new(RemittanceLocation4)
	return o.RelatedRemittanceInformation
}

func (o *OriginalItemReference4) AddRemittanceInformation() *RemittanceInformation11 {
	o.RemittanceInformation = new(RemittanceInformation11)
	return o.RemittanceInformation
}
