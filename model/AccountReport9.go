package model

// Set of elements providing further details on the account report.
type AccountReport9 struct {

	// Unique and unambiguous identification of the account report, assigned by the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Sequential number of the report, assigned by the account servicer. It is increased incrementally for each report sent electronically.
	ElectronicSequenceNumber *Number `xml:"ElctrncSeqNb,omitempty"`

	// Legal sequential number of the report, assigned by the account servicer. It is increased incrementally for each report sent.
	//
	// Usage : in those scenarios where eg a paper statement is a legal requirement, the paper statement may have a different numbering than the electronic sequential number. Paper statements can for instance only be sent if movement on the account has taken place, whereas electronic statements can be sent eg each day, regardless of whether movements have taken place or not.
	LegalSequenceNumber *Number `xml:"LglSeqNb,omitempty"`

	// Date and time at which the report was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Range of time between the start date and the end date for which the account report is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Specifies if this document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *CashAccount13 `xml:"Acct"`

	// Identifies the parent account of the reported account.
	RelatedAccount *CashAccount7 `xml:"RltdAcct,omitempty"`

	// Provides general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest1 `xml:"Intrst,omitempty"`

	// Set of elements defining the balance(s).
	Balance []*CashBalance1 `xml:"Bal,omitempty"`

	// Set of element providing summary information on entries.
	TransactionsSummary *TotalTransactions1 `xml:"TxsSummry,omitempty"`

	// Specifies the elements of an entry in the report.
	//
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	Entry []*ReportEntry1 `xml:"Ntry,omitempty"`

	// Further details on the account report.
	AdditionalReportInformation *Max500Text `xml:"AddtlRptInf,omitempty"`
}

func (a *AccountReport9) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountReport9) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountReport9) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountReport9) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountReport9) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountReport9) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountReport9) AddAccount() *CashAccount13 {
	a.Account = new(CashAccount13)
	return a.Account
}

func (a *AccountReport9) AddRelatedAccount() *CashAccount7 {
	a.RelatedAccount = new(CashAccount7)
	return a.RelatedAccount
}

func (a *AccountReport9) AddInterest() *AccountInterest1 {
	newValue := new(AccountInterest1)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountReport9) AddBalance() *CashBalance1 {
	newValue := new(CashBalance1)
	a.Balance = append(a.Balance, newValue)
	return newValue
}

func (a *AccountReport9) AddTransactionsSummary() *TotalTransactions1 {
	a.TransactionsSummary = new(TotalTransactions1)
	return a.TransactionsSummary
}

func (a *AccountReport9) AddEntry() *ReportEntry1 {
	newValue := new(ReportEntry1)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountReport9) SetAdditionalReportInformation(value string) {
	a.AdditionalReportInformation = (*Max500Text)(&value)
}
