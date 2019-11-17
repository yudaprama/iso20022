package model

// Provides further details on the interest that applies to the account at a particular moment in time.
type AccountInterest3 struct {

	// Specifies the type of interest.
	Type *InterestType1Choice `xml:"Tp,omitempty"`

	// Set of elements used to qualify the interest rate.
	Rate []*Rate3 `xml:"Rate,omitempty"`

	// Range of time between a start date and an end date for the calculation of the interest.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Specifies the reason for the interest.
	Reason *Max35Text `xml:"Rsn,omitempty"`

	// Provides details on the tax applied to charges.
	Tax *TaxCharges2 `xml:"Tax,omitempty"`
}

func (a *AccountInterest3) AddType() *InterestType1Choice {
	a.Type = new(InterestType1Choice)
	return a.Type
}

func (a *AccountInterest3) AddRate() *Rate3 {
	newValue := new(Rate3)
	a.Rate = append(a.Rate, newValue)
	return newValue
}

func (a *AccountInterest3) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountInterest3) SetReason(value string) {
	a.Reason = (*Max35Text)(&value)
}

func (a *AccountInterest3) AddTax() *TaxCharges2 {
	a.Tax = new(TaxCharges2)
	return a.Tax
}
