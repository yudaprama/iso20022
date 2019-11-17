package model

// Provides information about the corporate action security option.
type SecuritiesOption35 struct {

	// Identification of the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Specifies whether the value is a debit or credit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd"`

	// Specifies that the security identified  is a temporary security identification used for processing reasons, for example, contra security used in the US.
	TemporaryFinancialInstrumentIndicator *TemporaryFinancialInstrumentIndicator1Choice `xml:"TempFinInstrmInd,omitempty"`

	// Quantity of securities that have been posted (credit or debit) to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType23Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to securities movement.
	DateDetails *SecurityDate6 `xml:"DtDtls"`

	// Provides information about the rates related to securities movement.
	RateDetails *CorporateActionRate29 `xml:"RateDtls,omitempty"`

	// Provides information about the prices related to securities movement.
	PriceDetails *CorporateActionPrice39 `xml:"PricDtls,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties24 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties24 `xml:"DlvrgSttlmPties,omitempty"`
}

func (s *SecuritiesOption35) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesOption35) SetCreditDebitIndicator(value string) {
	s.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (s *SecuritiesOption35) AddTemporaryFinancialInstrumentIndicator() *TemporaryFinancialInstrumentIndicator1Choice {
	s.TemporaryFinancialInstrumentIndicator = new(TemporaryFinancialInstrumentIndicator1Choice)
	return s.TemporaryFinancialInstrumentIndicator
}

func (s *SecuritiesOption35) AddPostingQuantity() *Quantity6Choice {
	s.PostingQuantity = new(Quantity6Choice)
	return s.PostingQuantity
}

func (s *SecuritiesOption35) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesOption35) AddFractionDisposition() *FractionDispositionType23Choice {
	s.FractionDisposition = new(FractionDispositionType23Choice)
	return s.FractionDisposition
}

func (s *SecuritiesOption35) SetCurrencyOption(value string) {
	s.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (s *SecuritiesOption35) AddDateDetails() *SecurityDate6 {
	s.DateDetails = new(SecurityDate6)
	return s.DateDetails
}

func (s *SecuritiesOption35) AddRateDetails() *CorporateActionRate29 {
	s.RateDetails = new(CorporateActionRate29)
	return s.RateDetails
}

func (s *SecuritiesOption35) AddPriceDetails() *CorporateActionPrice39 {
	s.PriceDetails = new(CorporateActionPrice39)
	return s.PriceDetails
}

func (s *SecuritiesOption35) AddReceivingSettlementParties() *SettlementParties24 {
	s.ReceivingSettlementParties = new(SettlementParties24)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesOption35) AddDeliveringSettlementParties() *SettlementParties24 {
	s.DeliveringSettlementParties = new(SettlementParties24)
	return s.DeliveringSettlementParties
}
