package model

// Redemption leg, or switch-out, of a switch transaction.
type SwitchRedemptionLegOrder2 struct {

	// Unique technical identifier for an instance of a leg within a switch.
	LegIdentification *Max35Text `xml:"LegId,omitempty"`

	// Investment fund class related to an order.
	FinancialInstrumentDetails *FinancialInstrument6 `xml:"FinInstrmDtls"`

	// Investment fund class related to an order.
	FinancialInstrumentQuantityChoice *FinancialInstrumentQuantity3Choice `xml:"FinInstrmQtyChc,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Tax group to which the purchased investment fund units belong. The investor indicates to the intermediary operating pooled nominees, which type of unit is to be sold.
	Group1Or2Units *UKTaxGroupUnitCode `xml:"Grp1Or2Units,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *CurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *CurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Amount of money associated with a service.
	ChargeDetails []*Charge8 `xml:"ChrgDtls,omitempty"`

	// Commission linked to the execution of an investment fund order.
	CommissionDetails []*Commission6 `xml:"ComssnDtls,omitempty"`

	// Tax related to an investment fund order.
	TaxDetails []*Tax6 `xml:"TaxDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters3 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`
}

func (s *SwitchRedemptionLegOrder2) SetLegIdentification(value string) {
	s.LegIdentification = (*Max35Text)(&value)
}

func (s *SwitchRedemptionLegOrder2) AddFinancialInstrumentDetails() *FinancialInstrument6 {
	s.FinancialInstrumentDetails = new(FinancialInstrument6)
	return s.FinancialInstrumentDetails
}

func (s *SwitchRedemptionLegOrder2) AddFinancialInstrumentQuantityChoice() *FinancialInstrumentQuantity3Choice {
	s.FinancialInstrumentQuantityChoice = new(FinancialInstrumentQuantity3Choice)
	return s.FinancialInstrumentQuantityChoice
}

func (s *SwitchRedemptionLegOrder2) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SwitchRedemptionLegOrder2) SetGroup1Or2Units(value string) {
	s.Group1Or2Units = (*UKTaxGroupUnitCode)(&value)
}

func (s *SwitchRedemptionLegOrder2) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*CurrencyCode)(&value)
}

func (s *SwitchRedemptionLegOrder2) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*CurrencyCode)(&value)
}

func (s *SwitchRedemptionLegOrder2) AddChargeDetails() *Charge8 {
	newValue := new(Charge8)
	s.ChargeDetails = append(s.ChargeDetails, newValue)
	return newValue
}

func (s *SwitchRedemptionLegOrder2) AddCommissionDetails() *Commission6 {
	newValue := new(Commission6)
	s.CommissionDetails = append(s.CommissionDetails, newValue)
	return newValue
}

func (s *SwitchRedemptionLegOrder2) AddTaxDetails() *Tax6 {
	newValue := new(Tax6)
	s.TaxDetails = append(s.TaxDetails, newValue)
	return newValue
}

func (s *SwitchRedemptionLegOrder2) AddSettlementAndCustodyDetails() *FundSettlementParameters3 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters3)
	return s.SettlementAndCustodyDetails
}

func (s *SwitchRedemptionLegOrder2) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchRedemptionLegOrder2) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	s.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return s.PhysicalDeliveryDetails
}
