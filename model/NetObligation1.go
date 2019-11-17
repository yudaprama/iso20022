package model

// Describes the amount, direction and parties involved in a payment obligation between two participants (and their netting group or trading party) of a netting service.
type NetObligation1 struct {

	// Unique identification for the obligation.
	ObligationIdentification *Max35Text `xml:"OblgtnId"`

	// Amount and currency of the obligation
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Describes the party or netting group (of the participant receiving the report) involved in the calculation of the obligation.
	ParticipantNettingIdentification *NettingIdentification1Choice `xml:"PtcptNetgId"`

	// Specifies the direction of the obligation.
	ObligationDirection *CreditDebit3Code `xml:"OblgtnDrctn"`

	// Describes the party or netting group (of the counterparty in the obligation) involved in the calculation of the obligation.
	CounterpartyNettingIdentification *NettingIdentification1Choice `xml:"CtrPtyNetgId"`

	// Describes the counterparty participant involved in the obligation.
	NetServiceCounterpartyIdentification *PartyIdentification73Choice `xml:"NetSvcCtrPtyId,omitempty"`

	// Specifies the standard settlement instructions used to issue payment to the counterparty in order to settle the obligation.
	CounterpartySettlementInstructions *SettlementParties29 `xml:"CtrPtySttlmInstrs,omitempty"`

	// Number of transactions used to calculate the obligation. This is used in reconciliation between the net report obligation and the previously provided transaction status updates.
	TransactionsNumber *Max10NumericText `xml:"TxsNb,omitempty"`
}

func (n *NetObligation1) SetObligationIdentification(value string) {
	n.ObligationIdentification = (*Max35Text)(&value)
}

func (n *NetObligation1) SetAmount(value, currency string) {
	n.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (n *NetObligation1) AddParticipantNettingIdentification() *NettingIdentification1Choice {
	n.ParticipantNettingIdentification = new(NettingIdentification1Choice)
	return n.ParticipantNettingIdentification
}

func (n *NetObligation1) SetObligationDirection(value string) {
	n.ObligationDirection = (*CreditDebit3Code)(&value)
}

func (n *NetObligation1) AddCounterpartyNettingIdentification() *NettingIdentification1Choice {
	n.CounterpartyNettingIdentification = new(NettingIdentification1Choice)
	return n.CounterpartyNettingIdentification
}

func (n *NetObligation1) AddNetServiceCounterpartyIdentification() *PartyIdentification73Choice {
	n.NetServiceCounterpartyIdentification = new(PartyIdentification73Choice)
	return n.NetServiceCounterpartyIdentification
}

func (n *NetObligation1) AddCounterpartySettlementInstructions() *SettlementParties29 {
	n.CounterpartySettlementInstructions = new(SettlementParties29)
	return n.CounterpartySettlementInstructions
}

func (n *NetObligation1) SetTransactionsNumber(value string) {
	n.TransactionsNumber = (*Max10NumericText)(&value)
}
