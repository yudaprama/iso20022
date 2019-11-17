package model

// Status information.
type AccountManagementStatusAndReason5 struct {

	// Status of the account management instruction that was previously received.
	Status *Status25Choice `xml:"Sts"`

	// Reason for the status of the account management instruction.
	StatusReason []*AcceptedStatusReason1Choice `xml:"StsRsn,omitempty"`

	// Unique and unambiguous identifier of the account opening or modification instruction at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Account to which the account opening is related.
	ExistingAccountIdentification []*Account23 `xml:"ExstgAcctId,omitempty"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Status of the account.
	AccountStatus *AccountStatus2 `xml:"AcctSts,omitempty"`

	// Specifies the account is blocked and other factors for the blocked account.
	BlockedStatus *BlockedStatusReason2Choice `xml:"BlckdSts,omitempty"`

	// Date provided by the account owner to inform the account servicer of the date on which the holdings must be reported before the account is subsequently closed.
	FATCAReportingDate *ISODate `xml:"FATCARptgDt,omitempty"`

	// Date provided by the account owner to inform the account servicer of the date on which the holdings must be reported before the account is subsequently closed.
	CRSReportingDate *ISODate `xml:"CRSRptgDt,omitempty"`
}

func (a *AccountManagementStatusAndReason5) AddStatus() *Status25Choice {
	a.Status = new(Status25Choice)
	return a.Status
}

func (a *AccountManagementStatusAndReason5) AddStatusReason() *AcceptedStatusReason1Choice {
	newValue := new(AcceptedStatusReason1Choice)
	a.StatusReason = append(a.StatusReason, newValue)
	return newValue
}

func (a *AccountManagementStatusAndReason5) SetAccountApplicationIdentification(value string) {
	a.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (a *AccountManagementStatusAndReason5) AddExistingAccountIdentification() *Account23 {
	newValue := new(Account23)
	a.ExistingAccountIdentification = append(a.ExistingAccountIdentification, newValue)
	return newValue
}

func (a *AccountManagementStatusAndReason5) SetAccountIdentification(value string) {
	a.AccountIdentification = (*Max35Text)(&value)
}

func (a *AccountManagementStatusAndReason5) AddAccountStatus() *AccountStatus2 {
	a.AccountStatus = new(AccountStatus2)
	return a.AccountStatus
}

func (a *AccountManagementStatusAndReason5) AddBlockedStatus() *BlockedStatusReason2Choice {
	a.BlockedStatus = new(BlockedStatusReason2Choice)
	return a.BlockedStatus
}

func (a *AccountManagementStatusAndReason5) SetFATCAReportingDate(value string) {
	a.FATCAReportingDate = (*ISODate)(&value)
}

func (a *AccountManagementStatusAndReason5) SetCRSReportingDate(value string) {
	a.CRSReportingDate = (*ISODate)(&value)
}
