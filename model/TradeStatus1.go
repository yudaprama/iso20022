package model

// Specifies the status of a trade in a central settlement system.
type TradeStatus1 struct {

	// Specifies whether a trade is alleged or not.
	AllegedTrade *YesNoIndicator `xml:"AllgdTrad,omitempty"`

	// Reference to the unique identification assigned to a trade by a central matching system.
	MatchingSystemUniqueReference *Max35Text `xml:"MtchgSysUnqRef"`

	// Specifies the status of a trade
	Status *TradeStatus1Code `xml:"Sts,omitempty"`

	// Description of the status of a trade when no coded form is available.
	ExtendedStatus *Extended350Code `xml:"XtndedSts,omitempty"`

	// Additional information on the status of a trade in a central system.
	StatusSubType *Max70Text `xml:"StsSubTp,omitempty"`

	// Specifies the time at which a status was assigned.
	StatusTime *ISODateTime `xml:"StsTm,omitempty"`

	// Identifies the party which assigned a status to a treasury trade.
	StatusOriginator *Max35Text `xml:"StsOrgtr,omitempty"`
}

func (t *TradeStatus1) SetAllegedTrade(value string) {
	t.AllegedTrade = (*YesNoIndicator)(&value)
}

func (t *TradeStatus1) SetMatchingSystemUniqueReference(value string) {
	t.MatchingSystemUniqueReference = (*Max35Text)(&value)
}

func (t *TradeStatus1) SetStatus(value string) {
	t.Status = (*TradeStatus1Code)(&value)
}

func (t *TradeStatus1) SetExtendedStatus(value string) {
	t.ExtendedStatus = (*Extended350Code)(&value)
}

func (t *TradeStatus1) SetStatusSubType(value string) {
	t.StatusSubType = (*Max70Text)(&value)
}

func (t *TradeStatus1) SetStatusTime(value string) {
	t.StatusTime = (*ISODateTime)(&value)
}

func (t *TradeStatus1) SetStatusOriginator(value string) {
	t.StatusOriginator = (*Max35Text)(&value)
}
