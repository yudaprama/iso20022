package model

// Set of elements used to provide the list of direct debit mandate elements that have been modified when the amendment indicator has been set.
type AmendmentInformationDetails6 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the original mandate.
	OriginalMandateIdentification *Max35Text `xml:"OrgnlMndtId,omitempty"`

	// Original creditor scheme identification that has been modified.
	OriginalCreditorSchemeIdentification *PartyIdentification32 `xml:"OrgnlCdtrSchmeId,omitempty"`

	// Original creditor agent that has been modified.
	OriginalCreditorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"OrgnlCdtrAgt,omitempty"`

	// Original creditor agent acount that has been modified.
	OriginalCreditorAgentAccount *CashAccount16 `xml:"OrgnlCdtrAgtAcct,omitempty"`

	// Original debtor that has been modified.
	OriginalDebtor *PartyIdentification32 `xml:"OrgnlDbtr,omitempty"`

	// Original debtor account that has been modified.
	OriginalDebtorAccount *CashAccount16 `xml:"OrgnlDbtrAcct,omitempty"`

	// Original debtor agent that has been modified.
	OriginalDebtorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"OrgnlDbtrAgt,omitempty"`

	// Original debtor agent account that has been modified.
	OriginalDebtorAgentAccount *CashAccount16 `xml:"OrgnlDbtrAgtAcct,omitempty"`

	// Original final collection date that has been modified.
	OriginalFinalCollectionDate *ISODate `xml:"OrgnlFnlColltnDt,omitempty"`

	// Original frequency that has been modified.
	OriginalFrequency *Frequency1Code `xml:"OrgnlFrqcy,omitempty"`
}

func (a *AmendmentInformationDetails6) SetOriginalMandateIdentification(value string) {
	a.OriginalMandateIdentification = (*Max35Text)(&value)
}

func (a *AmendmentInformationDetails6) AddOriginalCreditorSchemeIdentification() *PartyIdentification32 {
	a.OriginalCreditorSchemeIdentification = new(PartyIdentification32)
	return a.OriginalCreditorSchemeIdentification
}

func (a *AmendmentInformationDetails6) AddOriginalCreditorAgent() *BranchAndFinancialInstitutionIdentification4 {
	a.OriginalCreditorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return a.OriginalCreditorAgent
}

func (a *AmendmentInformationDetails6) AddOriginalCreditorAgentAccount() *CashAccount16 {
	a.OriginalCreditorAgentAccount = new(CashAccount16)
	return a.OriginalCreditorAgentAccount
}

func (a *AmendmentInformationDetails6) AddOriginalDebtor() *PartyIdentification32 {
	a.OriginalDebtor = new(PartyIdentification32)
	return a.OriginalDebtor
}

func (a *AmendmentInformationDetails6) AddOriginalDebtorAccount() *CashAccount16 {
	a.OriginalDebtorAccount = new(CashAccount16)
	return a.OriginalDebtorAccount
}

func (a *AmendmentInformationDetails6) AddOriginalDebtorAgent() *BranchAndFinancialInstitutionIdentification4 {
	a.OriginalDebtorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return a.OriginalDebtorAgent
}

func (a *AmendmentInformationDetails6) AddOriginalDebtorAgentAccount() *CashAccount16 {
	a.OriginalDebtorAgentAccount = new(CashAccount16)
	return a.OriginalDebtorAgentAccount
}

func (a *AmendmentInformationDetails6) SetOriginalFinalCollectionDate(value string) {
	a.OriginalFinalCollectionDate = (*ISODate)(&value)
}

func (a *AmendmentInformationDetails6) SetOriginalFrequency(value string) {
	a.OriginalFrequency = (*Frequency1Code)(&value)
}
