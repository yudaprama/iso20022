package model

// Provides general interest information that applies to the account at a particular moment in time.
type AccountInterest1 struct {

	// Specifies the type of interest.
	Type *InterestType1Choice `xml:"Tp,omitempty"`

	// Set of elements qualifying the interest rate.
	Rate []*Rate1 `xml:"Rate,omitempty"`

	// Range of time between a start date and an end date for the calculation of the interest.
	FromToDate *DateTimePeriodDetails `xml:"FrToDt,omitempty"`

	// Underlying reason for the interest, eg, yearly credit interest on a savings account.
	Reason *Max35Text `xml:"Rsn,omitempty"`
}

func (a *AccountInterest1) AddType() *InterestType1Choice {
	a.Type = new(InterestType1Choice)
	return a.Type
}

func (a *AccountInterest1) AddRate() *Rate1 {
	newValue := new(Rate1)
	a.Rate = append(a.Rate, newValue)
	return newValue
}

func (a *AccountInterest1) AddFromToDate() *DateTimePeriodDetails {
	a.FromToDate = new(DateTimePeriodDetails)
	return a.FromToDate
}

func (a *AccountInterest1) SetReason(value string) {
	a.Reason = (*Max35Text)(&value)
}
