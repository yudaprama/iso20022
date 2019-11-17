package model

// A safekeeping account is an account on which a securities entry is made.
type SafekeepingAccount6 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification40Choice `xml:"AcctOwnr,omitempty"`

	// Identification of a subaccount within the safekeeping account
	SubAccountDetails *SubAccount2 `xml:"SubAcctDtls,omitempty"`

	// Quantity of securities in the sub-balance.
	InstructedBalance []*HoldingBalance8 `xml:"InstdBal"`

	// Owner of the voting rights.
	RightsHolder []*PartyIdentification40Choice `xml:"RghtsHldr,omitempty"`
}

func (s *SafekeepingAccount6) SetAccountIdentification(value string) {
	s.AccountIdentification = (*Max35Text)(&value)
}

func (s *SafekeepingAccount6) AddAccountOwner() *PartyIdentification40Choice {
	s.AccountOwner = new(PartyIdentification40Choice)
	return s.AccountOwner
}

func (s *SafekeepingAccount6) AddSubAccountDetails() *SubAccount2 {
	s.SubAccountDetails = new(SubAccount2)
	return s.SubAccountDetails
}

func (s *SafekeepingAccount6) AddInstructedBalance() *HoldingBalance8 {
	newValue := new(HoldingBalance8)
	s.InstructedBalance = append(s.InstructedBalance, newValue)
	return newValue
}

func (s *SafekeepingAccount6) AddRightsHolder() *PartyIdentification40Choice {
	newValue := new(PartyIdentification40Choice)
	s.RightsHolder = append(s.RightsHolder, newValue)
	return newValue
}
