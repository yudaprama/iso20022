package model

// Information about the debtor.
type Debtor2 struct {

	// Party that owes an amount of money to the (ultimate) creditor. In the context of the payment model, the debtor is also the debit account owner.
	Debtor *PartyIdentification2Choice `xml:"Dbtr,omitempty"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentificationAndName3 `xml:"AcctId"`

	// Financial institution that receives the payment transaction from the account owner, or other authorised party, and processes the instruction.
	FirstAgent *FinancialInstitutionIdentification3Choice `xml:"FrstAgt"`
}

func (d *Debtor2) AddDebtor() *PartyIdentification2Choice {
	d.Debtor = new(PartyIdentification2Choice)
	return d.Debtor
}

func (d *Debtor2) AddAccountIdentification() *AccountIdentificationAndName3 {
	d.AccountIdentification = new(AccountIdentificationAndName3)
	return d.AccountIdentification
}

func (d *Debtor2) AddFirstAgent() *FinancialInstitutionIdentification3Choice {
	d.FirstAgent = new(FinancialInstitutionIdentification3Choice)
	return d.FirstAgent
}
