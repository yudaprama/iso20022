package model

// Instruction from an investor to sell investment fund units back to the fund.
type RedemptionOrder3 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Specifies the category of the investment fund order.
	OrderType []*FundOrderType1 `xml:"OrdrTp,omitempty"`

	// Account impacted by an investment fund order.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls"`

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson2 `xml:"BnfcryDtls,omitempty"`

	// Quantity of investment fund units redeemed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Net amount of money used to derive the quantity of investment fund units to be sold.
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt"`

	// Portion of the investor's holdings, in a specific investment fund/ fund class, that is redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Amount of money used to derive the quantity of investment fund units to be sold, including all charges, commissions, and tax.
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms5 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

	// Charge for the placement of an order.
	ChargeDetails []*Charge8 `xml:"ChrgDtls,omitempty"`

	// Commission linked to the execution of an investment fund order.
	CommissionDetails []*Commission6 `xml:"ComssnDtls,omitempty"`

	// Tax applicable to an investment fund order.
	TaxDetails []*Tax6 `xml:"TaxDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters3 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Payment transaction resulting from the investment fund order execution.
	CashSettlementDetails *PaymentTransaction18 `xml:"CshSttlmDtls,omitempty"`
}

func (r *RedemptionOrder3) SetOrderReference(value string) {
	r.OrderReference = (*Max35Text)(&value)
}

func (r *RedemptionOrder3) AddOrderType() *FundOrderType1 {
	newValue := new(FundOrderType1)
	r.OrderType = append(r.OrderType, newValue)
	return newValue
}

func (r *RedemptionOrder3) AddInvestmentAccountDetails() *InvestmentAccount13 {
	r.InvestmentAccountDetails = new(InvestmentAccount13)
	return r.InvestmentAccountDetails
}

func (r *RedemptionOrder3) AddBeneficiaryDetails() *IndividualPerson2 {
	r.BeneficiaryDetails = new(IndividualPerson2)
	return r.BeneficiaryDetails
}

func (r *RedemptionOrder3) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	r.UnitsNumber = new(FinancialInstrumentQuantity1)
	return r.UnitsNumber
}

func (r *RedemptionOrder3) SetNetAmount(value, currency string) {
	r.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder3) SetHoldingsRedemptionRate(value string) {
	r.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (r *RedemptionOrder3) SetRounding(value string) {
	r.Rounding = (*RoundingDirection2Code)(&value)
}

func (r *RedemptionOrder3) SetGrossAmount(value, currency string) {
	r.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RedemptionOrder3) AddForeignExchangeDetails() *ForeignExchangeTerms5 {
	r.ForeignExchangeDetails = new(ForeignExchangeTerms5)
	return r.ForeignExchangeDetails
}

func (r *RedemptionOrder3) SetIncomePreference(value string) {
	r.IncomePreference = (*IncomePreference1Code)(&value)
}

func (r *RedemptionOrder3) SetGroup1Or2Units(value string) {
	r.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (r *RedemptionOrder3) AddChargeDetails() *Charge8 {
	newValue := new(Charge8)
	r.ChargeDetails = append(r.ChargeDetails, newValue)
	return newValue
}

func (r *RedemptionOrder3) AddCommissionDetails() *Commission6 {
	newValue := new(Commission6)
	r.CommissionDetails = append(r.CommissionDetails, newValue)
	return newValue
}

func (r *RedemptionOrder3) AddTaxDetails() *Tax6 {
	newValue := new(Tax6)
	r.TaxDetails = append(r.TaxDetails, newValue)
	return newValue
}

func (r *RedemptionOrder3) AddSettlementAndCustodyDetails() *FundSettlementParameters3 {
	r.SettlementAndCustodyDetails = new(FundSettlementParameters3)
	return r.SettlementAndCustodyDetails
}

func (r *RedemptionOrder3) SetPhysicalDeliveryIndicator(value string) {
	r.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (r *RedemptionOrder3) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	r.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return r.PhysicalDeliveryDetails
}

func (r *RedemptionOrder3) AddCashSettlementDetails() *PaymentTransaction18 {
	r.CashSettlementDetails = new(PaymentTransaction18)
	return r.CashSettlementDetails
}
