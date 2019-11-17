package model

// Provides further details on the list of direct debit mandate elements that have been modified when the amendment indicator has been set.
type AmendmentInformationDetails11 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the original mandate.
	OriginalMandateIdentification *Max35Text `xml:"OrgnlMndtId,omitempty"`

	// Original creditor scheme identification that has been modified.
	OriginalCreditorSchemeIdentification *PartyIdentification43 `xml:"OrgnlCdtrSchmeId,omitempty"`

	// Original creditor agent that has been modified.
	OriginalCreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlCdtrAgt,omitempty"`

	// Original creditor agent account that has been modified.
	OriginalCreditorAgentAccount *CashAccount24 `xml:"OrgnlCdtrAgtAcct,omitempty"`

	// Original debtor that has been modified.
	OriginalDebtor *PartyIdentification43 `xml:"OrgnlDbtr,omitempty"`

	// Original debtor account that has been modified.
	OriginalDebtorAccount *CashAccount24 `xml:"OrgnlDbtrAcct,omitempty"`

	// Original debtor agent that has been modified.
	OriginalDebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlDbtrAgt,omitempty"`

	// Original debtor agent account that has been modified.
	OriginalDebtorAgentAccount *CashAccount24 `xml:"OrgnlDbtrAgtAcct,omitempty"`

	// Original final collection date that has been modified.
	OriginalFinalCollectionDate *ISODate `xml:"OrgnlFnlColltnDt,omitempty"`

	// Original frequency that has been modified.
	OriginalFrequency *Frequency36Choice `xml:"OrgnlFrqcy,omitempty"`

	// Original reason for the mandate to allow the user to distinguish between different mandates for the same creditor.
	OriginalReason *MandateSetupReason1Choice `xml:"OrgnlRsn,omitempty"`

	// Original number of tracking days that has been modified.
	OriginalTrackingDays *Exact2NumericText `xml:"OrgnlTrckgDays,omitempty"`
}

func (a *AmendmentInformationDetails11) SetOriginalMandateIdentification(value string) {
	a.OriginalMandateIdentification = (*Max35Text)(&value)
}

func (a *AmendmentInformationDetails11) AddOriginalCreditorSchemeIdentification() *PartyIdentification43 {
	a.OriginalCreditorSchemeIdentification = new(PartyIdentification43)
	return a.OriginalCreditorSchemeIdentification
}

func (a *AmendmentInformationDetails11) AddOriginalCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.OriginalCreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.OriginalCreditorAgent
}

func (a *AmendmentInformationDetails11) AddOriginalCreditorAgentAccount() *CashAccount24 {
	a.OriginalCreditorAgentAccount = new(CashAccount24)
	return a.OriginalCreditorAgentAccount
}

func (a *AmendmentInformationDetails11) AddOriginalDebtor() *PartyIdentification43 {
	a.OriginalDebtor = new(PartyIdentification43)
	return a.OriginalDebtor
}

func (a *AmendmentInformationDetails11) AddOriginalDebtorAccount() *CashAccount24 {
	a.OriginalDebtorAccount = new(CashAccount24)
	return a.OriginalDebtorAccount
}

func (a *AmendmentInformationDetails11) AddOriginalDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	a.OriginalDebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return a.OriginalDebtorAgent
}

func (a *AmendmentInformationDetails11) AddOriginalDebtorAgentAccount() *CashAccount24 {
	a.OriginalDebtorAgentAccount = new(CashAccount24)
	return a.OriginalDebtorAgentAccount
}

func (a *AmendmentInformationDetails11) SetOriginalFinalCollectionDate(value string) {
	a.OriginalFinalCollectionDate = (*ISODate)(&value)
}

func (a *AmendmentInformationDetails11) AddOriginalFrequency() *Frequency36Choice {
	a.OriginalFrequency = new(Frequency36Choice)
	return a.OriginalFrequency
}

func (a *AmendmentInformationDetails11) AddOriginalReason() *MandateSetupReason1Choice {
	a.OriginalReason = new(MandateSetupReason1Choice)
	return a.OriginalReason
}

func (a *AmendmentInformationDetails11) SetOriginalTrackingDays(value string) {
	a.OriginalTrackingDays = (*Exact2NumericText)(&value)
}
