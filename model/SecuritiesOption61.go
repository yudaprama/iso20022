package model

// Provides information about the corporate action security option.
type SecuritiesOption61 struct {

	// Provides description of the financial instrument related to securities movement.
	SecurityDetails *FinancialInstrumentAttributes83 `xml:"SctyDtls"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator4Choice `xml:"TempFinInstrmInd,omitempty"`

	// Specifies information regarding outturn resources that cannot be processed by the Central Securities Depository (CSD). Special delivery instruction is required from the account owner for the corporate action outcome to be credited.
	NonEligibleProceedsIndicator *NonEligibleProceedsIndicator4Choice `xml:"NonElgblPrcdsInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *IssuerOfferorTaxabilityIndicator1Choice `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Indicates whether the securities are newly issued or not.
	NewSecuritiesIssuanceIndicator *NewSecuritiesIssuanceType5Code `xml:"NewSctiesIssncInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification47 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification47 `xml:"XmptnTp,omitempty"`

	// Quantity of securities based on the terms of the corporate action event and balance of underlying securities entitled to the account owner. (This quantity can be positive or negative).
	EntitledQuantity *Quantity10Choice `xml:"EntitldQty,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat17Choice `xml:"SfkpgPlc,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType31Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period3Choice `xml:"TradgPrd,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate14 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate77 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice66 `xml:"PricDtls,omitempty"`
}

func (s *SecuritiesOption61) AddSecurityDetails() *FinancialInstrumentAttributes83 {
	s.SecurityDetails = new(FinancialInstrumentAttributes83)
	return s.SecurityDetails
}

func (s *SecuritiesOption61) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption61) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator4Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator4Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption61) AddNonEligibleProceedsIndicator() *NonEligibleProceedsIndicator4Choice {
	s.NonEligibleProceedsIndicator = new(NonEligibleProceedsIndicator4Choice)
	return s.NonEligibleProceedsIndicator
}

func (s *SecuritiesOption61) AddIssuerOfferorTaxabilityIndicator() *IssuerOfferorTaxabilityIndicator1Choice {
	s.IssuerOfferorTaxabilityIndicator = new(IssuerOfferorTaxabilityIndicator1Choice)
	return s.IssuerOfferorTaxabilityIndicator
}

func (s *SecuritiesOption61) SetNewSecuritiesIssuanceIndicator(value string) {
	s.NewSecuritiesIssuanceIndicator = (*NewSecuritiesIssuanceType5Code)(&value)
}

func (s *SecuritiesOption61) AddIncomeType() *GenericIdentification47 {
	s.IncomeType = new(GenericIdentification47)
	return s.IncomeType
}

func (s *SecuritiesOption61) AddExemptionType() *GenericIdentification47 {
	newValue := new(GenericIdentification47)
	s.ExemptionType = append(s.ExemptionType, newValue)
	return newValue
}

func (s *SecuritiesOption61) AddEntitledQuantity() *Quantity10Choice {
	s.EntitledQuantity = new(Quantity10Choice)
	return s.EntitledQuantity
}

func (s *SecuritiesOption61) AddSafekeepingPlace() *SafekeepingPlaceFormat17Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat17Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption61) SetCountryOfIncomeSource(value string) {
	s.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (s *SecuritiesOption61) AddFractionDisposition() *FractionDispositionType31Choice {
	s.FractionDisposition = new(FractionDispositionType31Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption61) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption61) AddTradingPeriod() *Period3Choice {
	s.TradingPeriod = new(Period3Choice)
	return s.TradingPeriod
}

func (s *SecuritiesOption61) AddDateDetails() *SecurityDate14 {
	s.DateDetails = new(SecurityDate14)
	return s.DateDetails
}

func (s *SecuritiesOption61) AddRateDetails() *CorporateActionRate77 {
	s.RateDetails = new(CorporateActionRate77)
	return s.RateDetails
}

func (s *SecuritiesOption61) AddPriceDetails() *CorporateActionPrice66 {
	s.PriceDetails = new(CorporateActionPrice66)
	return s.PriceDetails
}
