package model

// Provides the clearing details.
type Clearing4 struct {

	// Indicates to the clearing member whether the trade is eligible for settlement netting or not.
	SettlementNettingEligibleCode *NettingEligible1Code `xml:"SttlmNetgElgblCd"`

	// Clearing organisation that will clear the trade.
	// Note: This field allows clearing member firm to segregate flows coming from clearing counterparty's clearing system. Indeed, clearing member firms receive messages from the same system (same sender) and this field allows them to know if the message is related to equities or derivatives.
	ClearingSegment *PartyIdentification35Choice `xml:"ClrSgmt,omitempty"`

	// Indicates if the position is guaranteed or non-guaranteed by the central counterparty, that is whether the CCP has done the novation and then guarantees the trade, or not.
	GuaranteedTrade *YesNoIndicator `xml:"GrntedTrad,omitempty"`

	// In case of trades that are not guaranteed by the central counterparty (this is when the central counterparty has not done the novation), provides details such as the trade counterparty member identification or the trade counterparty clearing member identification.
	NonGuaranteedTrade *NonGuaranteedTrade3 `xml:"NonGrntedTrad,omitempty"`
}

func (c *Clearing4) SetSettlementNettingEligibleCode(value string) {
	c.SettlementNettingEligibleCode = (*NettingEligible1Code)(&value)
}

func (c *Clearing4) AddClearingSegment() *PartyIdentification35Choice {
	c.ClearingSegment = new(PartyIdentification35Choice)
	return c.ClearingSegment
}

func (c *Clearing4) SetGuaranteedTrade(value string) {
	c.GuaranteedTrade = (*YesNoIndicator)(&value)
}

func (c *Clearing4) AddNonGuaranteedTrade() *NonGuaranteedTrade3 {
	c.NonGuaranteedTrade = new(NonGuaranteedTrade3)
	return c.NonGuaranteedTrade
}
