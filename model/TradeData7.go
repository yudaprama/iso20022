package model

// Provides information on the status of a trade.
type TradeData7 struct {

	// Identification of the present message assigned by the party issuing the message. This identification must be unique amongst all messages of same type sent by the same party.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Represents the original reference of the instruction for which the status is given, as assigned by the participant that submitted the foreign exchange trade.
	OriginatorReference *Max35Text `xml:"OrgtrRef,omitempty"`

	// Reference to the unique system identification assigned to the trade  by the central matching system.
	MatchingSystemUniqueReference *Max35Text `xml:"MtchgSysUnqRef"`

	// Unique matching identification assigned to the trade and to the matching trade from the counterparty by the central matching system.
	MatchingSystemMatchingReference *Max35Text `xml:"MtchgSysMtchgRef,omitempty"`

	// Identification to the unique reference from the central settlement system that allows the removal of alleged trades once the matched status notification for the matching side has been received.
	MatchingSystemMatchedSideReference *Max35Text `xml:"MtchgSysMtchdSdRef,omitempty"`

	// Party that assigned the status to the foreign exchange or derivative trade.
	StatusOriginator *Max20Text `xml:"StsOrgtr,omitempty"`

	// Specifies the new status of a trade.
	CurrentStatus *StatusAndSubStatus1 `xml:"CurSts"`

	// Additional information on the current status of a trade in a central system.
	CurrentStatusSubType *StatusSubType1Code `xml:"CurStsSubTp,omitempty"`

	// Specifies the date and time at which the current status was assigned.
	CurrentStatusDateTime *ISODateTime `xml:"CurStsDtTm,omitempty"`

	// Specifies the previous status of a trade.
	PreviousStatus *Status6Choice `xml:"PrvsSts,omitempty"`

	// Additional information on the previous status of a trade in a central system.
	PreviousStatusSubType *StatusSubType1Code `xml:"PrvsStsSubTp,omitempty"`

	// Specifies the date and time at which the previous status was assigned.
	PreviousStatusDateTime *ISODateTime `xml:"PrvsStsDtTm,omitempty"`

	// Specifies the product for which the status of the confirmation is reported.
	ProductType *Max35Text `xml:"PdctTp,omitempty"`

	// To indicate the requested CLS Settlement Session that the related trade is part of.
	SettlementSessionIdentifier *Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// To indicate if the trade is split.
	SplitTradeIndicator *YesNoIndicator `xml:"SpltTradInd,omitempty"`
}

func (t *TradeData7) SetMessageIdentification(value string) {
	t.MessageIdentification = (*Max35Text)(&value)
}

func (t *TradeData7) SetOriginatorReference(value string) {
	t.OriginatorReference = (*Max35Text)(&value)
}

func (t *TradeData7) SetMatchingSystemUniqueReference(value string) {
	t.MatchingSystemUniqueReference = (*Max35Text)(&value)
}

func (t *TradeData7) SetMatchingSystemMatchingReference(value string) {
	t.MatchingSystemMatchingReference = (*Max35Text)(&value)
}

func (t *TradeData7) SetMatchingSystemMatchedSideReference(value string) {
	t.MatchingSystemMatchedSideReference = (*Max35Text)(&value)
}

func (t *TradeData7) SetStatusOriginator(value string) {
	t.StatusOriginator = (*Max20Text)(&value)
}

func (t *TradeData7) AddCurrentStatus() *StatusAndSubStatus1 {
	t.CurrentStatus = new(StatusAndSubStatus1)
	return t.CurrentStatus
}

func (t *TradeData7) SetCurrentStatusSubType(value string) {
	t.CurrentStatusSubType = (*StatusSubType1Code)(&value)
}

func (t *TradeData7) SetCurrentStatusDateTime(value string) {
	t.CurrentStatusDateTime = (*ISODateTime)(&value)
}

func (t *TradeData7) AddPreviousStatus() *Status6Choice {
	t.PreviousStatus = new(Status6Choice)
	return t.PreviousStatus
}

func (t *TradeData7) SetPreviousStatusSubType(value string) {
	t.PreviousStatusSubType = (*StatusSubType1Code)(&value)
}

func (t *TradeData7) SetPreviousStatusDateTime(value string) {
	t.PreviousStatusDateTime = (*ISODateTime)(&value)
}

func (t *TradeData7) SetProductType(value string) {
	t.ProductType = (*Max35Text)(&value)
}

func (t *TradeData7) SetSettlementSessionIdentifier(value string) {
	t.SettlementSessionIdentifier = (*Exact4AlphaNumericText)(&value)
}

func (t *TradeData7) SetSplitTradeIndicator(value string) {
	t.SplitTradeIndicator = (*YesNoIndicator)(&value)
}
