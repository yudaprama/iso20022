package model

// Set of elements used to provide details of the account report.
type AccountReport11 struct {

	// Unique identification, as assigned by the account servicer, to unambiguously identify the account report.
	Identification *Max35Text `xml:"Id"`

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
	Account *CashAccount20 `xml:"Acct"`

	// Identifies the parent account of the account for which the report has been issued.
	RelatedAccount *CashAccount16 `xml:"RltdAcct,omitempty"`

	// Set of elements used to provide general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest2 `xml:"Intrst,omitempty"`

	// Set of elements used to define the balance as a numerical representation of the net increases and decreases in an account at a specific point in time.
	Balance []*CashBalance3 `xml:"Bal,omitempty"`

	// Set of elements used to provide summary information on entries.
	TransactionsSummary *TotalTransactions2 `xml:"TxsSummry,omitempty"`

	// Set of elements used to specify an entry in the report.
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	Entry []*ReportEntry2 `xml:"Ntry,omitempty"`

	// Further details of the account report.
	AdditionalReportInformation *Max500Text `xml:"AddtlRptInf,omitempty"`
}

func (a *AccountReport11) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountReport11) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountReport11) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountReport11) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountReport11) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountReport11) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountReport11) AddReportingSource() *ReportingSource1Choice {
	a.ReportingSource = new(ReportingSource1Choice)
	return a.ReportingSource
}

func (a *AccountReport11) AddAccount() *CashAccount20 {
	a.Account = new(CashAccount20)
	return a.Account
}

func (a *AccountReport11) AddRelatedAccount() *CashAccount16 {
	a.RelatedAccount = new(CashAccount16)
	return a.RelatedAccount
}

func (a *AccountReport11) AddInterest() *AccountInterest2 {
	newValue := new(AccountInterest2)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountReport11) AddBalance() *CashBalance3 {
	newValue := new(CashBalance3)
	a.Balance = append(a.Balance, newValue)
	return newValue
}

func (a *AccountReport11) AddTransactionsSummary() *TotalTransactions2 {
	a.TransactionsSummary = new(TotalTransactions2)
	return a.TransactionsSummary
}

func (a *AccountReport11) AddEntry() *ReportEntry2 {
	newValue := new(ReportEntry2)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountReport11) SetAdditionalReportInformation(value string) {
	a.AdditionalReportInformation = (*Max500Text)(&value)
}
