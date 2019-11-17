package model

// Order to invest the investor's principal in an investment fund.
type SubscriptionBulkOrder2 struct {

	// Market in which the advised trade transaction was executed.
	PlaceOfTrade *CountryCode `xml:"PlcOfTrad,omitempty"`

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Date on which the order expires.
	ExpiryDateTime *ISODateTime `xml:"XpryDtTm,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	CancellationRight *CancellationRight1 `xml:"CxlRght,omitempty"`

	// Investment fund class related to an order.
	FinancialInstrumentDetails *FinancialInstrument6 `xml:"FinInstrmDtls"`

	// Order to invest the investor's principal in an investment fund.
	IndividualOrderDetails []*SubscriptionOrder3 `xml:"IndvOrdrDtls"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *CurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *CurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Payment transaction resulting from the investment fund order execution.
	BulkCashSettlementDetails *PaymentTransaction17 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (s *SubscriptionBulkOrder2) SetPlaceOfTrade(value string) {
	s.PlaceOfTrade = (*CountryCode)(&value)
}

func (s *SubscriptionBulkOrder2) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SubscriptionBulkOrder2) SetExpiryDateTime(value string) {
	s.ExpiryDateTime = (*ISODateTime)(&value)
}

func (s *SubscriptionBulkOrder2) AddCancellationRight() *CancellationRight1 {
	s.CancellationRight = new(CancellationRight1)
	return s.CancellationRight
}

func (s *SubscriptionBulkOrder2) AddFinancialInstrumentDetails() *FinancialInstrument6 {
	s.FinancialInstrumentDetails = new(FinancialInstrument6)
	return s.FinancialInstrumentDetails
}

func (s *SubscriptionBulkOrder2) AddIndividualOrderDetails() *SubscriptionOrder3 {
	newValue := new(SubscriptionOrder3)
	s.IndividualOrderDetails = append(s.IndividualOrderDetails, newValue)
	return newValue
}

func (s *SubscriptionBulkOrder2) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*CurrencyCode)(&value)
}

func (s *SubscriptionBulkOrder2) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*CurrencyCode)(&value)
}

func (s *SubscriptionBulkOrder2) AddBulkCashSettlementDetails() *PaymentTransaction17 {
	s.BulkCashSettlementDetails = new(PaymentTransaction17)
	return s.BulkCashSettlementDetails
}
