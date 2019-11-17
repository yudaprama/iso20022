package model

// Provides information about the corporate action security option.
type SecuritiesOption59 struct {

	// Provides description of the financial instrument related to securities movement.
	SecurityDetails *FinancialInstrumentAttributes80 `xml:"SctyDtls"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator3Choice `xml:"TempFinInstrmInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator3Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerOfferorTaxabilityIndicator1Choice `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Indicates whether the securities are newly issued or not.
	NewSecuritiesIssuanceIndicator *NewSecuritiesIssuanceType5Code `xml:"NewSctiesIssncInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification30 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification30 `xml:"XmptnTp,omitempty"`

	// Quantity of securities based on the terms of the corporate action event and balance of underlying securities entitled to the account owner. (This quantity can be positive or negative).
	EntitledQuantity *Quantity6Choice `xml:"EntitldQty,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat10Choice `xml:"SfkpgPlc,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType26Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period3Choice `xml:"TradgPrd,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate12 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate69 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice56 `xml:"PricDtls,omitempty"`
}

func (s *SecuritiesOption59) AddSecurityDetails() *FinancialInstrumentAttributes80 {
	s.SecurityDetails = new(FinancialInstrumentAttributes80)
	return s.SecurityDetails
}

func (s *SecuritiesOption59) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption59) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator3Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator3Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption59) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator3Choice {
	s.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator3Choice)
	return s.NonEligibleProceedsIndicator
}

func (s *SecuritiesOption59) AddIssuerOfferorTaxabilityIndicator() *IssuerOfferorTaxabilityIndicator1Choice {
	s.IssuerOfferorTaxabilityIndicator = new(IssuerOfferorTaxabilityIndicator1Choice)
	return s.IssuerOfferorTaxabilityIndicator
}

func (s *SecuritiesOption59) SetNewSecuritiesIssuanceIndicator(value string) {
	s.NewSecuritiesIssuanceIndicator = (*NewSecuritiesIssuanceType5Code)(&value)
}

func (s *SecuritiesOption59) AddIncomeType() *GenericIdentification30 {
	s.IncomeType = new(GenericIdentification30)
	return s.IncomeType
}

func (s *SecuritiesOption59) AddExemptionType() *GenericIdentification30 {
	newValue := new(GenericIdentification30)
	s.ExemptionType = append(s.ExemptionType, newValue)
	return newValue
}

func (s *SecuritiesOption59) AddEntitledQuantity() *Quantity6Choice {
	s.EntitledQuantity = new(Quantity6Choice)
	return s.EntitledQuantity
}

func (s *SecuritiesOption59) AddSafekeepingPlace() *SafekeepingPlaceFormat10Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat10Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption59) SetCountryOfIncomeSource(value string) {
	s.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (s *SecuritiesOption59) AddFractionDisposition() *FractionDispositionType26Choice {
	s.FractionDisposition = new(FractionDispositionType26Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption59) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption59) AddTradingPeriod() *Period3Choice {
	s.TradingPeriod = new(Period3Choice)
	return s.TradingPeriod
}

func (s *SecuritiesOption59) AddDateDetails() *SecurityDate12 {
	s.DateDetails = new(SecurityDate12)
	return s.DateDetails
}

func (s *SecuritiesOption59) AddRateDetails() *CorporateActionRate69 {
	s.RateDetails = new(CorporateActionRate69)
	return s.RateDetails
}

func (s *SecuritiesOption59) AddPriceDetails() *CorporateActionPrice56 {
	s.PriceDetails = new(CorporateActionPrice56)
	return s.PriceDetails
}
