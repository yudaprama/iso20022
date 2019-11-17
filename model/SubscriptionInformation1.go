package model

// Execution of a subscription order.
type SubscriptionInformation1 struct {

	// Date the investment plan starts.
	DateOfFirstSubscription *ISODate `xml:"DtOfFrstSbcpt"`

	// Amount subscribed in the current tax year into equity (not including dividends).
	EquityComponent *ActiveCurrencyAndAmount `xml:"EqtyCmpnt,omitempty"`

	// Amount subscribed in the current tax year into cash.
	CashComponent *ActiveCurrencyAndAmount `xml:"CshCmpnt,omitempty"`

	// Total amount subscribed in the current tax year.
	TotalAmountYearToDate *ActiveCurrencyAndAmount `xml:"TtlAmtYrToDt"`
}

func (s *SubscriptionInformation1) SetDateOfFirstSubscription(value string) {
	s.DateOfFirstSubscription = (*ISODate)(&value)
}

func (s *SubscriptionInformation1) SetEquityComponent(value, currency string) {
	s.EquityComponent = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionInformation1) SetCashComponent(value, currency string) {
	s.CashComponent = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionInformation1) SetTotalAmountYearToDate(value, currency string) {
	s.TotalAmountYearToDate = NewActiveCurrencyAndAmount(value, currency)
}
