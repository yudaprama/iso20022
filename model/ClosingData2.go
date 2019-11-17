package model

// Set of data specified for the fixing of a non deliverable trade.
type ClosingData2 struct {

	// Date at which the trading parties have agreed on a valuation rate for a non deliverable trade.
	TradeDate *ISODate `xml:"TradDt"`

	// Refers to the identification of a trade assigned by the trading side of a non deliverable forward trade.
	NotificationIdentification *Max35Text `xml:"NtfctnId"`

	// Reference common to the parties of a trade.
	CommonReference *Max35Text `xml:"CmonRef,omitempty"`

	// Refers to the identification of a previous event in the life of a  non deliverable forward trade.
	RelatedReference *Max35Text `xml:"RltdRef,omitempty"`

	// Describes the reason for the cancellation or the amendment.
	AmendOrCancelReason *Max35Text `xml:"AmdOrCclRsn,omitempty"`

	// Specifies the amounts traded at the valuation of a non-deliverable trade.
	TradeAmounts *AmountsAndValueDate1 `xml:"TradAmts"`

	// Rate obtained at valuation time by following the valuation conditions (agreed upon by the trading parties at the opening of the non-deliverable contract).
	ValuationRate *AgreedRate1 `xml:"ValtnRate"`

	// Set of parameters used to calculate the valuation rate to be applied to a non-deliverable agreement.
	ValuationInformation *ValuationData2 `xml:"ValtnInf"`
}

func (c *ClosingData2) SetTradeDate(value string) {
	c.TradeDate = (*ISODate)(&value)
}

func (c *ClosingData2) SetNotificationIdentification(value string) {
	c.NotificationIdentification = (*Max35Text)(&value)
}

func (c *ClosingData2) SetCommonReference(value string) {
	c.CommonReference = (*Max35Text)(&value)
}

func (c *ClosingData2) SetRelatedReference(value string) {
	c.RelatedReference = (*Max35Text)(&value)
}

func (c *ClosingData2) SetAmendOrCancelReason(value string) {
	c.AmendOrCancelReason = (*Max35Text)(&value)
}

func (c *ClosingData2) AddTradeAmounts() *AmountsAndValueDate1 {
	c.TradeAmounts = new(AmountsAndValueDate1)
	return c.TradeAmounts
}

func (c *ClosingData2) AddValuationRate() *AgreedRate1 {
	c.ValuationRate = new(AgreedRate1)
	return c.ValuationRate
}

func (c *ClosingData2) AddValuationInformation() *ValuationData2 {
	c.ValuationInformation = new(ValuationData2)
	return c.ValuationInformation
}
