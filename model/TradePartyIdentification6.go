package model

// Entity involved in an activity.
type TradePartyIdentification6 struct {

	// Party that submits the foreign exchange trade to the matching system or to the settlement system or to the counterparty.
	SubmittingParty *PartyIdentification73Choice `xml:"SubmitgPty"`

	// Party that originated the foreign exchange trade. This party may be the same as the submitting party.
	TradeParty *PartyIdentification73Choice `xml:"TradPty,omitempty"`

	// Identifies the fund that is one of the parties in the foreign exchange trade.
	//
	FundIdentification []*FundIdentification4 `xml:"FndId,omitempty"`
}

func (t *TradePartyIdentification6) AddSubmittingParty() *PartyIdentification73Choice {
	t.SubmittingParty = new(PartyIdentification73Choice)
	return t.SubmittingParty
}

func (t *TradePartyIdentification6) AddTradeParty() *PartyIdentification73Choice {
	t.TradeParty = new(PartyIdentification73Choice)
	return t.TradeParty
}

func (t *TradePartyIdentification6) AddFundIdentification() *FundIdentification4 {
	newValue := new(FundIdentification4)
	t.FundIdentification = append(t.FundIdentification, newValue)
	return newValue
}
