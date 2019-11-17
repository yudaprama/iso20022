package model

// A safekeeping account is an account on which a securities entry is made.
type SafekeepingAccount3 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification9Choice `xml:"AcctOwnr,omitempty"`

	// Identification of a subaccount within the safekeeping account
	SubAccountDetails *SubAccount2 `xml:"SubAcctDtls,omitempty"`

	// Quantity of securities in the sub-balance.
	InstructedBalance []*HoldingBalance4 `xml:"InstdBal"`

	// Owner of the voting rights.
	RightsHolder []*PartyIdentification9Choice `xml:"RghtsHldr,omitempty"`
}

func (s *SafekeepingAccount3) SetAccountIdentification(value string) {
	s.AccountIdentification = (*Max35Text)(&value)
}

func (s *SafekeepingAccount3) AddAccountOwner() *PartyIdentification9Choice {
	s.AccountOwner = new(PartyIdentification9Choice)
	return s.AccountOwner
}

func (s *SafekeepingAccount3) AddSubAccountDetails() *SubAccount2 {
	s.SubAccountDetails = new(SubAccount2)
	return s.SubAccountDetails
}

func (s *SafekeepingAccount3) AddInstructedBalance() *HoldingBalance4 {
	newValue := new(HoldingBalance4)
	s.InstructedBalance = append(s.InstructedBalance, newValue)
	return newValue
}

func (s *SafekeepingAccount3) AddRightsHolder() *PartyIdentification9Choice {
	newValue := new(PartyIdentification9Choice)
	s.RightsHolder = append(s.RightsHolder, newValue)
	return newValue
}
