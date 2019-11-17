package model

// Provides information on the status of a trade.
type TradeData9 struct {

	// Reference to the unique system identification assigned to the trade by the central matching system.
	MatchingSystemUniqueReference *Max35Text `xml:"MtchgSysUnqRef"`

	// Reference to the unique matching identification assigned to the trade and to the matching trade from the counterparty by the central matching system.
	MatchingSystemMatchingReference *Max35Text `xml:"MtchgSysMtchgRef,omitempty"`

	// Unique reference from the central settlement system that allows the removal of alleged trades once the matched status notification for the matching side has been received.
	MatchingSystemMatchedSideReference *Max35Text `xml:"MtchgSysMtchdSdRef,omitempty"`

	// Party that assigned the status to the trade.
	StatusOriginator *Max35Text `xml:"StsOrgtr,omitempty"`

	// Specifies the new status of the trade.
	CurrentStatus *StatusAndSubStatus1 `xml:"CurSts"`

	// Additional information about the current status of the trade.
	CurrentStatusSubType *StatusSubType1Code `xml:"CurStsSubTp,omitempty"`

	// Specifies the date and time at which the current status was assigned.
	CurrentStatusDateTime *ISODateTime `xml:"CurStsDtTm,omitempty"`

	// Specifies the previous status of the trade.
	PreviousStatus *Status5Choice `xml:"PrvsSts,omitempty"`

	// Specifies whether a trade is alleged or not.
	AllegedTrade *YesNoIndicator `xml:"AllgdTrad,omitempty"`

	// Additional information on the previous status of a trade in a central system.
	PreviousStatusSubType *StatusSubType1Code `xml:"PrvsStsSubTp,omitempty"`
}

func (t *TradeData9) SetMatchingSystemUniqueReference(value string) {
	t.MatchingSystemUniqueReference = (*Max35Text)(&value)
}

func (t *TradeData9) SetMatchingSystemMatchingReference(value string) {
	t.MatchingSystemMatchingReference = (*Max35Text)(&value)
}

func (t *TradeData9) SetMatchingSystemMatchedSideReference(value string) {
	t.MatchingSystemMatchedSideReference = (*Max35Text)(&value)
}

func (t *TradeData9) SetStatusOriginator(value string) {
	t.StatusOriginator = (*Max35Text)(&value)
}

func (t *TradeData9) AddCurrentStatus() *StatusAndSubStatus1 {
	t.CurrentStatus = new(StatusAndSubStatus1)
	return t.CurrentStatus
}

func (t *TradeData9) SetCurrentStatusSubType(value string) {
	t.CurrentStatusSubType = (*StatusSubType1Code)(&value)
}

func (t *TradeData9) SetCurrentStatusDateTime(value string) {
	t.CurrentStatusDateTime = (*ISODateTime)(&value)
}

func (t *TradeData9) AddPreviousStatus() *Status5Choice {
	t.PreviousStatus = new(Status5Choice)
	return t.PreviousStatus
}

func (t *TradeData9) SetAllegedTrade(value string) {
	t.AllegedTrade = (*YesNoIndicator)(&value)
}

func (t *TradeData9) SetPreviousStatusSubType(value string) {
	t.PreviousStatusSubType = (*StatusSubType1Code)(&value)
}
