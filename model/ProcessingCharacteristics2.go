package model

// Processing characteristics linked to the instrument, ie, not to  the market.
type ProcessingCharacteristics2 struct {

	// Currency in which a subscription or redemption is accepted.
	DealingCurrencyAccepted []*ActiveCurrencyCode `xml:"DealgCcyAccptd"`

	// Minimum initial quantity of securities, expressed as an amount that must be purchased at subscription.
	InitialInvestment *Forms `xml:"InitlInvstmt"`

	// Minimum quantity of securities, expressed as number of units/shares that must be purchased.
	SubsequentInvestment *Forms `xml:"SbsqntInvstmt"`

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

	// Enter the last business day following the day on which a subscription order is priced (T) by which settlement will be due
	// for orders placed with the main Fund Order Desk, eg. T+3. Enter "P" (pre-payment) if cleared funds may be required before a
	// subscription order can be executed.
	SettlementCycle *TimeFrame5Choice `xml:"SttlmCycl"`
}

func (p *ProcessingCharacteristics2) AddDealingCurrencyAccepted(value string) {
	p.DealingCurrencyAccepted = append(p.DealingCurrencyAccepted, (*ActiveCurrencyCode)(&value))
}

func (p *ProcessingCharacteristics2) AddInitialInvestment() *Forms {
	p.InitialInvestment = new(Forms)
	return p.InitialInvestment
}

func (p *ProcessingCharacteristics2) AddSubsequentInvestment() *Forms {
	p.SubsequentInvestment = new(Forms)
	return p.SubsequentInvestment
}

func (p *ProcessingCharacteristics2) SetAmountIndicator(value string) {
	p.AmountIndicator = (*YesNoIndicator)(&value)
}

func (p *ProcessingCharacteristics2) SetUnitsIndicator(value string) {
	p.UnitsIndicator = (*YesNoIndicator)(&value)
}

func (p *ProcessingCharacteristics2) AddMainFundOrderDeskLocation() *MainFundOrderDeskLocation1 {
	p.MainFundOrderDeskLocation = new(MainFundOrderDeskLocation1)
	return p.MainFundOrderDeskLocation
}

func (p *ProcessingCharacteristics2) SetDealingCutOffTime(value string) {
	p.DealingCutOffTime = (*ISOTime)(&value)
}

func (p *ProcessingCharacteristics2) AddDealingCutOffTimeFrame() *TimeFrame3 {
	p.DealingCutOffTimeFrame = new(TimeFrame3)
	return p.DealingCutOffTimeFrame
}

func (p *ProcessingCharacteristics2) SetDealingFrequency(value string) {
	p.DealingFrequency = (*EventFrequency5Code)(&value)
}

func (p *ProcessingCharacteristics2) SetDealingFrequencyDescription(value string) {
	p.DealingFrequencyDescription = (*Max350Text)(&value)
}

func (p *ProcessingCharacteristics2) SetLimitedPeriod(value string) {
	p.LimitedPeriod = (*Max350Text)(&value)
}

func (p *ProcessingCharacteristics2) AddSettlementCycle() *TimeFrame5Choice {
	p.SettlementCycle = new(TimeFrame5Choice)
	return p.SettlementCycle
}
