package model

// Year in which the ISA plan is issued.
type ISAYearsOfIssue5 struct {

	// ISA that was issued during the current fiscal year.
	CurrentYear *CurrentYearType1Choice `xml:"CurYr,omitempty"`

	// Indicates whether the ISA contains a cash component asset for transfer.
	CashComponentIndicator *YesNoIndicator `xml:"CshCmpntInd"`

	// Selection of investment plans issued during previous years.
	PreviousYears *PreviousYear3 `xml:"PrvsYrs,omitempty"`

	// Specifies the amounts already subscribed for the current year.
	CurrentYearSubscriptionDetails *SubscriptionInformation1 `xml:"CurYrSbcptDtls"`
}

func (i *ISAYearsOfIssue5) AddCurrentYear() *CurrentYearType1Choice {
	i.CurrentYear = new(CurrentYearType1Choice)
	return i.CurrentYear
}

func (i *ISAYearsOfIssue5) SetCashComponentIndicator(value string) {
	i.CashComponentIndicator = (*YesNoIndicator)(&value)
}

func (i *ISAYearsOfIssue5) AddPreviousYears() *PreviousYear3 {
	i.PreviousYears = new(PreviousYear3)
	return i.PreviousYears
}

func (i *ISAYearsOfIssue5) AddCurrentYearSubscriptionDetails() *SubscriptionInformation1 {
	i.CurrentYearSubscriptionDetails = new(SubscriptionInformation1)
	return i.CurrentYearSubscriptionDetails
}
