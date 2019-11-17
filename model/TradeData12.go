package model

// Provides information on the status of a trade.
type TradeData12 struct {

	// Identification of the present message assigned by the party issuing the message. This identification must be unique amongst all messages of same type sent by the same party.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Party that assigned the status to the foreign exchange trade.
	StatusOriginator *Max35Text `xml:"StsOrgtr,omitempty"`

	// Specifies the new status of the trade.
	CurrentStatus *StatusAndSubStatus2 `xml:"CurSts"`

	// Additional information about the current status of the trade.
	CurrentStatusSubType *StatusSubType2Code `xml:"CurStsSubTp,omitempty"`

	// Specifies the date and time at which the current status was assigned to all the trades, unless overwritten by a date and time assigned to an individual trade.
	CurrentStatusDateTime *ISODateTime `xml:"CurStsDtTm"`

	// Specifies the previous status of a trade.
	PreviousStatus *Status28Choice `xml:"PrvsSts,omitempty"`

	// Additional information on the previous status of a trade in a central system.
	PreviousStatusSubType *StatusSubType2Code `xml:"PrvsStsSubTp,omitempty"`

	// Specifies the product for which the status of the confirmation is reported, unless overwritten by a product type assigned to an individual trade.
	ProductType *Max35Text `xml:"PdctTp,omitempty"`

	// To indicate the requested CLS settlement session that the related trade is part of.
	SettlementSessionIdentifier *Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// The identification that links the quoted trades with a submitted Report issued by a central system.
	LinkedReportIdentification *Max35Text `xml:"LkdRptId,omitempty"`
}

func (t *TradeData12) SetMessageIdentification(value string) {
	t.MessageIdentification = (*Max35Text)(&value)
}

func (t *TradeData12) SetStatusOriginator(value string) {
	t.StatusOriginator = (*Max35Text)(&value)
}

func (t *TradeData12) AddCurrentStatus() *StatusAndSubStatus2 {
	t.CurrentStatus = new(StatusAndSubStatus2)
	return t.CurrentStatus
}

func (t *TradeData12) SetCurrentStatusSubType(value string) {
	t.CurrentStatusSubType = (*StatusSubType2Code)(&value)
}

func (t *TradeData12) SetCurrentStatusDateTime(value string) {
	t.CurrentStatusDateTime = (*ISODateTime)(&value)
}

func (t *TradeData12) AddPreviousStatus() *Status28Choice {
	t.PreviousStatus = new(Status28Choice)
	return t.PreviousStatus
}

func (t *TradeData12) SetPreviousStatusSubType(value string) {
	t.PreviousStatusSubType = (*StatusSubType2Code)(&value)
}

func (t *TradeData12) SetProductType(value string) {
	t.ProductType = (*Max35Text)(&value)
}

func (t *TradeData12) SetSettlementSessionIdentifier(value string) {
	t.SettlementSessionIdentifier = (*Exact4AlphaNumericText)(&value)
}

func (t *TradeData12) SetLinkedReportIdentification(value string) {
	t.LinkedReportIdentification = (*Max35Text)(&value)
}
