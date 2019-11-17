package model

// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
type InvestmentPlan4 struct {

	// Frequency of the investment or divestment.
	Frequency *EventFrequency1Code `xml:"Frqcy"`

	// Frequency of the investment or divestment.
	ExtendedFrequency *Extended350Code `xml:"XtndedFrqcy"`

	// Date the investment plan starts.
	StartDate *ISODate `xml:"StartDt"`

	// Date the investment plan stops.
	EndDate *ISODate `xml:"EndDt,omitempty"`

	// Currency and amount of the periodical payments.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an ordered amount is a gross amount (including all charges, commissions, tax). If it is not a gross amount, the ordered amount is a net amount (amount to be invested or redeemed from the fund to which other elements will be added).
	GrossAmountIndicator *YesNoIndicator `xml:"GrssAmtInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Number of pre-paid instalment periods at the time the investment plan is created.
	InitialNumberOfInstalment *Number `xml:"InitlNbOfInstlmt,omitempty"`

	// Total number of times the amount must be invested at the predefined frequency as of the start date of the investment plan.
	TotalNumberOfInstalment *Number `xml:"TtlNbOfInstlmt,omitempty"`

	// Indicates the rounding direction when an amount is to be spread over several funds.
	RoundingDirection *RoundingDirection1Code `xml:"RndgDrctn,omitempty"`

	// Security that an investment plan invests in, or from which the investment plan divests.
	SecurityDetails []*Repartition1 `xml:"SctyDtls"`

	// Cash settlement standing instruction associated to the investment plan.
	CashSettlement *InvestmentFundCashSettlementInformation3 `xml:"CshSttlm,omitempty"`
}

func (i *InvestmentPlan4) SetFrequency(value string) {
	i.Frequency = (*EventFrequency1Code)(&value)
}

func (i *InvestmentPlan4) SetExtendedFrequency(value string) {
	i.ExtendedFrequency = (*Extended350Code)(&value)
}

func (i *InvestmentPlan4) SetStartDate(value string) {
	i.StartDate = (*ISODate)(&value)
}

func (i *InvestmentPlan4) SetEndDate(value string) {
	i.EndDate = (*ISODate)(&value)
}

func (i *InvestmentPlan4) SetAmount(value, currency string) {
	i.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (i *InvestmentPlan4) SetGrossAmountIndicator(value string) {
	i.GrossAmountIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentPlan4) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentPlan4) SetInitialNumberOfInstalment(value string) {
	i.InitialNumberOfInstalment = (*Number)(&value)
}

func (i *InvestmentPlan4) SetTotalNumberOfInstalment(value string) {
	i.TotalNumberOfInstalment = (*Number)(&value)
}

func (i *InvestmentPlan4) SetRoundingDirection(value string) {
	i.RoundingDirection = (*RoundingDirection1Code)(&value)
}

func (i *InvestmentPlan4) AddSecurityDetails() *Repartition1 {
	newValue := new(Repartition1)
	i.SecurityDetails = append(i.SecurityDetails, newValue)
	return newValue
}

func (i *InvestmentPlan4) AddCashSettlement() *InvestmentFundCashSettlementInformation3 {
	i.CashSettlement = new(InvestmentFundCashSettlementInformation3)
	return i.CashSettlement
}
