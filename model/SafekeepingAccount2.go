package model

// Safekeeping or investment account. A safekeeping account is an account on which a securities entry is made. An investment account is an account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type SafekeepingAccount2 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentificationFormatChoice `xml:"Id"`

	// Indicates whether the securities in the account are fungible, ie, interchangeable.
	FungibleIndicator *YesNoIndicator `xml:"FngbInd"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Party that provides services relating to financial products to investors, eg, advice on products and placement of orders for the investment fund.
	IntermediaryInformation []*Intermediary11 `xml:"IntrmyInf,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification2Choice `xml:"AcctOwnr,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`
}

func (s *SafekeepingAccount2) AddIdentification() *AccountIdentificationFormatChoice {
	s.Identification = new(AccountIdentificationFormatChoice)
	return s.Identification
}

func (s *SafekeepingAccount2) SetFungibleIndicator(value string) {
	s.FungibleIndicator = (*YesNoIndicator)(&value)
}

func (s *SafekeepingAccount2) SetName(value string) {
	s.Name = (*Max35Text)(&value)
}

func (s *SafekeepingAccount2) SetDesignation(value string) {
	s.Designation = (*Max35Text)(&value)
}

func (s *SafekeepingAccount2) AddIntermediaryInformation() *Intermediary11 {
	newValue := new(Intermediary11)
	s.IntermediaryInformation = append(s.IntermediaryInformation, newValue)
	return newValue
}

func (s *SafekeepingAccount2) AddAccountOwner() *PartyIdentification2Choice {
	s.AccountOwner = new(PartyIdentification2Choice)
	return s.AccountOwner
}

func (s *SafekeepingAccount2) AddAccountServicer() *PartyIdentification2Choice {
	s.AccountServicer = new(PartyIdentification2Choice)
	return s.AccountServicer
}
