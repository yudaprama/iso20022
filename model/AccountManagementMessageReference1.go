package model

// Information about the message reference of the account management instruction message for which the status is requested .
type AccountManagementMessageReference1 struct {

	// Reference to a linked message sent in a proprietary way or reference of a system.
	OtherReference *AdditionalReference3 `xml:"OthrRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Specifies if the status request refers to an earlier account opening or modification instruction message.
	StatusRequestType *AccountManagementType1Code `xml:"StsReqTp"`

	// Unique and unambiguous identifier of the account opening or account modification instruction at application level.
	AccountApplicationIdentification *Max35Text `xml:"AcctApplId,omitempty"`

	// Account information for which the status of an account management instruction is requested.
	InvestmentAccount *InvestmentAccount14 `xml:"InvstmtAcct,omitempty"`
}

func (a *AccountManagementMessageReference1) AddOtherReference() *AdditionalReference3 {
	a.OtherReference = new(AdditionalReference3)
	return a.OtherReference
}

func (a *AccountManagementMessageReference1) AddPreviousReference() *AdditionalReference3 {
	a.PreviousReference = new(AdditionalReference3)
	return a.PreviousReference
}

func (a *AccountManagementMessageReference1) SetStatusRequestType(value string) {
	a.StatusRequestType = (*AccountManagementType1Code)(&value)
}

func (a *AccountManagementMessageReference1) SetAccountApplicationIdentification(value string) {
	a.AccountApplicationIdentification = (*Max35Text)(&value)
}

func (a *AccountManagementMessageReference1) AddInvestmentAccount() *InvestmentAccount14 {
	a.InvestmentAccount = new(InvestmentAccount14)
	return a.InvestmentAccount
}
