package model

// Provides the trade leg statement details.
type TradeLegStatement3 struct {

	// Identifies the clearing member account at the Central counterparty through which the trade must be cleared (sometimes called position account).
	ClearingAccount *SecuritiesAccount18 `xml:"ClrAcct,omitempty"`

	// Clearing organisation that will clear the trade.
	// Note: This field allows Clearing Member Firm to segregate flows coming from clearing counterparty's clearing system. Indeed, Clearing Member Firms receive messages from the same system (same sender) and this field allows them to know if the message is related to equities or derivatives.
	ClearingSegment *PartyIdentification35Choice `xml:"ClrSgmt,omitempty"`

	// Provides the identification for the non-clearing member.
	NonClearingMember *PartyIdentificationAndAccount31 `xml:"NonClrMmb,omitempty"`

	// Provides the lists of all trades during the period in consideration for the statement.
	TradeLegsDetails []*TradeLeg9 `xml:"TradLegsDtls"`
}

func (t *TradeLegStatement3) AddClearingAccount() *SecuritiesAccount18 {
	t.ClearingAccount = new(SecuritiesAccount18)
	return t.ClearingAccount
}

func (t *TradeLegStatement3) AddClearingSegment() *PartyIdentification35Choice {
	t.ClearingSegment = new(PartyIdentification35Choice)
	return t.ClearingSegment
}

func (t *TradeLegStatement3) AddNonClearingMember() *PartyIdentificationAndAccount31 {
	t.NonClearingMember = new(PartyIdentificationAndAccount31)
	return t.NonClearingMember
}

func (t *TradeLegStatement3) AddTradeLegsDetails() *TradeLeg9 {
	newValue := new(TradeLeg9)
	t.TradeLegsDetails = append(t.TradeLegsDetails, newValue)
	return newValue
}
