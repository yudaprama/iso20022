package model

// Provides further details of the account report.
type AccountReport12 struct {

	// Unique identification, as assigned by the account servicer, to unambiguously identify the account report.
	Identification *Max35Text `xml:"Id"`

	// Provides details on the page number of the report.
	//
	// Usage: The pagination of the report is only allowed when agreed between the parties.
	ReportPagination *Pagination `xml:"RptPgntn,omitempty"`

	// Sequential number of the report, as assigned by the account servicer.
	// Usage: The sequential number is increased incrementally for each report sent electronically.
	ElectronicSequenceNumber *Number `xml:"ElctrncSeqNb,omitempty"`

	// Legal sequential number of the report, as assigned by the account servicer. It is increased incrementally for each report sent.
	LegalSequenceNumber *Number `xml:"LglSeqNb,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Range of time between a start date and an end date for which the account report is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Indicates whether the document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Specifies the application used to generate the reporting.
	ReportingSource *ReportingSource1Choice `xml:"RptgSrc,omitempty"`

	// Unambiguous identification of the account to which credit and debit entries are made.
	Account *CashAccount25 `xml:"Acct"`

	// Identifies the parent account of the account for which the report has been issued.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Provides general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest3 `xml:"Intrst,omitempty"`

	// Set of elements used to define the balance as a numerical representation of the net increases and decreases in an account at a specific point in time.
	Balance []*CashBalance3 `xml:"Bal,omitempty"`

	// Provides summary information on entries.
	TransactionsSummary *TotalTransactions3 `xml:"TxsSummry,omitempty"`

	// Set of elements used to specify an entry in the report.
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	Entry []*ReportEntry3 `xml:"Ntry,omitempty"`

	// Further details of the account report.
	AdditionalReportInformation *Max500Text `xml:"AddtlRptInf,omitempty"`
}

func (a *AccountReport12) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountReport12) AddReportPagination() *Pagination {
	a.ReportPagination = new(Pagination)
	return a.ReportPagination
}

func (a *AccountReport12) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountReport12) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountReport12) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountReport12) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountReport12) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountReport12) AddReportingSource() *ReportingSource1Choice {
	a.ReportingSource = new(ReportingSource1Choice)
	return a.ReportingSource
}

func (a *AccountReport12) AddAccount() *CashAccount25 {
	a.Account = new(CashAccount25)
	return a.Account
}

func (a *AccountReport12) AddRelatedAccount() *CashAccount24 {
	a.RelatedAccount = new(CashAccount24)
	return a.RelatedAccount
}

func (a *AccountReport12) AddInterest() *AccountInterest3 {
	newValue := new(AccountInterest3)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountReport12) AddBalance() *CashBalance3 {
	newValue := new(CashBalance3)
	a.Balance = append(a.Balance, newValue)
	return newValue
}

func (a *AccountReport12) AddTransactionsSummary() *TotalTransactions3 {
	a.TransactionsSummary = new(TotalTransactions3)
	return a.TransactionsSummary
}

func (a *AccountReport12) AddEntry() *ReportEntry3 {
	newValue := new(ReportEntry3)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountReport12) SetAdditionalReportInformation(value string) {
	a.AdditionalReportInformation = (*Max500Text)(&value)
}
