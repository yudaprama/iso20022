package model

// Provides further details on the reporting request.
type ReportingRequest3 struct {

	// Unique identification, as assigned by the account owner, to unambiguously identify the account reporting request.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Specifies the type of the requested reporting message.
	RequestedMessageNameIdentification *Max35Text `xml:"ReqdMsgNmId"`

	// Unambiguous identification of the account to which the reporting request refers.
	Account *CashAccount24 `xml:"Acct,omitempty"`

	// Party that legally owns the account.
	AccountOwner *Party12Choice `xml:"AcctOwnr"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *BranchAndFinancialInstitutionIdentification5 `xml:"AcctSvcr,omitempty"`

	// Specifies the requested reporting period.
	ReportingPeriod *ReportingPeriod1 `xml:"RptgPrd,omitempty"`

	// Identifies the transactions to be reported.
	RequestedTransactionType *TransactionType1 `xml:"ReqdTxTp,omitempty"`

	// Provides details on the requested balance reporting.
	RequestedBalanceType []*BalanceType12 `xml:"ReqdBalTp,omitempty"`
}

func (r *ReportingRequest3) SetIdentification(value string) {
	r.Identification = (*Max35Text)(&value)
}

func (r *ReportingRequest3) SetRequestedMessageNameIdentification(value string) {
	r.RequestedMessageNameIdentification = (*Max35Text)(&value)
}

func (r *ReportingRequest3) AddAccount() *CashAccount24 {
	r.Account = new(CashAccount24)
	return r.Account
}

func (r *ReportingRequest3) AddAccountOwner() *Party12Choice {
	r.AccountOwner = new(Party12Choice)
	return r.AccountOwner
}

func (r *ReportingRequest3) AddAccountServicer() *BranchAndFinancialInstitutionIdentification5 {
	r.AccountServicer = new(BranchAndFinancialInstitutionIdentification5)
	return r.AccountServicer
}

func (r *ReportingRequest3) AddReportingPeriod() *ReportingPeriod1 {
	r.ReportingPeriod = new(ReportingPeriod1)
	return r.ReportingPeriod
}

func (r *ReportingRequest3) AddRequestedTransactionType() *TransactionType1 {
	r.RequestedTransactionType = new(TransactionType1)
	return r.RequestedTransactionType
}

func (r *ReportingRequest3) AddRequestedBalanceType() *BalanceType12 {
	newValue := new(BalanceType12)
	r.RequestedBalanceType = append(r.RequestedBalanceType, newValue)
	return newValue
}
