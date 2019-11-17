package model

// Processing characteristics linked to the instrument, ie, not to  the market.
type ProcessingCharacteristics3 struct {

	// Currency in which a subscription or redemption is accepted.
	DealingCurrencyAccepted []*ActiveCurrencyCode `xml:"DealgCcyAccptd"`

	// Authorization to claim redemption proceeds.
	RedemptionAuthorisation *Forms `xml:"RedAuthstn"`

	// Indicates whether a subscription or  a redemption can be instructed by amount.
	AmountIndicator *YesNoIndicator `xml:"AmtInd"`

	// Indicates whether subscriptions or redemptions may be placed as a number of units.
	UnitsIndicator *YesNoIndicator `xml:"UnitsInd"`

	// Specifies the location of the main fund order desk.
	MainFundOrderDeskLocation *MainFundOrderDeskLocation1 `xml:"MainFndOrdrDskLctn"`

	// Last date/time at which an order to subscribe or redeem can be given.
	DealingCutOffTime *ISOTime `xml:"DealgCutOffTm"`

	// TimeFrame or period concept that allows definition of a period as number of days before or after a defined activity.
	DealingCutOffTimeFrame *TimeFrame3 `xml:"DealgCutOffTmFrame"`

	// Frequency at which the subscriptions are done.
	DealingFrequency *EventFrequency5Code `xml:"DealgFrqcy"`

	// Description of frequency at which the subscription is done.
	DealingFrequencyDescription *Max350Text `xml:"DealgFrqcyDesc"`

	// Specific period, eg, for some guaranteed funds, during which the units/shares may be redeemed.
	LimitedPeriod *Max350Text `xml:"LtdPrd,omitempty"`

	// Indicate the last business day following the day on which a redemption order is priced (T) by which settlement will be due
	// for orders placed with the main Fund Order Desk.&nbsp; Alternatively, if proceeds will be paid following receipt of written
	// renunciation, indicate the last business day following receipt of the relevant renunciation documentation by the main Fund
	// Order Desk (R) by which the proceeds will be sent.&nbsp; Examples of the above would be T+3, R+4 etc.
	SettlementCycle *TimeFrame4Choice `xml:"SttlmCycl"`
}

func (p *ProcessingCharacteristics3) AddDealingCurrencyAccepted(value string) {
	p.DealingCurrencyAccepted = append(p.DealingCurrencyAccepted, (*ActiveCurrencyCode)(&value))
}

func (p *ProcessingCharacteristics3) AddRedemptionAuthorisation() *Forms {
	p.RedemptionAuthorisation = new(Forms)
	return p.RedemptionAuthorisation
}

func (p *ProcessingCharacteristics3) SetAmountIndicator(value string) {
	p.AmountIndicator = (*YesNoIndicator)(&value)
}

func (p *ProcessingCharacteristics3) SetUnitsIndicator(value string) {
	p.UnitsIndicator = (*YesNoIndicator)(&value)
}

func (p *ProcessingCharacteristics3) AddMainFundOrderDeskLocation() *MainFundOrderDeskLocation1 {
	p.MainFundOrderDeskLocation = new(MainFundOrderDeskLocation1)
	return p.MainFundOrderDeskLocation
}

func (p *ProcessingCharacteristics3) SetDealingCutOffTime(value string) {
	p.DealingCutOffTime = (*ISOTime)(&value)
}

func (p *ProcessingCharacteristics3) AddDealingCutOffTimeFrame() *TimeFrame3 {
	p.DealingCutOffTimeFrame = new(TimeFrame3)
	return p.DealingCutOffTimeFrame
}

func (p *ProcessingCharacteristics3) SetDealingFrequency(value string) {
	p.DealingFrequency = (*EventFrequency5Code)(&value)
}

func (p *ProcessingCharacteristics3) SetDealingFrequencyDescription(value string) {
	p.DealingFrequencyDescription = (*Max350Text)(&value)
}

func (p *ProcessingCharacteristics3) SetLimitedPeriod(value string) {
	p.LimitedPeriod = (*Max350Text)(&value)
}

func (p *ProcessingCharacteristics3) AddSettlementCycle() *TimeFrame4Choice {
	p.SettlementCycle = new(TimeFrame4Choice)
	return p.SettlementCycle
}
