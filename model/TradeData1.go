package model

// Provides information on the status of a trade.
type TradeData1 struct {

	// Refers to the identification of a notification.
	NotificationIdentification *Max35Text `xml:"NtfctnId"`

	// Reference to the unique identification assigned to a trade by a central matching system.
	MatchingSystemUniqueReference *Max35Text `xml:"MtchgSysUnqRef"`

	// Identifies the party which assigned a status to a treasury trade.
	StatusOriginator *Max35Text `xml:"StsOrgtr,omitempty"`

	// Specifies the new status of a trade.
	CurrentStatus *TradeStatus1Code `xml:"CurSts"`

	// Description of the status of a trade when no coded form is available.
	ExtendedCurrentStatus *Extended350Code `xml:"XtndedCurSts"`

	// Additional information on the current status of a trade in a central system.
	CurrentStatusSubType *Max70Text `xml:"CurStsSubTp,omitempty"`

	// Specifies the time at which the current status was assigned.
	CurrentStatusTime *ISODateTime `xml:"CurStsTm,omitempty"`

	// Specifies the previous status of a trade.
	PreviousStatus *TradeStatus1Code `xml:"PrvsSts,omitempty"`

	// Description of the status of a trade when no coded form is available.
	ExtendedPreviousStatus *Extended350Code `xml:"XtndedPrvsSts,omitempty"`

	// Additional information on the previous status of a trade in a central system.
	PreviousStatusSubType *Max70Text `xml:"PrvsStsSubTp,omitempty"`

	// Specifies the time at which the previous status was assigned.
	PreviousStatusTime *ISODateTime `xml:"PrvsStsTm,omitempty"`

	// Specifies the product for which the status of the confirmation is reported.
	ProductType *Max4AlphaNumericText `xml:"PdctTp,omitempty"`
}

func (t *TradeData1) SetNotificationIdentification(value string) {
	t.NotificationIdentification = (*Max35Text)(&value)
}

func (t *TradeData1) SetMatchingSystemUniqueReference(value string) {
	t.MatchingSystemUniqueReference = (*Max35Text)(&value)
}

func (t *TradeData1) SetStatusOriginator(value string) {
	t.StatusOriginator = (*Max35Text)(&value)
}

func (t *TradeData1) SetCurrentStatus(value string) {
	t.CurrentStatus = (*TradeStatus1Code)(&value)
}

func (t *TradeData1) SetExtendedCurrentStatus(value string) {
	t.ExtendedCurrentStatus = (*Extended350Code)(&value)
}

func (t *TradeData1) SetCurrentStatusSubType(value string) {
	t.CurrentStatusSubType = (*Max70Text)(&value)
}

func (t *TradeData1) SetCurrentStatusTime(value string) {
	t.CurrentStatusTime = (*ISODateTime)(&value)
}

func (t *TradeData1) SetPreviousStatus(value string) {
	t.PreviousStatus = (*TradeStatus1Code)(&value)
}

func (t *TradeData1) SetExtendedPreviousStatus(value string) {
	t.ExtendedPreviousStatus = (*Extended350Code)(&value)
}

func (t *TradeData1) SetPreviousStatusSubType(value string) {
	t.PreviousStatusSubType = (*Max70Text)(&value)
}

func (t *TradeData1) SetPreviousStatusTime(value string) {
	t.PreviousStatusTime = (*ISODateTime)(&value)
}

func (t *TradeData1) SetProductType(value string) {
	t.ProductType = (*Max4AlphaNumericText)(&value)
}
