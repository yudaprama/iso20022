package model

// Provides information about the corporate action security option.
type SecuritiesOption50 struct {

	// Identification of the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator3Choice `xml:"TempFinInstrmInd,omitempty"`

	// Indicates whether the securities are newly issued or not.
	NewSecuritiesIssuanceIndicator *NewSecuritiesIssuanceType6Code `xml:"NewSctiesIssncInd,omitempty"`

	// Proceeds are taxable according to the information provided by the issuer / offeror.
	IssuerOfferorTaxabilityIndicator *GenericIdentification30 `xml:"IssrOfferrTaxbltyInd,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification30 `xml:"IncmTp,omitempty"`

	// Specifies the basis for the reduced rate of withholding.
	ExemptionType []*GenericIdentification30 `xml:"XmptnTp,omitempty"`

	// Indicates the country from which the income originates.
	CountryOfIncomeSource *CountryCode `xml:"CtryOfIncmSrc,omitempty"`

	// Quantity of securities that have been posted (credit or debit) to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat10Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType27Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate11 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate72 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice59 `xml:"PricDtls,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties42 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties42 `xml:"DlvrgSttlmPties,omitempty"`
}

func (s *SecuritiesOption50) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesOption50) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption50) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator3Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator3Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption50) SetNewSecuritiesIssuanceIndicator(value string) {
	s.NewSecuritiesIssuanceIndicator = (*NewSecuritiesIssuanceType6Code)(&value)
}

func (s *SecuritiesOption50) AddIssuerOfferorTaxabilityIndicator() *GenericIdentification30 {
	s.IssuerOfferorTaxabilityIndicator = new(GenericIdentification30)
	return s.IssuerOfferorTaxabilityIndicator
}

func (s *SecuritiesOption50) AddIncomeType() *GenericIdentification30 {
	s.IncomeType = new(GenericIdentification30)
	return s.IncomeType
}

func (s *SecuritiesOption50) AddExemptionType() *GenericIdentification30 {
	newValue := new(GenericIdentification30)
	s.ExemptionType = append(s.ExemptionType, newValue)
	return newValue
}

func (s *SecuritiesOption50) SetCountryOfIncomeSource(value string) {
	s.CountryOfIncomeSource = (*CountryCode)(&value)
}

func (s *SecuritiesOption50) AddPostingQuantity() *Quantity6Choice {
	s.PostingQuantity = new(Quantity6Choice)
	return s.PostingQuantity
}

func (s *SecuritiesOption50) AddSafekeepingPlace() *SafekeepingPlaceFormat10Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat10Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption50) AddFractionDisposition() *FractionDispositionType27Choice {
	s.FractionDisposition = new(FractionDispositionType27Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption50) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption50) AddDateDetails() *SecurityDate11 {
	s.DateDetails = new(SecurityDate11)
	return s.DateDetails
}

func (s *SecuritiesOption50) AddRateDetails() *CorporateActionRate72 {
	s.RateDetails = new(CorporateActionRate72)
	return s.RateDetails
}

func (s *SecuritiesOption50) AddPriceDetails() *CorporateActionPrice59 {
	s.PriceDetails = new(CorporateActionPrice59)
	return s.PriceDetails
}

func (s *SecuritiesOption50) AddReceivingSettlementParties() *SettlementParties42 {
	s.ReceivingSettlementParties = new(SettlementParties42)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesOption50) AddDeliveringSettlementParties() *SettlementParties42 {
	s.DeliveringSettlementParties = new(SettlementParties42)
	return s.DeliveringSettlementParties
}
