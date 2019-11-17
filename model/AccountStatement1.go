package model

// Set of elements providing further details on the account statement.
type AccountStatement1 struct {

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

	// Range of time between the start date and the end date for which the account statement is issued.
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
	Balance []*CashBalance2 `xml:"Bal"`

	// Set of element providing summary information on entries.
	TransactionsSummary *TotalTransactions1 `xml:"TxsSummry,omitempty"`

	// Specifies the elements of an entry in the statement.
	//
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	Entry []*StatementEntry1 `xml:"Ntry,omitempty"`

	// Further details on the account statement.
	AdditionalStatementInformation *Max500Text `xml:"AddtlStmtInf,omitempty"`
}

func (a *AccountStatement1) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountStatement1) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountStatement1) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountStatement1) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountStatement1) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountStatement1) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountStatement1) AddAccount() *CashAccount13 {
	a.Account = new(CashAccount13)
	return a.Account
}

func (a *AccountStatement1) AddRelatedAccount() *CashAccount7 {
	a.RelatedAccount = new(CashAccount7)
	return a.RelatedAccount
}

func (a *AccountStatement1) AddInterest() *AccountInterest1 {
	newValue := new(AccountInterest1)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountStatement1) AddBalance() *CashBalance2 {
	newValue := new(CashBalance2)
	a.Balance = append(a.Balance, newValue)
	return newValue
}

func (a *AccountStatement1) AddTransactionsSummary() *TotalTransactions1 {
	a.TransactionsSummary = new(TotalTransactions1)
	return a.TransactionsSummary
}

func (a *AccountStatement1) AddEntry() *StatementEntry1 {
	newValue := new(StatementEntry1)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountStatement1) SetAdditionalStatementInformation(value string) {
	a.AdditionalStatementInformation = (*Max500Text)(&value)
}
