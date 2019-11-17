package model

// Account and holding of the next sub-level (Level 5).
type AccountSubLevel15 struct {

	// Unique and unambiguous identification for the sub-account between the account owner and the account servicer.
	AccountIdentification *SecuritiesAccount19 `xml:"AcctId"`

	// Party that legally owns the sub-account.
	AccountOwner *PartyIdentification100 `xml:"AcctOwnr"`

	// Party that manages the sub-level account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification100 `xml:"AcctSvcr"`

	// Individual or entity that is ultimately entitled to the benefit of income and rights in a financial instrument, as opposed to a nominal or legal owner.
	BeneficialOwner []*BeneficialOwner2 `xml:"BnfclOwnr,omitempty"`

	// Report on the net position of a financial instrument on the sub-account (sub-account level 5),  for a certain date. The agent, for example, a trade intermediary, may also be specified.
	BalanceDetails []*AggregateHoldingBalance3 `xml:"BalDtls,omitempty"`

	// Holdings of level 5.
	AccountSubLevel6 []*AccountSubLevel16 `xml:"AcctSubLvl6,omitempty"`

	// Difference in holdings between the safekeeping account and the sub-accounts of level 6.
	AccountSubLevel6Difference []*AggregateHoldingBalance2 `xml:"AcctSubLvl6Diff,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AccountSubLevel15) AddAccountIdentification() *SecuritiesAccount19 {
	a.AccountIdentification = new(SecuritiesAccount19)
	return a.AccountIdentification
}

func (a *AccountSubLevel15) AddAccountOwner() *PartyIdentification100 {
	a.AccountOwner = new(PartyIdentification100)
	return a.AccountOwner
}

func (a *AccountSubLevel15) AddAccountServicer() *PartyIdentification100 {
	a.AccountServicer = new(PartyIdentification100)
	return a.AccountServicer
}

func (a *AccountSubLevel15) AddBeneficialOwner() *BeneficialOwner2 {
	newValue := new(BeneficialOwner2)
	a.BeneficialOwner = append(a.BeneficialOwner, newValue)
	return newValue
}

func (a *AccountSubLevel15) AddBalanceDetails() *AggregateHoldingBalance3 {
	newValue := new(AggregateHoldingBalance3)
	a.BalanceDetails = append(a.BalanceDetails, newValue)
	return newValue
}

func (a *AccountSubLevel15) AddAccountSubLevel6() *AccountSubLevel16 {
	newValue := new(AccountSubLevel16)
	a.AccountSubLevel6 = append(a.AccountSubLevel6, newValue)
	return newValue
}

func (a *AccountSubLevel15) AddAccountSubLevel6Difference() *AggregateHoldingBalance2 {
	newValue := new(AggregateHoldingBalance2)
	a.AccountSubLevel6Difference = append(a.AccountSubLevel6Difference, newValue)
	return newValue
}

func (a *AccountSubLevel15) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
