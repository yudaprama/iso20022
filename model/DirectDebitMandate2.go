package model

// Instruction, initiated by the creditor, to debit a debtor's account in favour of the creditor. A direct debit can be pre-authorised or not. In most countries, authorisation is in the form of a mandate between the debtor and creditor.
type DirectDebitMandate2 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	DebtorAccountIdentification *CashAccountIdentification1Choice `xml:"DbtrAcctId"`

	// Party that owes the cash to the creditor/final party. The debtor is also the debit account owner.
	DebtorIdentification *PartyIdentification2Choice `xml:"DbtrId,omitempty"`

	// Party that receives an amount of money from the debtor. In the context of the payment model, the creditor is also the credit account owner.
	CreditorIdentification *PartyIdentification2Choice `xml:"CdtrId,omitempty"`

	// Financial institution that receives the direct debit instruction from the creditor or other authorised party.
	FirstAgent *FinancialInstitutionIdentification3Choice `xml:"FrstAgt"`

	// Financial institution that receives the payment transaction on behalf of the creditor, or other nominated party, and credits the account.
	FinalAgent *FinancialInstitutionIdentification3Choice `xml:"FnlAgt,omitempty"`

	// Reference assigned to a creditor by its financial institution, or relevant authority, authorising the creditor to take part in a direct debit scheme.
	RegistrationIdentification *Max35Text `xml:"RegnId,omitempty"`

	// Reference of the direct debit mandate that has been agreed upon by the debtor and creditor.
	MandateIdentification *Max35Text `xml:"MndtId,omitempty"`
}

func (d *DirectDebitMandate2) AddDebtorAccountIdentification() *CashAccountIdentification1Choice {
	d.DebtorAccountIdentification = new(CashAccountIdentification1Choice)
	return d.DebtorAccountIdentification
}

func (d *DirectDebitMandate2) AddDebtorIdentification() *PartyIdentification2Choice {
	d.DebtorIdentification = new(PartyIdentification2Choice)
	return d.DebtorIdentification
}

func (d *DirectDebitMandate2) AddCreditorIdentification() *PartyIdentification2Choice {
	d.CreditorIdentification = new(PartyIdentification2Choice)
	return d.CreditorIdentification
}

func (d *DirectDebitMandate2) AddFirstAgent() *FinancialInstitutionIdentification3Choice {
	d.FirstAgent = new(FinancialInstitutionIdentification3Choice)
	return d.FirstAgent
}

func (d *DirectDebitMandate2) AddFinalAgent() *FinancialInstitutionIdentification3Choice {
	d.FinalAgent = new(FinancialInstitutionIdentification3Choice)
	return d.FinalAgent
}

func (d *DirectDebitMandate2) SetRegistrationIdentification(value string) {
	d.RegistrationIdentification = (*Max35Text)(&value)
}

func (d *DirectDebitMandate2) SetMandateIdentification(value string) {
	d.MandateIdentification = (*Max35Text)(&value)
}
