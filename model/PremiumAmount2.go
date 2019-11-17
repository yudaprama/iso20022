package model

// Commercial agreement in which the buyer agrees to pay the seller an amount of cash. Some aspects of the payment may be defined in the agreement, eg, the method of the payment
type PremiumAmount2 struct {

	// Specifies the calculation method of the premium amount.
	PremiumQuote *PremiumQuote1Choice `xml:"PrmQt,omitempty"`

	// Result of the calculation of the premium amount on the basis of the premium quote and one of the amounts of the underlying foreign exchange trade.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Date on which the premium must be settled.
	SettlementDate *ISODate `xml:"SttlmDt"`

	// Party that settles the premium amount on behalf of the paying party.  It may contain the BIC of a central settlement system, eg. CLSBUS33.
	SettlementParty *PartyIdentification8Choice `xml:"SttlmPty,omitempty"`
}

func (p *PremiumAmount2) AddPremiumQuote() *PremiumQuote1Choice {
	p.PremiumQuote = new(PremiumQuote1Choice)
	return p.PremiumQuote
}

func (p *PremiumAmount2) SetAmount(value, currency string) {
	p.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PremiumAmount2) SetSettlementDate(value string) {
	p.SettlementDate = (*ISODate)(&value)
}

func (p *PremiumAmount2) AddSettlementParty() *PartyIdentification8Choice {
	p.SettlementParty = new(PartyIdentification8Choice)
	return p.SettlementParty
}
