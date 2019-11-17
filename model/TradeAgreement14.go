package model

// Date and identification of a trade.
type TradeAgreement14 struct {

	// Date on which the trading parties agreed on the trade.
	TradeDate *ISODate `xml:"TradDt"`

	// Reference of the present instruction assigned by the party issuing the message. This reference must be unique amongst all messages of same type sent by the same party.
	OriginatorReference *Max35Text `xml:"OrgtrRef"`

	// Reference common to both parties of the trade.
	CommonReference *Max35Text `xml:"CmonRef,omitempty"`

	// Specifies the type of underlying transaction, for example cancellation (CANC).
	OperationType *Max4Text `xml:"OprTp,omitempty"`

	// Specifies the business role between the submitter and the trade party, for example Agent (AGNT).
	OperationScope *Max4Text `xml:"OprScp,omitempty"`

	// Specifies the product for which the status of the confirmation is reported.
	ProductType *Max35Text `xml:"PdctTp,omitempty"`

	// To indicate the requested CLS settlement session that the related trade is part of.
	SettlementSessionIdentifier *Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// Specifies if the FX transaction is PVP settlement. Payment versus payment (PvP) settlement arrangement allows for two currencies in a foreign exchange (FX) contract to exchange simultaneously on a central settlement platform to eliminate the settlement risk. To apply PvP, the two parties in the FX contract need to have a pre-agreement with the central settlement platform, for example, USD/MYR FX deals require both parties to have an agreement to settle via HK Interbank Clearing Ltd settlement platform.
	PaymentVersusPaymentIndicator *YesNoIndicator `xml:"PmtVrssPmtInd,omitempty"`
}

func (t *TradeAgreement14) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *TradeAgreement14) SetOriginatorReference(value string) {
	t.OriginatorReference = (*Max35Text)(&value)
}

func (t *TradeAgreement14) SetCommonReference(value string) {
	t.CommonReference = (*Max35Text)(&value)
}

func (t *TradeAgreement14) SetOperationType(value string) {
	t.OperationType = (*Max4Text)(&value)
}

func (t *TradeAgreement14) SetOperationScope(value string) {
	t.OperationScope = (*Max4Text)(&value)
}

func (t *TradeAgreement14) SetProductType(value string) {
	t.ProductType = (*Max35Text)(&value)
}

func (t *TradeAgreement14) SetSettlementSessionIdentifier(value string) {
	t.SettlementSessionIdentifier = (*Exact4AlphaNumericText)(&value)
}

func (t *TradeAgreement14) SetPaymentVersusPaymentIndicator(value string) {
	t.PaymentVersusPaymentIndicator = (*YesNoIndicator)(&value)
}
