package model

// Provides information on the status of a trade.
type TradeData10 struct {

	// Identification of the present message assigned by the party issuing the message. This identification must be unique amongst all messages of same type sent by the same party.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Party that assigned the status to the foreign exchange trade.
	StatusOriginator *Max35Text `xml:"StsOrgtr,omitempty"`

	// Specifies the new status of the trade.
	CurrentStatus *StatusAndSubStatus1 `xml:"CurSts"`

	// Additional information about the current status of the trade.
	CurrentStatusSubType *StatusSubType1Code `xml:"CurStsSubTp,omitempty"`

	// Specifies the date and time at which the current status was assigned to all the trades, unless overwritten by a date and time assigned to an individual trade.
	CurrentStatusDateTime *ISODateTime `xml:"CurStsDtTm"`

	// Specifies the previous status of a trade.
	PreviousStatus *Status5Choice `xml:"PrvsSts,omitempty"`

	// Additional information on the previous status of a trade in a central system.
	PreviousStatusSubType *StatusSubType1Code `xml:"PrvsStsSubTp,omitempty"`

	// Specifies the product for which the status of the confirmation is reported, unless overwritten by a product type assigned to an individual trade.
	ProductType *Max35Text `xml:"PdctTp,omitempty"`
}

func (t *TradeData10) SetMessageIdentification(value string) {
	t.MessageIdentification = (*Max35Text)(&value)
}

func (t *TradeData10) SetStatusOriginator(value string) {
	t.StatusOriginator = (*Max35Text)(&value)
}

func (t *TradeData10) AddCurrentStatus() *StatusAndSubStatus1 {
	t.CurrentStatus = new(StatusAndSubStatus1)
	return t.CurrentStatus
}

func (t *TradeData10) SetCurrentStatusSubType(value string) {
	t.CurrentStatusSubType = (*StatusSubType1Code)(&value)
}

func (t *TradeData10) SetCurrentStatusDateTime(value string) {
	t.CurrentStatusDateTime = (*ISODateTime)(&value)
}

func (t *TradeData10) AddPreviousStatus() *Status5Choice {
	t.PreviousStatus = new(Status5Choice)
	return t.PreviousStatus
}

func (t *TradeData10) SetPreviousStatusSubType(value string) {
	t.PreviousStatusSubType = (*StatusSubType1Code)(&value)
}

func (t *TradeData10) SetProductType(value string) {
	t.ProductType = (*Max35Text)(&value)
}
