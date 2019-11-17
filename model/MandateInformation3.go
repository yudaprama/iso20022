package model

// Information that serves as a basis to debit an account.
type MandateInformation3 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId"`

	// Identification for the mandate request, as assigned by the initiating party.
	MandateRequestIdentification *Max35Text `xml:"MndtReqId,omitempty"`

	// Specifies the type of mandate, such as paper, electronic or scheme.
	Type *MandateTypeInformation1 `xml:"Tp,omitempty"`

	// Set of elements used to provide details of the duration of the mandate and occurrence of the underlying transactions.
	Occurrences *MandateOccurrences1 `xml:"Ocrncs,omitempty"`

	// Fixed amount to be collected from the debtor's account.
	CollectionAmount *ActiveCurrencyAndAmount `xml:"ColltnAmt,omitempty"`

	// Maximum amount that may be collected from the debtor's account, per instruction.
	MaximumAmount *ActiveCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification32 `xml:"CdtrSchmeId,omitempty"`

	// Party that signs the mandate and to whom an amount of money is due.
	Creditor *PartyIdentification32 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount16 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"CdtrAgt,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification32 `xml:"UltmtCdtr,omitempty"`

	// Party that signs the mandate and owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification32 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor, to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount16 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification4 `xml:"DbtrAgt,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification32 `xml:"UltmtDbtr,omitempty"`

	// Set of elements used to provide information to identify the underlying documents associated with the mandate.
	ReferredDocument *ReferredDocumentInformation3 `xml:"RfrdDoc,omitempty"`
}

func (m *MandateInformation3) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *MandateInformation3) SetMandateRequestIdentification(value string) {
	m.MandateRequestIdentification = (*Max35Text)(&value)
}

func (m *MandateInformation3) AddType() *MandateTypeInformation1 {
	m.Type = new(MandateTypeInformation1)
	return m.Type
}

func (m *MandateInformation3) AddOccurrences() *MandateOccurrences1 {
	m.Occurrences = new(MandateOccurrences1)
	return m.Occurrences
}

func (m *MandateInformation3) SetCollectionAmount(value, currency string) {
	m.CollectionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MandateInformation3) SetMaximumAmount(value, currency string) {
	m.MaximumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MandateInformation3) AddCreditorSchemeIdentification() *PartyIdentification32 {
	m.CreditorSchemeIdentification = new(PartyIdentification32)
	return m.CreditorSchemeIdentification
}

func (m *MandateInformation3) AddCreditor() *PartyIdentification32 {
	m.Creditor = new(PartyIdentification32)
	return m.Creditor
}

func (m *MandateInformation3) AddCreditorAccount() *CashAccount16 {
	m.CreditorAccount = new(CashAccount16)
	return m.CreditorAccount
}

func (m *MandateInformation3) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification4 {
	m.CreditorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return m.CreditorAgent
}

func (m *MandateInformation3) AddUltimateCreditor() *PartyIdentification32 {
	m.UltimateCreditor = new(PartyIdentification32)
	return m.UltimateCreditor
}

func (m *MandateInformation3) AddDebtor() *PartyIdentification32 {
	m.Debtor = new(PartyIdentification32)
	return m.Debtor
}

func (m *MandateInformation3) AddDebtorAccount() *CashAccount16 {
	m.DebtorAccount = new(CashAccount16)
	return m.DebtorAccount
}

func (m *MandateInformation3) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification4 {
	m.DebtorAgent = new(BranchAndFinancialInstitutionIdentification4)
	return m.DebtorAgent
}

func (m *MandateInformation3) AddUltimateDebtor() *PartyIdentification32 {
	m.UltimateDebtor = new(PartyIdentification32)
	return m.UltimateDebtor
}

func (m *MandateInformation3) AddReferredDocument() *ReferredDocumentInformation3 {
	m.ReferredDocument = new(ReferredDocumentInformation3)
	return m.ReferredDocument
}
