package model

// Account on which a securities entry is made.
type SafekeepingAccount7 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *SecuritiesAccount19 `xml:"AcctId"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification100 `xml:"AcctOwnr"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification100 `xml:"AcctSvcr"`

	// Individual or entity that is ultimately entitled to the benefit of income and rights in a financial instrument, as opposed to a nominal or legal owner.
	BeneficialOwner []*BeneficialOwner2 `xml:"BnfclOwnr,omitempty"`

	// Report on the net position of a financial instrument on the account,  for a certain date. The agent, for example, a trade intermediary, may also be specified.
	BalanceDetails []*AggregateHoldingBalance3 `xml:"BalDtls,omitempty"`

	// Holdings of level 1.
	AccountSubLevel1 []*AccountSubLevel11 `xml:"AcctSubLvl1,omitempty"`

	// Difference in holdings between the safekeeping account and the sub-accounts of level 1.
	AccountSubLevel1Difference []*AggregateHoldingBalance2 `xml:"AcctSubLvl1Diff,omitempty"`
}

func (s *SafekeepingAccount7) AddAccountIdentification() *SecuritiesAccount19 {
	s.AccountIdentification = new(SecuritiesAccount19)
	return s.AccountIdentification
}

func (s *SafekeepingAccount7) AddAccountOwner() *PartyIdentification100 {
	s.AccountOwner = new(PartyIdentification100)
	return s.AccountOwner
}

func (s *SafekeepingAccount7) AddAccountServicer() *PartyIdentification100 {
	s.AccountServicer = new(PartyIdentification100)
	return s.AccountServicer
}

func (s *SafekeepingAccount7) AddBeneficialOwner() *BeneficialOwner2 {
	newValue := new(BeneficialOwner2)
	s.BeneficialOwner = append(s.BeneficialOwner, newValue)
	return newValue
}

func (s *SafekeepingAccount7) AddBalanceDetails() *AggregateHoldingBalance3 {
	newValue := new(AggregateHoldingBalance3)
	s.BalanceDetails = append(s.BalanceDetails, newValue)
	return newValue
}

func (s *SafekeepingAccount7) AddAccountSubLevel1() *AccountSubLevel11 {
	newValue := new(AccountSubLevel11)
	s.AccountSubLevel1 = append(s.AccountSubLevel1, newValue)
	return newValue
}

func (s *SafekeepingAccount7) AddAccountSubLevel1Difference() *AggregateHoldingBalance2 {
	newValue := new(AggregateHoldingBalance2)
	s.AccountSubLevel1Difference = append(s.AccountSubLevel1Difference, newValue)
	return newValue
}
