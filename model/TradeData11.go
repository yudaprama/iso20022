package model

// Provides information and details on the status of a trade.
type TradeData11 struct {

	// Represents the original reference of the instruction for which the status is given, as assigned by the participant that submitted the foreign exchange trade.
	OriginatorReference *Max35Text `xml:"OrgtrRef,omitempty"`

	// Reference to the unique system identification assigned to the trade  by the central matching system.
	MatchingSystemUniqueReference *Max35Text `xml:"MtchgSysUnqRef"`

	// Reference to the unique matching identification assigned to the trade and to the matching trade from the counterparty by the central matching system.
	MatchingSystemMatchingReference *Max35Text `xml:"MtchgSysMtchgRef,omitempty"`

	// Unique reference from the central settlement system that allows the removal of alleged trades once the matched status notification for the matching side has been received.
	MatchingSystemMatchedSideReference *Max35Text `xml:"MtchgSysMtchdSdRef,omitempty"`

	// The current settlement date of the notification.
	CurrentSettlementDate *ISODate `xml:"CurSttlmDt,omitempty"`

	// Settlement date has been amended.
	NewSettlementDate *ISODate `xml:"NewSttlmDt,omitempty"`

	// Specifies the date and time at which the current status was assigned to the individual trade.
	CurrentStatusDateTime *ISODateTime `xml:"CurStsDtTm,omitempty"`

	// Product type of the individual trade.
	ProductType *Max35Text `xml:"PdctTp,omitempty"`

	// To indicate the requested CLS settlement session that the related trade is part of.
	SettlementSessionIdentifier *Exact4AlphaNumericText `xml:"SttlmSsnIdr,omitempty"`

	// Information that is to be provided to trade repositories in the context of the regulatory standards around over-the-counter (OTC) derivatives, central counterparties and trade repositories.
	RegulatoryReporting *RegulatoryReporting4 `xml:"RgltryRptg,omitempty"`
}

func (t *TradeData11) SetOriginatorReference(value string) {
	t.OriginatorReference = (*Max35Text)(&value)
}

func (t *TradeData11) SetMatchingSystemUniqueReference(value string) {
	t.MatchingSystemUniqueReference = (*Max35Text)(&value)
}

func (t *TradeData11) SetMatchingSystemMatchingReference(value string) {
	t.MatchingSystemMatchingReference = (*Max35Text)(&value)
}

func (t *TradeData11) SetMatchingSystemMatchedSideReference(value string) {
	t.MatchingSystemMatchedSideReference = (*Max35Text)(&value)
}

func (t *TradeData11) SetCurrentSettlementDate(value string) {
	t.CurrentSettlementDate = (*ISODate)(&value)
}

func (t *TradeData11) SetNewSettlementDate(value string) {
	t.NewSettlementDate = (*ISODate)(&value)
}

func (t *TradeData11) SetCurrentStatusDateTime(value string) {
	t.CurrentStatusDateTime = (*ISODateTime)(&value)
}

func (t *TradeData11) SetProductType(value string) {
	t.ProductType = (*Max35Text)(&value)
}

func (t *TradeData11) SetSettlementSessionIdentifier(value string) {
	t.SettlementSessionIdentifier = (*Exact4AlphaNumericText)(&value)
}

func (t *TradeData11) AddRegulatoryReporting() *RegulatoryReporting4 {
	t.RegulatoryReporting = new(RegulatoryReporting4)
	return t.RegulatoryReporting
}
