package model

// Provides information about the corporate action security option.
type SecuritiesOption18 struct {

	// Identification of the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator1Choice `xml:"TempFinInstrmInd,omitempty"`

	// Quantity of securities that have been posted (credit or debit) to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType4Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate6 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate21 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice9 `xml:"PricDtls,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties15 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties15 `xml:"DlvrgSttlmPties,omitempty"`
}

func (s *SecuritiesOption18) AddSecurityIdentification() *SecurityIdentification14 {
	s.SecurityIdentification = new(SecurityIdentification14)
	return s.SecurityIdentification
}

func (s *SecuritiesOption18) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption18) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator1Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator1Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption18) AddPostingQuantity() *Quantity6Choice {
	s.PostingQuantity = new(Quantity6Choice)
	return s.PostingQuantity
}

func (s *SecuritiesOption18) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption18) AddFractionDisposition() *FractionDispositionType4Choice {
	s.FractionDisposition = new(FractionDispositionType4Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption18) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption18) AddDateDetails() *SecurityDate6 {
	s.DateDetails = new(SecurityDate6)
	return s.DateDetails
}

func (s *SecuritiesOption18) AddRateDetails() *CorporateActionRate21 {
	s.RateDetails = new(CorporateActionRate21)
	return s.RateDetails
}

func (s *SecuritiesOption18) AddPriceDetails() *CorporateActionPrice9 {
	s.PriceDetails = new(CorporateActionPrice9)
	return s.PriceDetails
}

func (s *SecuritiesOption18) AddReceivingSettlementParties() *SettlementParties15 {
	s.ReceivingSettlementParties = new(SettlementParties15)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesOption18) AddDeliveringSettlementParties() *SettlementParties15 {
	s.DeliveringSettlementParties = new(SettlementParties15)
	return s.DeliveringSettlementParties
}
