package model

// Provides information on the status of a trade.
type TradeData16 struct {

	// Reference to the unique system identification assigned to the trade by the central matching system.
	MatchingSystemUniqueReference *Max35Text `xml:"MtchgSysUnqRef"`

	// Reference to the unique matching identification assigned to the trade and to the matching trade from the counterparty by the central matching system.
	MatchingSystemMatchingReference *Max35Text `xml:"MtchgSysMtchgRef,omitempty"`

	// Unique reference from the central settlement system that allows the removal of alleged trades once the matched status notification for the matching side has been received.
	MatchingSystemMatchedSideReference *Max35Text `xml:"MtchgSysMtchdSdRef,omitempty"`

	// Party that assigned the status to the trade.
	StatusOriginator *Max35Text `xml:"StsOrgtr,omitempty"`

	// Specifies the new status of a trade.
	CurrentStatus *StatusAndSubStatus2 `xml:"CurSts"`

	// Additional information on the current status of a trade in a central system.
	CurrentStatusSubType *StatusSubType2Code `xml:"CurStsSubTp,omitempty"`

	// Specifies the date and time at which the current status was assigned.
	CurrentStatusDateTime *ISODateTime `xml:"CurStsDtTm,omitempty"`

	// Specifies the previous status of a trade.
	PreviousStatus *Status28Choice `xml:"PrvsSts,omitempty"`

	// Additional information on the previous status of a trade in a central system.
	PreviousStatusSubType *StatusSubType2Code `xml:"PrvsStsSubTp,omitempty"`
}

func (t *TradeData16) SetMatchingSystemUniqueReference(value string) {
	t.MatchingSystemUniqueReference = (*Max35Text)(&value)
}

func (t *TradeData16) SetMatchingSystemMatchingReference(value string) {
	t.MatchingSystemMatchingReference = (*Max35Text)(&value)
}

func (t *TradeData16) SetMatchingSystemMatchedSideReference(value string) {
	t.MatchingSystemMatchedSideReference = (*Max35Text)(&value)
}

func (t *TradeData16) SetStatusOriginator(value string) {
	t.StatusOriginator = (*Max35Text)(&value)
}

func (t *TradeData16) AddCurrentStatus() *StatusAndSubStatus2 {
	t.CurrentStatus = new(StatusAndSubStatus2)
	return t.CurrentStatus
}

func (t *TradeData16) SetCurrentStatusSubType(value string) {
	t.CurrentStatusSubType = (*StatusSubType2Code)(&value)
}

func (t *TradeData16) SetCurrentStatusDateTime(value string) {
	t.CurrentStatusDateTime = (*ISODateTime)(&value)
}

func (t *TradeData16) AddPreviousStatus() *Status28Choice {
	t.PreviousStatus = new(Status28Choice)
	return t.PreviousStatus
}

func (t *TradeData16) SetPreviousStatusSubType(value string) {
	t.PreviousStatusSubType = (*StatusSubType2Code)(&value)
}
