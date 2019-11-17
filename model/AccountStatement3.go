package model

// Provides further details of the account statement.
type AccountStatement3 struct {

	// Unique identification, as assigned by the account servicer, to unambiguously identify the account statement.
	Identification *Max35Text `xml:"Id"`

	// Provides details on the page number of the statement.
	//
	// Usage: The pagination of the statement is only allowed when agreed between the parties.
	StatementPagination *Pagination `xml:"StmtPgntn,omitempty"`

	// Sequential number of the statement, as assigned by the account servicer.
	// Usage: The sequential number is increased incrementally for each statement sent electronically.
	ElectronicSequenceNumber *Number `xml:"ElctrncSeqNb,omitempty"`

	// Legal sequential number of the statement, as assigned by the account servicer. It is increased incrementally for each statement sent.
	//
	// Usage: Where a paper statement is a legal requirement, it may have a number different from the electronic sequential number. Paper statements could for instance only be sent if movement on the account has taken place, whereas electronic statements could be sent at the end of each reporting period, regardless of whether movements have taken place or not.
	LegalSequenceNumber *Number `xml:"LglSeqNb,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Range of time between a start date and an end date for which the account statement is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Indicates whether the document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Specifies the application used to generate the reporting.
	ReportingSource *ReportingSource1Choice `xml:"RptgSrc,omitempty"`

	// Unambiguous identification of the account to which credit and debit entries are made.
	Account *CashAccount25 `xml:"Acct"`

	// Identifies the parent account of the account for which the statement has been issued.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Provides general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest2 `xml:"Intrst,omitempty"`

	// Set of elements used to define the balance as a numerical representation of the net increases and decreases in an account at a specific point in time.
	Balance []*CashBalance3 `xml:"Bal"`

	// Provides summary information on entries.
	TransactionsSummary *TotalTransactions2 `xml:"TxsSummry,omitempty"`

	// Set of elements used to specify an entry in the statement.
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	Entry []*ReportEntry3 `xml:"Ntry,omitempty"`

	// Further details of the account statement.
	AdditionalStatementInformation *Max500Text `xml:"AddtlStmtInf,omitempty"`
}

func (a *AccountStatement3) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountStatement3) AddStatementPagination() *Pagination {
	a.StatementPagination = new(Pagination)
	return a.StatementPagination
}

func (a *AccountStatement3) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountStatement3) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountStatement3) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountStatement3) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountStatement3) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountStatement3) AddReportingSource() *ReportingSource1Choice {
	a.ReportingSource = new(ReportingSource1Choice)
	return a.ReportingSource
}

func (a *AccountStatement3) AddAccount() *CashAccount25 {
	a.Account = new(CashAccount25)
	return a.Account
}

func (a *AccountStatement3) AddRelatedAccount() *CashAccount24 {
	a.RelatedAccount = new(CashAccount24)
	return a.RelatedAccount
}

func (a *AccountStatement3) AddInterest() *AccountInterest2 {
	newValue := new(AccountInterest2)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountStatement3) AddBalance() *CashBalance3 {
	newValue := new(CashBalance3)
	a.Balance = append(a.Balance, newValue)
	return newValue
}

func (a *AccountStatement3) AddTransactionsSummary() *TotalTransactions2 {
	a.TransactionsSummary = new(TotalTransactions2)
	return a.TransactionsSummary
}

func (a *AccountStatement3) AddEntry() *ReportEntry3 {
	newValue := new(ReportEntry3)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountStatement3) SetAdditionalStatementInformation(value string) {
	a.AdditionalStatementInformation = (*Max500Text)(&value)
}
