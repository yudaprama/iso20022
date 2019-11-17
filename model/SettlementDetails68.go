package model

// Details of settlement of a transaction.
type SettlementDetails68 struct {

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType9Choice `xml:"SctiesTxTp,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition12Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn1Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType10Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the forex standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction1Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity4Choice `xml:"SttlgCpcty,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking1Choice `xml:"Trckg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed1Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification20 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification20 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails68) AddSecuritiesTransactionType() *SecuritiesTransactionType9Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType9Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails68) AddSettlementTransactionCondition() *SettlementTransactionCondition12Choice {
	newValue := new(SettlementTransactionCondition12Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails68) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails68) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails68) AddDeliveryReturnReason() *DeliveryReturn1Choice {
	s.DeliveryReturnReason = new(DeliveryReturn1Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails68) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails68) AddExposureType() *ExposureType10Choice {
	s.ExposureType = new(ExposureType10Choice)
	return s.ExposureType
}

func (s *SettlementDetails68) AddFXStandingInstruction() *FXStandingInstruction1Choice {
	s.FXStandingInstruction = new(FXStandingInstruction1Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails68) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails68) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails68) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails68) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails68) AddSettlingCapacity() *SettlingCapacity4Choice {
	s.SettlingCapacity = new(SettlingCapacity4Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails68) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails68) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails68) AddTracking() *Tracking1Choice {
	s.Tracking = new(Tracking1Choice)
	return s.Tracking
}

func (s *SettlementDetails68) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails68) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails68) AddModificationCancellationAllowed() *ModificationCancellationAllowed1Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed1Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails68) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails68) AddSecuritiesSubBalanceType() *GenericIdentification20 {
	s.SecuritiesSubBalanceType = new(GenericIdentification20)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails68) AddCashSubBalanceType() *GenericIdentification20 {
	s.CashSubBalanceType = new(GenericIdentification20)
	return s.CashSubBalanceType
}
