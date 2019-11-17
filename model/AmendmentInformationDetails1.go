package model

// Amendment information details providing the list of direct debit mandate elements that have been modified when the amendment indicator has been set.
type AmendmentInformationDetails1 struct {

	// Original mandate identification that has been modified.
	OriginalMandateIdentification *Max35Text `xml:"OrgnlMndtId,omitempty"`

	// Original creditor scheme identification that has been modified.
	OriginalCreditorSchemeIdentification *PartyIdentification8 `xml:"OrgnlCdtrSchmeId,omitempty"`

	// Original creditor agent that has been modified.
	OriginalCreditorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"OrgnlCdtrAgt,omitempty"`

	// Original creditor agent acount that has been modified.
	OriginalCreditorAgentAccount *CashAccount7 `xml:"OrgnlCdtrAgtAcct,omitempty"`

	// Original debtor that has been modified.
	OriginalDebtor *PartyIdentification8 `xml:"OrgnlDbtr,omitempty"`

	// Original debtor account that has been modified.
	OriginalDebtorAccount *CashAccount7 `xml:"OrgnlDbtrAcct,omitempty"`

	// Original debtor's agent that has been modified.
	OriginalDebtorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"OrgnlDbtrAgt,omitempty"`

	// Original debtor agent account that has been modified.
	OriginalDebtorAgentAccount *CashAccount7 `xml:"OrgnlDbtrAgtAcct,omitempty"`

	// Original final collection date that has been modified.
	OriginalFinalCollectionDate *ISODate `xml:"OrgnlFnlColltnDt,omitempty"`

	// Original frequency that has been modified.
	OriginalFrequency *Frequency1Code `xml:"OrgnlFrqcy,omitempty"`
}

func (a *AmendmentInformationDetails1) SetOriginalMandateIdentification(value string) {
	a.OriginalMandateIdentification = (*Max35Text)(&value)
}

func (a *AmendmentInformationDetails1) AddOriginalCreditorSchemeIdentification() *PartyIdentification8 {
	a.OriginalCreditorSchemeIdentification = new(PartyIdentification8)
	return a.OriginalCreditorSchemeIdentification
}

func (a *AmendmentInformationDetails1) AddOriginalCreditorAgent() *BranchAndFinancialInstitutionIdentification3 {
	a.OriginalCreditorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return a.OriginalCreditorAgent
}

func (a *AmendmentInformationDetails1) AddOriginalCreditorAgentAccount() *CashAccount7 {
	a.OriginalCreditorAgentAccount = new(CashAccount7)
	return a.OriginalCreditorAgentAccount
}

func (a *AmendmentInformationDetails1) AddOriginalDebtor() *PartyIdentification8 {
	a.OriginalDebtor = new(PartyIdentification8)
	return a.OriginalDebtor
}

func (a *AmendmentInformationDetails1) AddOriginalDebtorAccount() *CashAccount7 {
	a.OriginalDebtorAccount = new(CashAccount7)
	return a.OriginalDebtorAccount
}

func (a *AmendmentInformationDetails1) AddOriginalDebtorAgent() *BranchAndFinancialInstitutionIdentification3 {
	a.OriginalDebtorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return a.OriginalDebtorAgent
}

func (a *AmendmentInformationDetails1) AddOriginalDebtorAgentAccount() *CashAccount7 {
	a.OriginalDebtorAgentAccount = new(CashAccount7)
	return a.OriginalDebtorAgentAccount
}

func (a *AmendmentInformationDetails1) SetOriginalFinalCollectionDate(value string) {
	a.OriginalFinalCollectionDate = (*ISODate)(&value)
}

func (a *AmendmentInformationDetails1) SetOriginalFrequency(value string) {
	a.OriginalFrequency = (*Frequency1Code)(&value)
}
