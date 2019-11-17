package model

// Provides details on the interest statement.
type InterestStatement3 struct {

	// Provides the period during which the interest rate has been applied.
	InterestPeriod *DatePeriodDetails `xml:"IntrstPrd"`

	// Provides the total amount of interest that is due to partyA.
	TotalInterestAmountDueToA *ActiveCurrencyAndAmount `xml:"TtlIntrstAmtDueToA,omitempty"`

	// Provides the total amount of interest that is due to partyB.
	TotalInterestAmountDueToB *ActiveCurrencyAndAmount `xml:"TtlIntrstAmtDueToB,omitempty"`

	// Indicates the value date of the interest statement.
	ValueDate *ISODate `xml:"ValDt"`

	// Provides the reference to the interest payment request.
	InterestPaymentRequestIdentification *Max35Text `xml:"IntrstPmtReqId,omitempty"`

	// Provides the details of the interest calculation.
	InterestCalculation []*InterestCalculation3 `xml:"IntrstClctn,omitempty"`
}

func (i *InterestStatement3) AddInterestPeriod() *DatePeriodDetails {
	i.InterestPeriod = new(DatePeriodDetails)
	return i.InterestPeriod
}

func (i *InterestStatement3) SetTotalInterestAmountDueToA(value, currency string) {
	i.TotalInterestAmountDueToA = NewActiveCurrencyAndAmount(value, currency)
}

func (i *InterestStatement3) SetTotalInterestAmountDueToB(value, currency string) {
	i.TotalInterestAmountDueToB = NewActiveCurrencyAndAmount(value, currency)
}

func (i *InterestStatement3) SetValueDate(value string) {
	i.ValueDate = (*ISODate)(&value)
}

func (i *InterestStatement3) SetInterestPaymentRequestIdentification(value string) {
	i.InterestPaymentRequestIdentification = (*Max35Text)(&value)
}

func (i *InterestStatement3) AddInterestCalculation() *InterestCalculation3 {
	newValue := new(InterestCalculation3)
	i.InterestCalculation = append(i.InterestCalculation, newValue)
	return newValue
}
