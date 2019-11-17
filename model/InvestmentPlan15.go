package model

// Plan that allows investors to schedule periodical investments or divestments, according to pre-defined criteria.
type InvestmentPlan15 struct {

	// Frequency of the investment or divestment.
	Frequency *Frequency20Choice `xml:"Frqcy"`

	// Date the investment plan starts.
	StartDate *ISODate `xml:"StartDt,omitempty"`

	// Date the investment plan stops.
	EndDate *ISODate `xml:"EndDt,omitempty"`

	// Amount of the periodical payments.
	Quantity *UnitsOrAmount1Choice `xml:"Qty"`

	// Indicates whether an ordered amount is a gross amount (including transaction overhead). If it is not a gross amount, the ordered amount is a net amount (amount to be invested or redeemed from the fund to which other elements will be added).
	GrossAmountIndicator *YesNoIndicator `xml:"GrssAmtInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference2Code `xml:"IncmPref,omitempty"`

	// Initial amount or number of initial instalments.
	InitialAmount *InitialAmount1Choice `xml:"InitlAmt,omitempty"`

	// Total number of times the amount must be invested at the predefined frequency as of the start date of the investment plan.
	TotalNumberOfInstalments *Number `xml:"TtlNbOfInstlmts,omitempty"`

	// Indicates the rounding direction when an amount is to be spread over several funds.
	RoundingDirection *RoundingDirection1Code `xml:"RndgDrctn,omitempty"`

	// Security that an investment plan invests in, or from which the investment plan divests.
	SecurityDetails []*Repartition5 `xml:"SctyDtls"`

	// Cash settlement standing instruction associated to the investment plan and to be either inserted or deleted.
	ModifiedCashSettlement []*CashSettlement2 `xml:"ModfdCshSttlm,omitempty"`

	// Reference of the underlying investment contract. In some markets, such as Italy, this might be required to segregate holdings between the same investment account.
	ContractReference *Max35Text `xml:"CtrctRef,omitempty"`

	// Reference of the previous contract to which this savings or withdrawal plan is related.
	RelatedContractReference *Max35Text `xml:"RltdCtrctRef,omitempty"`

	// Identification of the product as designated by the fund manager. In some markets, such as Italy, the financial product or service related to a savings plan or withdrawal plan are identified by a product identification or number.
	ProductIdentification *Max35Text `xml:"PdctId,omitempty"`

	// Reference of the underlying service level agreement (SLA) governing fees.
	SLAChargeAndCommissionReference *Max35Text `xml:"SLAChrgAndComssnRef,omitempty"`

	// Specifies the type of insurance contract to which the savings investment plan is linked.
	InsuranceCover *InsuranceType2Choice `xml:"InsrncCover,omitempty"`

	// Status of the savings or withdrawal investment plan.
	PlanStatus *PlanStatus2Choice `xml:"PlanSts,omitempty"`

	// Role or function of the instalment manager.
	InstalmentManagerRole *PartyRole4Choice `xml:"InstlmtMgrRole,omitempty"`
}

func (i *InvestmentPlan15) AddFrequency() *Frequency20Choice {
	i.Frequency = new(Frequency20Choice)
	return i.Frequency
}

func (i *InvestmentPlan15) SetStartDate(value string) {
	i.StartDate = (*ISODate)(&value)
}

func (i *InvestmentPlan15) SetEndDate(value string) {
	i.EndDate = (*ISODate)(&value)
}

func (i *InvestmentPlan15) AddQuantity() *UnitsOrAmount1Choice {
	i.Quantity = new(UnitsOrAmount1Choice)
	return i.Quantity
}

func (i *InvestmentPlan15) SetGrossAmountIndicator(value string) {
	i.GrossAmountIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentPlan15) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference2Code)(&value)
}

func (i *InvestmentPlan15) AddInitialAmount() *InitialAmount1Choice {
	i.InitialAmount = new(InitialAmount1Choice)
	return i.InitialAmount
}

func (i *InvestmentPlan15) SetTotalNumberOfInstalments(value string) {
	i.TotalNumberOfInstalments = (*Number)(&value)
}

func (i *InvestmentPlan15) SetRoundingDirection(value string) {
	i.RoundingDirection = (*RoundingDirection1Code)(&value)
}

func (i *InvestmentPlan15) AddSecurityDetails() *Repartition5 {
	newValue := new(Repartition5)
	i.SecurityDetails = append(i.SecurityDetails, newValue)
	return newValue
}

func (i *InvestmentPlan15) AddModifiedCashSettlement() *CashSettlement2 {
	newValue := new(CashSettlement2)
	i.ModifiedCashSettlement = append(i.ModifiedCashSettlement, newValue)
	return newValue
}

func (i *InvestmentPlan15) SetContractReference(value string) {
	i.ContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan15) SetRelatedContractReference(value string) {
	i.RelatedContractReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan15) SetProductIdentification(value string) {
	i.ProductIdentification = (*Max35Text)(&value)
}

func (i *InvestmentPlan15) SetSLAChargeAndCommissionReference(value string) {
	i.SLAChargeAndCommissionReference = (*Max35Text)(&value)
}

func (i *InvestmentPlan15) AddInsuranceCover() *InsuranceType2Choice {
	i.InsuranceCover = new(InsuranceType2Choice)
	return i.InsuranceCover
}

func (i *InvestmentPlan15) AddPlanStatus() *PlanStatus2Choice {
	i.PlanStatus = new(PlanStatus2Choice)
	return i.PlanStatus
}

func (i *InvestmentPlan15) AddInstalmentManagerRole() *PartyRole4Choice {
	i.InstalmentManagerRole = new(PartyRole4Choice)
	return i.InstalmentManagerRole
}
