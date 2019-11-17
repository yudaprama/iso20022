package model

// Date and identification of a trade together with references to previous events in its life.
type TradeAgreement12 struct {

	// Date on which the trading parties agreed on the trade.
	TradeDate *ISODate `xml:"TradDt"`

	// Identification of the present message assigned by the party issuing the message. This identification must be unique amongst all messages of same type sent by the same party.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Represents the original reference of the instruction for which the status is given, as assigned by the participant that submitted the foreign exchange trade.
	OriginatorReference *Max35Text `xml:"OrgtrRef"`

	// Reference common to both parties of the trade.
	CommonReference *Max35Text `xml:"CmonRef,omitempty"`

	// Specifies the reason for the cancellation or the amendment.
	AmendOrCancelReason *Max35Text `xml:"AmdOrCclRsn,omitempty"`

	// Reference to the identification of a previous event in the life of a trade which is amended or cancelled.
	RelatedReference *Max35Text `xml:"RltdRef,omitempty"`

	// Specifies the product for which the status of the confirmation is reported.
	ProductType *Max35Text `xml:"PdctTp,omitempty"`

	// Specifies the type of underlying transaction, for example cancellation (CANC).
	OperationType *Max4Text `xml:"OprTp,omitempty"`

	// Specifies the business role between the submitter and the trade party, for example, agent (AGNT).
	OperationScope *Max4Text `xml:"OprScp,omitempty"`

	// To indicate the requested CLS settlement session that the related trade is part of.
	SettlementSessionIdentifier *Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// To indicate if the trade is split.
	SplitTradeIndicator *YesNoIndicator `xml:"SpltTradInd"`

	// Specifies if the FX transaction is PVP settlement. Payment versus payment (PvP) settlement arrangement allows for two currencies in a foreign exchange (FX) contract to exchange simultaneously on a central settlement platform to eliminate the settlement risk. To apply PvP, the two parties in the FX contract need to have a pre-agreement with the central settlement platform, for example, USD/MYR FX deals require both parties to have an agreement to settle via HK Interbank Clearing Ltd settlement platform.
	PaymentVersusPaymentIndicator *YesNoIndicator `xml:"PmtVrssPmtInd,omitempty"`
}

func (t *TradeAgreement12) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *TradeAgreement12) SetMessageIdentification(value string) {
	t.MessageIdentification = (*Max35Text)(&value)
}

func (t *TradeAgreement12) SetOriginatorReference(value string) {
	t.OriginatorReference = (*Max35Text)(&value)
}

func (t *TradeAgreement12) SetCommonReference(value string) {
	t.CommonReference = (*Max35Text)(&value)
}

func (t *TradeAgreement12) SetAmendOrCancelReason(value string) {
	t.AmendOrCancelReason = (*Max35Text)(&value)
}

func (t *TradeAgreement12) SetRelatedReference(value string) {
	t.RelatedReference = (*Max35Text)(&value)
}

func (t *TradeAgreement12) SetProductType(value string) {
	t.ProductType = (*Max35Text)(&value)
}

func (t *TradeAgreement12) SetOperationType(value string) {
	t.OperationType = (*Max4Text)(&value)
}

func (t *TradeAgreement12) SetOperationScope(value string) {
	t.OperationScope = (*Max4Text)(&value)
}

func (t *TradeAgreement12) SetSettlementSessionIdentifier(value string) {
	t.SettlementSessionIdentifier = (*Exact4AlphaNumericText)(&value)
}

func (t *TradeAgreement12) SetSplitTradeIndicator(value string) {
	t.SplitTradeIndicator = (*YesNoIndicator)(&value)
}

func (t *TradeAgreement12) SetPaymentVersusPaymentIndicator(value string) {
	t.PaymentVersusPaymentIndicator = (*YesNoIndicator)(&value)
}
