package model

// Cash settlement chain parties and accounts.
type CashParties24 struct {

	// Party to which the payment amount must be ultimately delivered. In some cases, this may be a fund.
	//
	Creditor *PartyIdentificationAndAccount96 `xml:"Cdtr"`

	// Financial institution that services the cash account of the beneficiary (creditor). In some markets, this is also known as receiving agent. The creditor agent is the party where the payment amount must be ultimately delivered on behalf of the beneficiary (creditor), that is, the party where the beneficiary has its account.
	CreditorAgent *PartyIdentificationAndAccount97 `xml:"CdtrAgt"`

	// Financial institution through which the transaction must pass to reach the account with institution (creditor agent).
	Intermediary *PartyIdentificationAndAccount97 `xml:"Intrmy,omitempty"`

	// Financial institution through which the transaction must pass to reach the account with institution (creditor agent).
	Intermediary2 *PartyIdentificationAndAccount97 `xml:"Intrmy2,omitempty"`
}

func (c *CashParties24) AddCreditor() *PartyIdentificationAndAccount96 {
	c.Creditor = new(PartyIdentificationAndAccount96)
	return c.Creditor
}

func (c *CashParties24) AddCreditorAgent() *PartyIdentificationAndAccount97 {
	c.CreditorAgent = new(PartyIdentificationAndAccount97)
	return c.CreditorAgent
}

func (c *CashParties24) AddIntermediary() *PartyIdentificationAndAccount97 {
	c.Intermediary = new(PartyIdentificationAndAccount97)
	return c.Intermediary
}

func (c *CashParties24) AddIntermediary2() *PartyIdentificationAndAccount97 {
	c.Intermediary2 = new(PartyIdentificationAndAccount97)
	return c.Intermediary2
}
