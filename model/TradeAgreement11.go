package model

// Date and identification of a trade together with references to previous events in its life.
type TradeAgreement11 struct {

	// Date on which the trading parties agreed to amend or cancel the trade.
	TradeDate *ISODate `xml:"TradDt"`

	// Reference of the present instruction assigned by the party issuing the message. This reference must be unique amongst all messages of same type sent by the same party.
	OriginatorReference *Max35Text `xml:"OrgtrRef"`

	// Identification of a matching system reference by a choice between a matching system unique identification or the related reference.
	MatchingSystemReference *MatchingSystemReference1Choice `xml:"MtchgSysRef"`

	// Reference common to both parties of the trade.
	CommonReference *Max35Text `xml:"CmonRef,omitempty"`

	// Describes the reason for the cancellation or the amendment.
	AmendOrCancelReason *Max35Text `xml:"AmdOrCclRsn,omitempty"`

	// Specifies the type of  underlying transaction, for example cancellation (CANC).
	OperationType *Max4Text `xml:"OprTp,omitempty"`

	// Specifies the business role between the submitter and the trade party, for example Agent (AGNT).
	OperationScope *Max4Text `xml:"OprScp,omitempty"`

	// To indicate the requested CLS settlement session that the related trade is part of.
	SettlementSessionIdentifier *Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// Specifies if the FX transaction is PVP settlement. Payment versus payment (PvP) settlement arrangement allows for two currencies in a foreign exchange (FX) contract to exchange simultaneously on a central settlement platform to eliminate the settlement risk. To apply PvP, the two parties in the FX contract need to have a pre-agreement with the central settlement platform, for example, USD/MYR FX deals require both parties to have an agreement to settle via HK Interbank Clearing Ltd settlement platform.
	PaymentVersusPaymentIndicator *YesNoIndicator `xml:"PmtVrssPmtInd,omitempty"`
}

func (t *TradeAgreement11) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *TradeAgreement11) SetOriginatorReference(value string) {
	t.OriginatorReference = (*Max35Text)(&value)
}

func (t *TradeAgreement11) AddMatchingSystemReference() *MatchingSystemReference1Choice {
	t.MatchingSystemReference = new(MatchingSystemReference1Choice)
	return t.MatchingSystemReference
}

func (t *TradeAgreement11) SetCommonReference(value string) {
	t.CommonReference = (*Max35Text)(&value)
}

func (t *TradeAgreement11) SetAmendOrCancelReason(value string) {
	t.AmendOrCancelReason = (*Max35Text)(&value)
}

func (t *TradeAgreement11) SetOperationType(value string) {
	t.OperationType = (*Max4Text)(&value)
}

func (t *TradeAgreement11) SetOperationScope(value string) {
	t.OperationScope = (*Max4Text)(&value)
}

func (t *TradeAgreement11) SetSettlementSessionIdentifier(value string) {
	t.SettlementSessionIdentifier = (*Exact4AlphaNumericText)(&value)
}

func (t *TradeAgreement11) SetPaymentVersusPaymentIndicator(value string) {
	t.PaymentVersusPaymentIndicator = (*YesNoIndicator)(&value)
}
