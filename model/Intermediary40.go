package model

// Party that provides services to investors relating to financial products.
type Intermediary40 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification113 `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account22 `xml:"Acct,omitempty"`

	// Counterparties eligibility as defined by article 24 of the EU MiFID Directive applicable to transactions executed by investment firms for eligible counterparties.
	OrderOriginatorEligibility *OrderOriginatorEligibility1Code `xml:"OrdrOrgtrElgblty,omitempty"`

	// Function performed by the intermediary.
	Role *InvestmentFundRole2Choice `xml:"Role,omitempty"`
}

func (i *Intermediary40) AddIdentification() *PartyIdentification113 {
	i.Identification = new(PartyIdentification113)
	return i.Identification
}

func (i *Intermediary40) AddAccount() *Account22 {
	i.Account = new(Account22)
	return i.Account
}

func (i *Intermediary40) SetOrderOriginatorEligibility(value string) {
	i.OrderOriginatorEligibility = (*OrderOriginatorEligibility1Code)(&value)
}

func (i *Intermediary40) AddRole() *InvestmentFundRole2Choice {
	i.Role = new(InvestmentFundRole2Choice)
	return i.Role
}
