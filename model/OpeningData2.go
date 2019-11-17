package model

// List of elements which specify the opening of a non deliverable trade.
type OpeningData2 struct {

	// Date at which the trading parties execute a treasury trade.
	TradeDate *ISODate `xml:"TradDt"`

	// Refers to the identification of a notification assigned by the trading side.
	NotificationIdentification *Max35Text `xml:"NtfctnId"`

	// Reference common to the parties of a trade.
	CommonReference *Max35Text `xml:"CmonRef,omitempty"`

	// Refers to the identification of a previous event in the life of a  non deliverable forward trade.
	RelatedReference *Max35Text `xml:"RltdRef,omitempty"`

	// Describes the reason for the cancellation or the amendment.
	AmendOrCancelReason *Max35Text `xml:"AmdOrCclRsn,omitempty"`

	// Specifies the amounts of the non deliverable trade which is reported.
	TradeAmounts *AmountsAndValueDate1 `xml:"TradAmts"`

	// Exchange rate between two currencies. The rate is agreed by the trading parties during the negotiation process.
	AgreedRate *AgreedRate1 `xml:"AgrdRate"`

	// Set of parameters used to calculate the valuation rate to be applied to a non-deliverable agreement.
	ValuationConditions *NonDeliverableForwardValuationConditions2 `xml:"ValtnConds"`
}

func (o *OpeningData2) SetTradeDate(value string) {
	o.TradeDate = (*ISODate)(&value)
}

func (o *OpeningData2) SetNotificationIdentification(value string) {
	o.NotificationIdentification = (*Max35Text)(&value)
}

func (o *OpeningData2) SetCommonReference(value string) {
	o.CommonReference = (*Max35Text)(&value)
}

func (o *OpeningData2) SetRelatedReference(value string) {
	o.RelatedReference = (*Max35Text)(&value)
}

func (o *OpeningData2) SetAmendOrCancelReason(value string) {
	o.AmendOrCancelReason = (*Max35Text)(&value)
}

func (o *OpeningData2) AddTradeAmounts() *AmountsAndValueDate1 {
	o.TradeAmounts = new(AmountsAndValueDate1)
	return o.TradeAmounts
}

func (o *OpeningData2) AddAgreedRate() *AgreedRate1 {
	o.AgreedRate = new(AgreedRate1)
	return o.AgreedRate
}

func (o *OpeningData2) AddValuationConditions() *NonDeliverableForwardValuationConditions2 {
	o.ValuationConditions = new(NonDeliverableForwardValuationConditions2)
	return o.ValuationConditions
}
