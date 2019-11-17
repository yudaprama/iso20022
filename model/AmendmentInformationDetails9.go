package model

// Provides further details on the list of direct debit mandate elements that have been modified when the amendment indicator has been set.
type AmendmentInformationDetails9 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the original mandate.
	OriginalMandateIdentification *Max35Text `xml:"OrgnlMndtId,omitempty"`

	// Original creditor scheme identification that has been modified.
	OriginalCreditorSchemeIdentification *PartyIdentification43 `xml:"OrgnlCdtrSchmeId,omitempty"`

	// Original creditor agent that has been modified.
	OriginalCreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlCdtrAgt,omitempty"`

	// Original debtor that has been modified.
	OriginalDebtor *PartyIdentification43 `xml:"OrgnlDbtr,omitempty"`

	// Original debtor account that has been modified.
	OriginalDebtorAccount *CashAccount24 `xml:"OrgnlDbtrAcct,omitempty"`

	// Original debtor agent that has been modified.
	OriginalDebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlDbtrAgt,omitempty"`

	// Original final collection date that has been modified.
	OriginalFinalCollectionDate *ISODate `xml:"OrgnlFnlColltnDt,omitempty"`

	// Original frequency that has been modified.
	OriginalFrequency *Frequency6Code `xml:"OrgnlFrqcy,omitempty"`
}

func (a *AmendmentInformationDetails9) SetOriginalMandateIdentification(value string) {
	a.OriginalMandateIdentification = (*Max35Text)(&value)
}

func (a *AmendmentInformationDetails9) AddOriginalCreditorSchemeIdentification() *PartyIdentification43 {
	a.OriginalCreditorSchemeIdentification = new(PartyIdentification43)
	return a.OriginalCreditorSchemeIdentification
}

func (a *AmendmentInformationDetails9) AddOriginalCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.OriginalCreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.OriginalCreditorAgent
}

func (a *AmendmentInformationDetails9) AddOriginalDebtor() *PartyIdentification43 {
	a.OriginalDebtor = new(PartyIdentification43)
	return a.OriginalDebtor
}

func (a *AmendmentInformationDetails9) AddOriginalDebtorAccount() *CashAccount24 {
	a.OriginalDebtorAccount = new(CashAccount24)
	return a.OriginalDebtorAccount
}

func (a *AmendmentInformationDetails9) AddOriginalDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.OriginalDebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.OriginalDebtorAgent
}

func (a *AmendmentInformationDetails9) SetOriginalFinalCollectionDate(value string) {
	a.OriginalFinalCollectionDate = (*ISODate)(&value)
}

func (a *AmendmentInformationDetails9) SetOriginalFrequency(value string) {
	a.OriginalFrequency = (*Frequency6Code)(&value)
}
