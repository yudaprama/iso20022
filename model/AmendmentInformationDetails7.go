package model

// Set of elements used to provide the list of direct debit mandate elements that have been modified when the amendment indicator has been set.
type AmendmentInformationDetails7 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the original mandate.
	OriginalMandateIdentification *Max35Text `xml:"OrgnlMndtId,omitempty"`

	// Original creditor scheme identification that has been modified.
	OriginalCreditorSchemeIdentification *PartyIdentification43 `xml:"OrgnlCdtrSchmeId,omitempty"`

	// Original creditor agent that has been modified.
	OriginalCreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlCdtrAgt,omitempty"`

	// Original debtor that has been modified.
	OriginalDebtor *PartyIdentification43 `xml:"OrgnlDbtr,omitempty"`

	// Original debtor account that has been modified.
	OriginalDebtorAccount *CashAccount16 `xml:"OrgnlDbtrAcct,omitempty"`

	// Original debtor agent that has been modified.
	OriginalDebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlDbtrAgt,omitempty"`

	// Original final collection date that has been modified.
	OriginalFinalCollectionDate *ISODate `xml:"OrgnlFnlColltnDt,omitempty"`

	// Original frequency that has been modified.
	OriginalFrequency *Frequency1Code `xml:"OrgnlFrqcy,omitempty"`
}

func (a *AmendmentInformationDetails7) SetOriginalMandateIdentification(value string) {
	a.OriginalMandateIdentification = (*Max35Text)(&value)
}

func (a *AmendmentInformationDetails7) AddOriginalCreditorSchemeIdentification() *PartyIdentification43 {
	a.OriginalCreditorSchemeIdentification = new(PartyIdentification43)
	return a.OriginalCreditorSchemeIdentification
}

func (a *AmendmentInformationDetails7) AddOriginalCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.OriginalCreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.OriginalCreditorAgent
}

func (a *AmendmentInformationDetails7) AddOriginalDebtor() *PartyIdentification43 {
	a.OriginalDebtor = new(PartyIdentification43)
	return a.OriginalDebtor
}

func (a *AmendmentInformationDetails7) AddOriginalDebtorAccount() *CashAccount16 {
	a.OriginalDebtorAccount = new(CashAccount16)
	return a.OriginalDebtorAccount
}

func (a *AmendmentInformationDetails7) AddOriginalDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.OriginalDebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.OriginalDebtorAgent
}

func (a *AmendmentInformationDetails7) SetOriginalFinalCollectionDate(value string) {
	a.OriginalFinalCollectionDate = (*ISODate)(&value)
}

func (a *AmendmentInformationDetails7) SetOriginalFrequency(value string) {
	a.OriginalFrequency = (*Frequency1Code)(&value)
}
