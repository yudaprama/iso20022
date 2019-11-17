package model

// Information about the creditor.
type Creditor2 struct {

	// Party that receives an amount of money from the debtor. In the context of the payment model, the creditor is also the credit account owner.
	Creditor *PartyIdentification2Choice `xml:"Cdtr,omitempty"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentificationAndName3 `xml:"AcctId"`

	// Party that is the ultimate beneficiary of the credit transfer. The final party is mentioned when different from the creditor, whose account will be credited by the final agent.
	FinalAgent *FinancialInstitutionIdentification3Choice `xml:"FnlAgt"`
}

func (c *Creditor2) AddCreditor() *PartyIdentification2Choice {
	c.Creditor = new(PartyIdentification2Choice)
	return c.Creditor
}

func (c *Creditor2) AddAccountIdentification() *AccountIdentificationAndName3 {
	c.AccountIdentification = new(AccountIdentificationAndName3)
	return c.AccountIdentification
}

func (c *Creditor2) AddFinalAgent() *FinancialInstitutionIdentification3Choice {
	c.FinalAgent = new(FinancialInstitutionIdentification3Choice)
	return c.FinalAgent
}
