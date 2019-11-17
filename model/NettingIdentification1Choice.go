package model

// Choice between netting group identification of an individual trading party.
type NettingIdentification1Choice struct {

	// Describes the individual trading party.
	TradeParty *PartyIdentification73Choice `xml:"TradPty"`

	// Describes the netting group.
	NettingGroupIdentification *Max35Text `xml:"NetgGrpId"`
}

func (n *NettingIdentification1Choice) AddTradeParty() *PartyIdentification73Choice {
	n.TradeParty = new(PartyIdentification73Choice)
	return n.TradeParty
}

func (n *NettingIdentification1Choice) SetNettingGroupIdentification(value string) {
	n.NettingGroupIdentification = (*Max35Text)(&value)
}
