package model

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails34 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *RestrictedFINXMax16Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *RestrictedFINXMax16Text `xml:"ClsgLegId,omitempty"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate5Choice `xml:"TermntnDt,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType67Choice `xml:"RateTp,omitempty"`

	// Legal framework of the transaction.
	LegalFramework *LegalFramework4Choice `xml:"LglFrmwk,omitempty"`

	// Specifies whether the maturity date of the securities financing transaction may be modified.
	MaturityDateModification *YesNoIndicator `xml:"MtrtyDtMod,omitempty"`

	// Specifies whether the interest is to be paid to the collateral taker. If set to no, the interest is paid to the collateral giver.
	InterestPayment *YesNoIndicator `xml:"IntrstPmt,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName2 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Minimum number of days' notice a counterparty needs for terminating the transaction.
	TransactionCallDelay *Exact3NumericText `xml:"TxCallDely,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection59 `xml:"AcrdIntrstAmt,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection59 `xml:"TermntnTxAmt,omitempty"`

	// Provides additional information about the second leg in narrative form.
	SecondLegNarrative *RestrictedFINXMax140Text `xml:"ScndLegNrrtv,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails34) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails34) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails34) AddTerminationDate() *TerminationDate5Choice {
	s.TerminationDate = new(TerminationDate5Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails34) AddRateType() *RateType67Choice {
	s.RateType = new(RateType67Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails34) AddLegalFramework() *LegalFramework4Choice {
	s.LegalFramework = new(LegalFramework4Choice)
	return s.LegalFramework
}

func (s *SecuritiesFinancingTransactionDetails34) SetMaturityDateModification(value string) {
	s.MaturityDateModification = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails34) SetInterestPayment(value string) {
	s.InterestPayment = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails34) AddVariableRateSupport() *RateName2 {
	s.VariableRateSupport = new(RateName2)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails34) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails34) SetTransactionCallDelay(value string) {
	s.TransactionCallDelay = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails34) AddAccruedInterestAmount() *AmountAndDirection59 {
	s.AccruedInterestAmount = new(AmountAndDirection59)
	return s.AccruedInterestAmount
}

func (s *SecuritiesFinancingTransactionDetails34) AddTerminationTransactionAmount() *AmountAndDirection59 {
	s.TerminationTransactionAmount = new(AmountAndDirection59)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails34) SetSecondLegNarrative(value string) {
	s.SecondLegNarrative = (*RestrictedFINXMax140Text)(&value)
}
