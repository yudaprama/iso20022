package model

// Provides further details of the account notification.
type AccountNotification7 struct {

	// Unique identification, as assigned by the account servicer, to unambiguously identify the account notification.
	Identification *Max35Text `xml:"Id"`

	// Provides details on the page number of the notification.
	//
	// Usage: The pagination of the notification is only allowed when agreed between the parties.
	NotificationPagination *Pagination `xml:"NtfctnPgntn,omitempty"`

	// Sequential number of the notification, as assigned by the account servicer.
	// Usage: The sequential number is increased incrementally for each notification sent electronically.
	ElectronicSequenceNumber *Number `xml:"ElctrncSeqNb,omitempty"`

	// Legal sequential number of the notification, as assigned by the account servicer. It is increased incrementally for each notification sent.
	LegalSequenceNumber *Number `xml:"LglSeqNb,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Range of time between a start date and an end date for which the account notification is issued.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Indicates whether the document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicateIndicator *CopyDuplicate1Code `xml:"CpyDplctInd,omitempty"`

	// Specifies the application used to generate the reporting.
	ReportingSource *ReportingSource1Choice `xml:"RptgSrc,omitempty"`

	// Unambiguous identification of the account to which credit and debit entries are made.
	Account *CashAccount25 `xml:"Acct"`

	// Identifies the parent account of the account for which the notification has been issued.
	RelatedAccount *CashAccount24 `xml:"RltdAcct,omitempty"`

	// Provides general interest information that applies to the account at a particular moment in time.
	Interest []*AccountInterest2 `xml:"Intrst,omitempty"`

	// Provides summary information on entries.
	TransactionsSummary *TotalTransactions4 `xml:"TxsSummry,omitempty"`

	// Set of elements used to specify an entry in the debit credit notification.
	// Usage: At least one reference must be provided to identify the entry and its underlying transaction(s).
	Entry []*ReportEntry4 `xml:"Ntry,omitempty"`

	// Further details of the account notification.
	AdditionalNotificationInformation *Max500Text `xml:"AddtlNtfctnInf,omitempty"`
}

func (a *AccountNotification7) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AccountNotification7) AddNotificationPagination() *Pagination {
	a.NotificationPagination = new(Pagination)
	return a.NotificationPagination
}

func (a *AccountNotification7) SetElectronicSequenceNumber(value string) {
	a.ElectronicSequenceNumber = (*Number)(&value)
}

func (a *AccountNotification7) SetLegalSequenceNumber(value string) {
	a.LegalSequenceNumber = (*Number)(&value)
}

func (a *AccountNotification7) SetCreationDateTime(value string) {
	a.CreationDateTime = (*ISODateTime)(&value)
}

func (a *AccountNotification7) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountNotification7) SetCopyDuplicateIndicator(value string) {
	a.CopyDuplicateIndicator = (*CopyDuplicate1Code)(&value)
}

func (a *AccountNotification7) AddReportingSource() *ReportingSource1Choice {
	a.ReportingSource = new(ReportingSource1Choice)
	return a.ReportingSource
}

func (a *AccountNotification7) AddAccount() *CashAccount25 {
	a.Account = new(CashAccount25)
	return a.Account
}

func (a *AccountNotification7) AddRelatedAccount() *CashAccount24 {
	a.RelatedAccount = new(CashAccount24)
	return a.RelatedAccount
}

func (a *AccountNotification7) AddInterest() *AccountInterest2 {
	newValue := new(AccountInterest2)
	a.Interest = append(a.Interest, newValue)
	return newValue
}

func (a *AccountNotification7) AddTransactionsSummary() *TotalTransactions4 {
	a.TransactionsSummary = new(TotalTransactions4)
	return a.TransactionsSummary
}

func (a *AccountNotification7) AddEntry() *ReportEntry4 {
	newValue := new(ReportEntry4)
	a.Entry = append(a.Entry, newValue)
	return newValue
}

func (a *AccountNotification7) SetAdditionalNotificationInformation(value string) {
	a.AdditionalNotificationInformation = (*Max500Text)(&value)
}
