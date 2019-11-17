package model

// Details of settlement of a transaction.
type SettlementDetails100 struct {

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition19Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod4Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`
}

func (s *SettlementDetails100) AddSettlementTransactionCondition() *SettlementTransactionCondition19Choice {
	newValue := new(SettlementTransactionCondition19Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails100) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails100) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails100) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails100) AddSettlementSystemMethod() *SettlementSystemMethod4Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod4Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails100) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails100) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}
