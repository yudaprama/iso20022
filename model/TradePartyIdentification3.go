package model

// Entity involved in an activity.
type TradePartyIdentification3 struct {

	// Identifies the fund which is one of the parties in a treasury trade.
	FundInformation *FundIdentification2 `xml:"FndInf,omitempty"`

	// Specifies the party which submits a treasury trade to a matching system or to a settlement system or to a counterparty.
	SubmittingParty *PartyIdentification8Choice `xml:"SubmitgPty"`

	// Specifies the party which originated a treasury trade. This party may be the same as the submitting party.
	TradeParty *PartyIdentification8Choice `xml:"TradPty"`
}

func (t *TradePartyIdentification3) AddFundInformation() *FundIdentification2 {
	t.FundInformation = new(FundIdentification2)
	return t.FundInformation
}

func (t *TradePartyIdentification3) AddSubmittingParty() *PartyIdentification8Choice {
	t.SubmittingParty = new(PartyIdentification8Choice)
	return t.SubmittingParty
}

func (t *TradePartyIdentification3) AddTradeParty() *PartyIdentification8Choice {
	t.TradeParty = new(PartyIdentification8Choice)
	return t.TradeParty
}
