package model

// Instruction, initiated by the creditor, to debit a debtor's account in favour of the creditor. A direct debit can be pre-authorised or not. In most countries, authorisation is in the form of a mandate between the debtor and creditor.
type DirectDebitMandate6 struct {

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *AccountIdentificationAndName5 `xml:"DbtrAcct"`

	// Party that owes the cash to the creditor/final party. The debtor is also the debit account owner.
	Debtor *PartyIdentification113 `xml:"Dbtr,omitempty"`

	// Number assigned by a tax authority to an entity.
	DebtorTaxIdentificationNumber *Max35Text `xml:"DbtrTaxIdNb,omitempty"`

	// Number assigned by a national registration authority to an entity.
	DebtorNationalRegistrationNumber *Max35Text `xml:"DbtrNtlRegnNb,omitempty"`

	// Party that receives an amount of money from the debtor. In the context of the payment model, the creditor is also the credit account owner.
	Creditor *PartyIdentification113 `xml:"Cdtr,omitempty"`

	// Financial institution that receives the direct debit instruction from the creditor or other authorised party.
	DebtorAgent *FinancialInstitutionIdentification10 `xml:"DbtrAgt"`

	// Information identifying a specific branch of a financial institution.
	//
	// Usage : this component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	DebtorAgentBranch *BranchData `xml:"DbtrAgtBrnch,omitempty"`

	// Financial institution that receives the payment transaction on behalf of the creditor, or other nominated party, and credits the account.
	CreditorAgent *FinancialInstitutionIdentification10 `xml:"CdtrAgt,omitempty"`

	// Information identifying a specific branch of a financial institution.
	//
	// Usage : this component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	CreditorAgentBranch *BranchData `xml:"CdtrAgtBrnch,omitempty"`

	// Reference assigned to a creditor by its financial institution, or relevant authority, authorising the creditor to take part in a direct debit scheme.
	RegistrationIdentification *Max35Text `xml:"RegnId,omitempty"`

	// Reference of the direct debit mandate that has been agreed upon by the debtor and creditor.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`
}

func (d *DirectDebitMandate6) AddDebtorAccount() *AccountIdentificationAndName5 {
	d.DebtorAccount = new(AccountIdentificationAndName5)
	return d.DebtorAccount
}

func (d *DirectDebitMandate6) AddDebtor() *PartyIdentification113 {
	d.Debtor = new(PartyIdentification113)
	return d.Debtor
}

func (d *DirectDebitMandate6) SetDebtorTaxIdentificationNumber(value string) {
	d.DebtorTaxIdentificationNumber = (*Max35Text)(&value)
}

func (d *DirectDebitMandate6) SetDebtorNationalRegistrationNumber(value string) {
	d.DebtorNationalRegistrationNumber = (*Max35Text)(&value)
}

func (d *DirectDebitMandate6) AddCreditor() *PartyIdentification113 {
	d.Creditor = new(PartyIdentification113)
	return d.Creditor
}

func (d *DirectDebitMandate6) AddDebtorAgent() *FinancialInstitutionIdentification10 {
	d.DebtorAgent = new(FinancialInstitutionIdentification10)
	return d.DebtorAgent
}

func (d *DirectDebitMandate6) AddDebtorAgentBranch() *BranchData {
	d.DebtorAgentBranch = new(BranchData)
	return d.DebtorAgentBranch
}

func (d *DirectDebitMandate6) AddCreditorAgent() *FinancialInstitutionIdentification10 {
	d.CreditorAgent = new(FinancialInstitutionIdentification10)
	return d.CreditorAgent
}

func (d *DirectDebitMandate6) AddCreditorAgentBranch() *BranchData {
	d.CreditorAgentBranch = new(BranchData)
	return d.CreditorAgentBranch
}

func (d *DirectDebitMandate6) SetRegistrationIdentification(value string) {
	d.RegistrationIdentification = (*Max35Text)(&value)
}

func (d *DirectDebitMandate6) SetMandateIdentification(value string) {
	d.MandateIdentification = (*Max35Text)(&value)
}
