package model

// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
type InvestmentPlan9 struct {

	// Frequency of the investment or divestment.
	Frequency *Frequency20Choice `xml:"Frqcy"`

	// Date the investment plan starts.
	StartDate *ISODate `xml:"StartDt,omitempty"`

	// Date the investment plan stops.
	EndDate *ISODate `xml:"EndDt,omitempty"`

	// Amount of the periodical payments.
	Quantity *UnitsOrAmount1Choice `xml:"Qty"`

	// Indicates whether an ordered amount is a gross amount (including all charges, commissions, tax). If it is not a gross amount, the ordered amount is a net amount (amount to be invested or redeemed from the fund to which other elements will be added).
	GrossAmountIndicator *YesNoIndicator `xml:"GrssAmtInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Number of pre-paid instalment periods at the time the investment plan is created.
	InitialNumberOfInstalments *Number `xml:"InitlNbOfInstlmts,omitempty"`

	// Total number of times the amount must be invested at the predefined frequency as of the start date of the investment plan.
	TotalNumberOfInstalments *Number `xml:"TtlNbOfInstlmts,omitempty"`

	// Indicates the rounding direction when an amount is to be spread over several funds.
	RoundingDirection *RoundingDirection1Code `xml:"RndgDrctn,omitempty"`

	// Security that an investment plan invests in, or from which the investment plan divests.
	SecurityDetails []*Repartition2 `xml:"SctyDtls"`

	// Cash settlement standing instruction associated to the investment plan.
	CashSettlement []*InvestmentFundCashSettlementInformation7 `xml:"CshSttlm,omitempty"`

	// Reference of the underlying investment contract. In some markets, such as Italy, this might be required to segregate holdings between the same investment account.
	ContractReference *Max35Text `xml:"CtrctRef,omitempty"`

	// Reference of the previous contract to which this savings or withdrawal plan is related.
	RelatedContractReference *Max35Text `xml:"RltdCtrctRef,omitempty"`

	// Identification of the product as designated by the fund manager. In some markets, such as Italy, the financial product or service related to a savings plan or withdrawal plan are identified by a product identification or number.
	ProductIdentification *Max35Text `xml:"PdctId,omitempty"`

	// Reference of the underlying service level agreement (SLA) governing charges and commission.
	SLAChargeAndCommissionReference *Max35Text `xml:"SLAChrgAndComssnRef,omitempty"`

	// Specifies the type of insurance contract to which the savings investment plan is linked.
	InsuranceCover *InsuranceType1Choice `xml:"InsrncCover,omitempty"`
}

func (i *InvestmentPlan9) AddFrequency() *Frequency20Choice {
	i.Frequency = new(Frequency20Choice)
	return i.Frequency
}

func (i *InvestmentPlan9) SetStartDate(value string) {
	i.StartDate = (*ISODate)(&value)
}

func (i *InvestmentPlan9) SetEndDate(value string) {
	i.EndDate = (*ISODate)(&value)
}

func (i *InvestmentPlan9) AddQuantity() *UnitsOrAmount1Choice {
	i.Quantity = new(UnitsOrAmount1Choice)
	return i.Quantity
}

func (i *InvestmentPlan9) SetGrossAmountIndicator(value string) {
	i.GrossAmountIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentPlan9) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentPlan9) SetInitialNumberOfInstalments(value string) {
	i.InitialNumberOfInstalments = (*Number)(&value)
}

func (i *InvestmentPlan9) SetTotalNumberOfInstalments(value string) {
	i.TotalNumberOfInstalments = (*Number)(&value)
}

func (i *InvestmentPlan9) SetRoundingDirection(value string) {
	i.RoundingDirection = (*RoundingDirection1Code)(&value)
}

func (i *InvestmentPlan9) AddSecurityDetails() *Repartition2 {
	newValue := new(Repartition2)
	i.SecurityDetails = append(i.SecurityDetails, newValue)
	return newValue
}

func (i *InvestmentPlan9) AddCashSettlement() *InvestmentFundCashSettlementInformation7 {
	newValue := new(InvestmentFundCashSettlementInformation7)
	i.CashSettlement = append(i.CashSettlement, newValue)
	return newValue
}

func (i *InvestmentPlan9) SetContractReference(value string) {
	i.ContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan9) SetRelatedContractReference(value string) {
	i.RelatedContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan9) SetProductIdentification(value string) {
	i.ProductIdentification = (*Max35Text)(&value)
}

func (i *InvestmentPlan9) SetSLAChargeAndCommissionReference(value string) {
	i.SLAChargeAndCommissionReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan9) AddInsuranceCover() *InsuranceType1Choice {
	i.InsuranceCover = new(InsuranceType1Choice)
	return i.InsuranceCover
}
