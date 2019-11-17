package model

// Provides the non guaranteed trade details.
type NonGuaranteedTrade3 struct {

	// External identification of the member who is the market counterpart member of the current trade leg (in case of non guarantee trades, this field allows buyer and seller to identify each other).
	TradeCounterpartyMemberIdentification *PartyIdentification35Choice `xml:"TradCtrPtyMmbId"`

	// External identification of the clearing member of the market couterpart member (in case of non guarantee trades, this field allows buyer and seller to identify each other).
	TradeCounterpartyClearingMemberIdentification *PartyIdentification35Choice `xml:"TradCtrPtyClrMmbId"`

	// Provides details about the delivering parties involved in the settlement chain.
	DeliveringParties *DeliveringPartiesAndAccount11 `xml:"DlvrgPties,omitempty"`

	// Provides details about the receiving parties involved in the settlement chain.
	ReceivingParties *ReceivingPartiesAndAccount11 `xml:"RcvgPties,omitempty"`
}

func (n *NonGuaranteedTrade3) AddTradeCounterpartyMemberIdentification() *PartyIdentification35Choice {
	n.TradeCounterpartyMemberIdentification = new(PartyIdentification35Choice)
	return n.TradeCounterpartyMemberIdentification
}

func (n *NonGuaranteedTrade3) AddTradeCounterpartyClearingMemberIdentification() *PartyIdentification35Choice {
	n.TradeCounterpartyClearingMemberIdentification = new(PartyIdentification35Choice)
	return n.TradeCounterpartyClearingMemberIdentification
}

func (n *NonGuaranteedTrade3) AddDeliveringParties() *DeliveringPartiesAndAccount11 {
	n.DeliveringParties = new(DeliveringPartiesAndAccount11)
	return n.DeliveringParties
}

func (n *NonGuaranteedTrade3) AddReceivingParties() *ReceivingPartiesAndAccount11 {
	n.ReceivingParties = new(ReceivingPartiesAndAccount11)
	return n.ReceivingParties
}
