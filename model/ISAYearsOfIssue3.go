package model

// Year in which the ISA plan is issued.
type ISAYearsOfIssue3 struct {

	// ISA that was issued during the current fiscal year.
	CurrentYearType *ISAType1Code `xml:"CurYrTp,omitempty"`

	// Current year ISA is an ISA that was issued during the current fiscal year.
	ExtendedCurrentYearType *Extended350Code `xml:"XtndedCurYrTp,omitempty"`

	// Indicates whether the ISA contains a cash component asset for transfer.
	CashComponentIndicator *YesNoIndicator `xml:"CshCmpntInd"`

	// Specifies the amounts already subscribed for the current year.
	CurrentYearSubscriptionDetails *SubscriptionInformation1 `xml:"CurYrSbcptDtls"`

	// Selection of investment plans issued during previous years.
	PreviousYears *PreviousYear1 `xml:"PrvsYrs,omitempty"`
}

func (i *ISAYearsOfIssue3) SetCurrentYearType(value string) {
	i.CurrentYearType = (*ISAType1Code)(&value)
}

func (i *ISAYearsOfIssue3) SetExtendedCurrentYearType(value string) {
	i.ExtendedCurrentYearType = (*Extended350Code)(&value)
}

func (i *ISAYearsOfIssue3) SetCashComponentIndicator(value string) {
	i.CashComponentIndicator = (*YesNoIndicator)(&value)
}

func (i *ISAYearsOfIssue3) AddCurrentYearSubscriptionDetails() *SubscriptionInformation1 {
	i.CurrentYearSubscriptionDetails = new(SubscriptionInformation1)
	return i.CurrentYearSubscriptionDetails
}

func (i *ISAYearsOfIssue3) AddPreviousYears() *PreviousYear1 {
	i.PreviousYears = new(PreviousYear1)
	return i.PreviousYears
}
