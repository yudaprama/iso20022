package model

// TimeFrame elements that define a period as number of days before or after a activity.
type TimeFrame3 struct {

	// Specifies a description of any other TimeFrame that may be used for the DealingCutOffTimeFrame
	OtherTimeFrameDescription *Max350Text `xml:"OthrTmFrameDesc,omitempty"`

	// An agreed number of days before the Trade date (T) used to define standard timeframes e.g. T-1 Dealing cut off or T-2 prepayment condition
	//
	// Where = T is the date that the price is applied to a transaction,
	TradeMinus *Number `xml:"TMns,omitempty"`

	// Convention used for adjusting a date when it is not a business day.
	NonWorkingDayAdjustment *BusinessDayConvention1Code `xml:"NonWorkgDayAdjstmnt,omitempty"`

	// Refer to Order Desk
	ReferToOrderDesk *ReferToFundOrderDesk1Code `xml:"RefrToOrdrDsk,omitempty"`
}

func (t *TimeFrame3) SetOtherTimeFrameDescription(value string) {
	t.OtherTimeFrameDescription = (*Max350Text)(&value)
}

func (t *TimeFrame3) SetTradeMinus(value string) {
	t.TradeMinus = (*Number)(&value)
}

func (t *TimeFrame3) SetNonWorkingDayAdjustment(value string) {
	t.NonWorkingDayAdjustment = (*BusinessDayConvention1Code)(&value)
}

func (t *TimeFrame3) SetReferToOrderDesk(value string) {
	t.ReferToOrderDesk = (*ReferToFundOrderDesk1Code)(&value)
}
