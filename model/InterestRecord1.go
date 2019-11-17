package model

// Provides transaction specific interest information that applies to the underlying transaction.
type InterestRecord1 struct {

	// Amount of interest included in the entry amount.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Indicates whether the interest amount included in the entry is credit or debit amount.
	// Usage: A zero amount is considered to be a credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies the type of interest.
	Type *InterestType1Choice `xml:"Tp,omitempty"`

	// Set of elements used to qualify the interest rate.
	Rate *Rate3 `xml:"Rate,omitempty"`

	// Range of time between a start date and an end date for the calculation of the interest.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Specifies the reason for the interest.
	Reason *Max35Text `xml:"Rsn,omitempty"`

	// Provides details on the tax applied to charges.
	Tax *TaxCharges2 `xml:"Tax,omitempty"`
}

func (i *InterestRecord1) SetAmount(value, currency string) {
	i.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (i *InterestRecord1) SetCreditDebitIndicator(value string) {
	i.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (i *InterestRecord1) AddType() *InterestType1Choice {
	i.Type = new(InterestType1Choice)
	return i.Type
}

func (i *InterestRecord1) AddRate() *Rate3 {
	i.Rate = new(Rate3)
	return i.Rate
}

func (i *InterestRecord1) AddFromToDate() *DateTimePeriodDetails {
	i.FromToDate = new(DateTimePeriodDetails)
	return i.FromToDate
}

func (i *InterestRecord1) SetReason(value string) {
	i.Reason = (*Max35Text)(&value)
}

func (i *InterestRecord1) AddTax() *TaxCharges2 {
	i.Tax = new(TaxCharges2)
	return i.Tax
}
