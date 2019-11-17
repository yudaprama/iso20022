package model

// Information that serves as a basis to debit an account.
type Mandate8 struct {

	// Unique identification, as assigned by the responsible party (such as the creditor) or agent (such as the debtor agent), to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId"`

	// Identification for the mandate request, as assigned by the initiating party.
	MandateRequestIdentification *Max35Text `xml:"MndtReqId,omitempty"`

	// Specifies the transport authentication details related to the mandate.
	Authentication *MandateAuthentication1 `xml:"Authntcn,omitempty"`

	// Specifies the type of mandate, such as paper, electronic or scheme.
	Type *MandateTypeInformation2 `xml:"Tp,omitempty"`

	// Provides details of the duration of the mandate and occurrence of the underlying transactions.
	Occurrences *MandateOccurrences4 `xml:"Ocrncs,omitempty"`

	// Specifies whether the direct debit instructions should be automatically re-submitted periodically when bilaterally agreed.
	TrackingIndicator *TrueFalseIndicator `xml:"TrckgInd"`

	// Amount different from the collection amount, as it includes the costs associated with the first debited amount.
	FirstCollectionAmount *ActiveCurrencyAndAmount `xml:"FrstColltnAmt,omitempty"`

	// Fixed amount to be collected from the debtor's account.
	CollectionAmount *ActiveCurrencyAndAmount `xml:"ColltnAmt,omitempty"`

	// Maximum amount that may be collected from the debtor's account, per instruction.
	MaximumAmount *ActiveCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Specifies the characteristics of the adjustment applied to the collection amount of a direct debit instruction.
	Adjustment *MandateAdjustment1 `xml:"Adjstmnt,omitempty"`

	// Provides the reason for the setup of the mandate.
	Reason *MandateSetupReason1Choice `xml:"Rsn,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Party that signs the mandate and to whom an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Party that signs the mandate and owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor, to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Reference assigned by a creditor or ultimate creditor for internal usage for the mandate.
	MandateReference *Max35Text `xml:"MndtRef,omitempty"`

	// Provides information to identify the underlying documents associated with the mandate.
	ReferredDocument []*ReferredMandateDocument1 `xml:"RfrdDoc,omitempty"`
}

func (m *Mandate8) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *Mandate8) SetMandateRequestIdentification(value string) {
	m.MandateRequestIdentification = (*Max35Text)(&value)
}

func (m *Mandate8) AddAuthentication() *MandateAuthentication1 {
	m.Authentication = new(MandateAuthentication1)
	return m.Authentication
}

func (m *Mandate8) AddType() *MandateTypeInformation2 {
	m.Type = new(MandateTypeInformation2)
	return m.Type
}

func (m *Mandate8) AddOccurrences() *MandateOccurrences4 {
	m.Occurrences = new(MandateOccurrences4)
	return m.Occurrences
}

func (m *Mandate8) SetTrackingIndicator(value string) {
	m.TrackingIndicator = (*TrueFalseIndicator)(&value)
}

func (m *Mandate8) SetFirstCollectionAmount(value, currency string) {
	m.FirstCollectionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *Mandate8) SetCollectionAmount(value, currency string) {
	m.CollectionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *Mandate8) SetMaximumAmount(value, currency string) {
	m.MaximumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *Mandate8) AddAdjustment() *MandateAdjustment1 {
	m.Adjustment = new(MandateAdjustment1)
	return m.Adjustment
}

func (m *Mandate8) AddReason() *MandateSetupReason1Choice {
	m.Reason = new(MandateSetupReason1Choice)
	return m.Reason
}

func (m *Mandate8) AddCreditorSchemeIdentification() *PartyIdentification43 {
	m.CreditorSchemeIdentification = new(PartyIdentification43)
	return m.CreditorSchemeIdentification
}

func (m *Mandate8) AddCreditor() *PartyIdentification43 {
	m.Creditor = new(PartyIdentification43)
	return m.Creditor
}

func (m *Mandate8) AddCreditorAccount() *CashAccount24 {
	m.CreditorAccount = new(CashAccount24)
	return m.CreditorAccount
}

func (m *Mandate8) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.CreditorAgent
}

func (m *Mandate8) AddUltimateCreditor() *PartyIdentification43 {
	m.UltimateCreditor = new(PartyIdentification43)
	return m.UltimateCreditor
}

func (m *Mandate8) AddDebtor() *PartyIdentification43 {
	m.Debtor = new(PartyIdentification43)
	return m.Debtor
}

func (m *Mandate8) AddDebtorAccount() *CashAccount24 {
	m.DebtorAccount = new(CashAccount24)
	return m.DebtorAccount
}

func (m *Mandate8) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.DebtorAgent
}

func (m *Mandate8) AddUltimateDebtor() *PartyIdentification43 {
	m.UltimateDebtor = new(PartyIdentification43)
	return m.UltimateDebtor
}

func (m *Mandate8) SetMandateReference(value string) {
	m.MandateReference = (*Max35Text)(&value)
}

func (m *Mandate8) AddReferredDocument() *ReferredMandateDocument1 {
	newValue := new(ReferredMandateDocument1)
	m.ReferredDocument = append(m.ReferredDocument, newValue)
	return newValue
}
