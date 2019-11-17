package model

// Details of settlement of a transaction.
type SettlementDetails125 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType33Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition17Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType13Choice `xml:"RpTp,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`
}

func (s *SettlementDetails125) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails125) AddSecuritiesTransactionType() *SecuritiesTransactionType33Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType33Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails125) AddSettlementTransactionCondition() *SettlementTransactionCondition17Choice {
	newValue := new(SettlementTransactionCondition17Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails125) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails125) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails125) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails125) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails125) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails125) AddRepurchaseType() *RepurchaseType13Choice {
	s.RepurchaseType = new(RepurchaseType13Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails125) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails125) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}
