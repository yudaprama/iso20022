package model

// Payment instrument between a debtor and a creditor, which flows through one or more financial institutions or systems.
type CreditTransfer6 struct {

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	Reference *Max35Text `xml:"Ref,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor. In the context of the payment model, the debtor is also the debit account owner.
	Debtor *PartyIdentification2Choice `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *AccountIdentificationAndName3 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *FinancialInstitutionIdentification3Choice `xml:"DbtrAgt,omitempty"`

	// Identifies the account of the debtor's agent.
	DebtorAgentAccount *AccountIdentificationAndName3 `xml:"DbtrAgtAcct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	IntermediaryAgent1 *FinancialInstitutionIdentification3Choice `xml:"IntrmyAgt1,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 1 at its servicing agent in the payment chain.
	IntermediaryAgent1Account *AccountIdentificationAndName3 `xml:"IntrmyAgt1Acct,omitempty"`

	// Agent between the debtor's agent and the creditor's agent.
	IntermediaryAgent2 *FinancialInstitutionIdentification3Choice `xml:"IntrmyAgt2,omitempty"`

	// Unambiguous identification of the account of the intermediary agent 2 at its servicing agent in the payment chain.
	IntermediaryAgent2Account *AccountIdentificationAndName3 `xml:"IntrmyAgt2Acct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *FinancialInstitutionIdentification3Choice `xml:"CdtrAgt"`

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CreditorAgentAccount *AccountIdentificationAndName3 `xml:"CdtrAgtAcct,omitempty"`

	// Party that receives an amount of money from the debtor. In the context of the payment model, the creditor is also the credit account owner.
	Creditor *PartyIdentification2Choice `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *AccountIdentificationAndName3 `xml:"CdtrAcct"`
}

func (c *CreditTransfer6) SetReference(value string) {
	c.Reference = (*Max35Text)(&value)
}

func (c *CreditTransfer6) AddDebtor() *PartyIdentification2Choice {
	c.Debtor = new(PartyIdentification2Choice)
	return c.Debtor
}

func (c *CreditTransfer6) AddDebtorAccount() *AccountIdentificationAndName3 {
	c.DebtorAccount = new(AccountIdentificationAndName3)
	return c.DebtorAccount
}

func (c *CreditTransfer6) AddDebtorAgent() *FinancialInstitutionIdentification3Choice {
	c.DebtorAgent = new(FinancialInstitutionIdentification3Choice)
	return c.DebtorAgent
}

func (c *CreditTransfer6) AddDebtorAgentAccount() *AccountIdentificationAndName3 {
	c.DebtorAgentAccount = new(AccountIdentificationAndName3)
	return c.DebtorAgentAccount
}

func (c *CreditTransfer6) AddIntermediaryAgent1() *FinancialInstitutionIdentification3Choice {
	c.IntermediaryAgent1 = new(FinancialInstitutionIdentification3Choice)
	return c.IntermediaryAgent1
}

func (c *CreditTransfer6) AddIntermediaryAgent1Account() *AccountIdentificationAndName3 {
	c.IntermediaryAgent1Account = new(AccountIdentificationAndName3)
	return c.IntermediaryAgent1Account
}

func (c *CreditTransfer6) AddIntermediaryAgent2() *FinancialInstitutionIdentification3Choice {
	c.IntermediaryAgent2 = new(FinancialInstitutionIdentification3Choice)
	return c.IntermediaryAgent2
}

func (c *CreditTransfer6) AddIntermediaryAgent2Account() *AccountIdentificationAndName3 {
	c.IntermediaryAgent2Account = new(AccountIdentificationAndName3)
	return c.IntermediaryAgent2Account
}

func (c *CreditTransfer6) AddCreditorAgent() *FinancialInstitutionIdentification3Choice {
	c.CreditorAgent = new(FinancialInstitutionIdentification3Choice)
	return c.CreditorAgent
}

func (c *CreditTransfer6) AddCreditorAgentAccount() *AccountIdentificationAndName3 {
	c.CreditorAgentAccount = new(AccountIdentificationAndName3)
	return c.CreditorAgentAccount
}

func (c *CreditTransfer6) AddCreditor() *PartyIdentification2Choice {
	c.Creditor = new(PartyIdentification2Choice)
	return c.Creditor
}

func (c *CreditTransfer6) AddCreditorAccount() *AccountIdentificationAndName3 {
	c.CreditorAccount = new(AccountIdentificationAndName3)
	return c.CreditorAccount
}
