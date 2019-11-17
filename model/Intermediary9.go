package model

// Party that provides services to investors relating to financial products.
type Intermediary9 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification2Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account7 `xml:"Acct,omitempty"`

	// Counterparties eligibility as defined by article 24 of the EU MiFID Directive applicable to transactions executed by investment firms for eligible counterparties.
	OrderOriginatorEligibility *OrderOriginatorEligibility1Code `xml:"OrdrOrgtrElgblty,omitempty"`

	// Capacity of the party executing an order.
	TradingPartyCapacity *TradingCapacity2Code `xml:"TradgPtyCpcty,omitempty"`

	// Function performed by the intermediary.
	Role *InvestmentFundRole2Code `xml:"Role,omitempty"`

	// Function performed by the intermediary.
	ExtendedRole *Extended350Code `xml:"XtndedRole,omitempty"`
}

func (i *Intermediary9) AddIdentification() *PartyIdentification2Choice {
	i.Identification = new(PartyIdentification2Choice)
	return i.Identification
}

func (i *Intermediary9) AddAccount() *Account7 {
	i.Account = new(Account7)
	return i.Account
}

func (i *Intermediary9) SetOrderOriginatorEligibility(value string) {
	i.OrderOriginatorEligibility = (*OrderOriginatorEligibility1Code)(&value)
}

func (i *Intermediary9) SetTradingPartyCapacity(value string) {
	i.TradingPartyCapacity = (*TradingCapacity2Code)(&value)
}

func (i *Intermediary9) SetRole(value string) {
	i.Role = (*InvestmentFundRole2Code)(&value)
}

func (i *Intermediary9) SetExtendedRole(value string) {
	i.ExtendedRole = (*Extended350Code)(&value)
}
