package model

// Provides information about the CA security option.
type SecurityOption1 struct {

	// Identification of the financial instrument.
	SecurityIdentification *FinancialInstrumentDescription3 `xml:"SctyId"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Quantity of financial instrument.
	SecuritiesQuantity *UnitOrFaceAmount1Choice `xml:"SctiesQty,omitempty"`

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableSecuritiesQuantity *UnitOrFaceAmount1Choice `xml:"MinExrcblSctiesQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleSecuritiesQuantity *UnitOrFaceAmount1Choice `xml:"MinExrcblMltplSctiesQty,omitempty"`

	// New denomination of the financial instrument following, eg, an increase or decrease in nominal value.
	NewDenominationSecuritiesQuantity *UnitOrFaceAmount1Choice `xml:"NewDnmtnSctiesQty,omitempty"`

	// Quantity of equity that makes up the new board lot.
	NewBoardLotSecuritiesQuantity *UnitOrFaceAmount1Choice `xml:"NewBrdLotSctiesQty,omitempty"`

	// Specifies whether the shares are ranking for dividend or pari passu.
	ShareRanking *ShareRanking1FormatChoice `xml:"ShrRnkg,omitempty"`

	// Quantity of additional intermediate securities/new equities awarded for a given quantity of securities derived from subscription.
	AdditionalQuantityForSubscribedResultantSecurities *QuantityToQuantityRatio1 `xml:"AddtlQtyForSbcbdRsltntScties,omitempty"`

	// Provides information about the dates related to a securities movement.
	DateDetails *CorporateActionDate3 `xml:"DtDtls,omitempty"`

	// Provides information about the prices related to a securities movement.
	PriceDetails *CorporateActionPrice4 `xml:"PricDtls,omitempty"`

	// Period during which intermediate securities are tradable in a secondary market.
	TradingPeriod *Period1 `xml:"TradgPrd,omitempty"`

	// Quantity of additional securities for a given quantity of underlying securities where underlying securities are not exchanged or debited, eg, 1 for 1: 1 new equity credited for every 1 underlying equity = 2 resulting equities.
	AdditionalQuantityForExistingSecurities *QuantityToQuantityRatio1 `xml:"AddtlQtyForExstgScties,omitempty"`

	// Specifies that the security is a temporary security.
	TemporaryFinancialInstrumentIndicator *YesNoIndicator `xml:"TempFinInstrmInd,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType1FormatChoice `xml:"FrctnDspstn,omitempty"`
}

func (s *SecurityOption1) AddSecurityIdentification() *FinancialInstrumentDescription3 {
	s.SecurityIdentification = new(FinancialInstrumentDescription3)
	return s.SecurityIdentification
}

func (s *SecurityOption1) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecurityOption1) AddSecuritiesQuantity() *UnitOrFaceAmount1Choice {
	s.SecuritiesQuantity = new(UnitOrFaceAmount1Choice)
	return s.SecuritiesQuantity
}

func (s *SecurityOption1) AddMinimumExercisableSecuritiesQuantity() *UnitOrFaceAmount1Choice {
	s.MinimumExercisableSecuritiesQuantity = new(UnitOrFaceAmount1Choice)
	return s.MinimumExercisableSecuritiesQuantity
}

func (s *SecurityOption1) AddMinimumExercisableMultipleSecuritiesQuantity() *UnitOrFaceAmount1Choice {
	s.MinimumExercisableMultipleSecuritiesQuantity = new(UnitOrFaceAmount1Choice)
	return s.MinimumExercisableMultipleSecuritiesQuantity
}

func (s *SecurityOption1) AddNewDenominationSecuritiesQuantity() *UnitOrFaceAmount1Choice {
	s.NewDenominationSecuritiesQuantity = new(UnitOrFaceAmount1Choice)
	return s.NewDenominationSecuritiesQuantity
}

func (s *SecurityOption1) AddNewBoardLotSecuritiesQuantity() *UnitOrFaceAmount1Choice {
	s.NewBoardLotSecuritiesQuantity = new(UnitOrFaceAmount1Choice)
	return s.NewBoardLotSecuritiesQuantity
}

func (s *SecurityOption1) AddShareRanking() *ShareRanking1FormatChoice {
	s.ShareRanking = new(ShareRanking1FormatChoice)
	return s.ShareRanking
}

func (s *SecurityOption1) AddAdditionalQuantityForSubscribedResultantSecurities() *QuantityToQuantityRatio1 {
	s.AdditionalQuantityForSubscribedResultantSecurities = new(QuantityToQuantityRatio1)
	return s.AdditionalQuantityForSubscribedResultantSecurities
}

func (s *SecurityOption1) AddDateDetails() *CorporateActionDate3 {
	s.DateDetails = new(CorporateActionDate3)
	return s.DateDetails
}

func (s *SecurityOption1) AddPriceDetails() *CorporateActionPrice4 {
	s.PriceDetails = new(CorporateActionPrice4)
	return s.PriceDetails
}

func (s *SecurityOption1) AddTradingPeriod() *Period1 {
	s.TradingPeriod = new(Period1)
	return s.TradingPeriod
}

func (s *SecurityOption1) AddAdditionalQuantityForExistingSecurities() *QuantityToQuantityRatio1 {
	s.AdditionalQuantityForExistingSecurities = new(QuantityToQuantityRatio1)
	return s.AdditionalQuantityForExistingSecurities
}

func (s *SecurityOption1) SetTemporaryFinancialInstrumentIndicator(value string) {
	s.TemporaryFinancialInstrumentIndicator = (*YesNoIndicator)(&value)
}

func (s *SecurityOption1) AddFractionDisposition() *FractionDispositionType1FormatChoice {
	s.FractionDisposition = new(FractionDispositionType1FormatChoice)
	return s.FractionDisposition
}
