package model

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails29 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *Max35Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *Max35Text `xml:"ClsgLegId,omitempty"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate4Choice `xml:"TermntnDt,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType35Choice `xml:"RateTp,omitempty"`

	// Legal framework of the transaction.
	LegalFramework *LegalFramework3Choice `xml:"LglFrmwk,omitempty"`

	// Specifies whether the maturity date of the securities financing transaction may be modified.
	MaturityDateModification *YesNoIndicator `xml:"MtrtyDtMod,omitempty"`

	// Specifies whether the interest is to be paid to the collateral taker. If set to no, the interest is paid to the collateral giver.
	InterestPayment *YesNoIndicator `xml:"IntrstPmt,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName1 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Minimum number of days' notice a counterparty needs for terminating the transaction.
	TransactionCallDelay *Exact3NumericText `xml:"TxCallDely,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection21 `xml:"AcrdIntrstAmt,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection21 `xml:"TermntnTxAmt,omitempty"`

	// Provides additional information about the second leg in narrative form.
	SecondLegNarrative *Max140Text `xml:"ScndLegNrrtv,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails29) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails29) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails29) AddTerminationDate() *TerminationDate4Choice {
	s.TerminationDate = new(TerminationDate4Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails29) AddRateType() *RateType35Choice {
	s.RateType = new(RateType35Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails29) AddLegalFramework() *LegalFramework3Choice {
	s.LegalFramework = new(LegalFramework3Choice)
	return s.LegalFramework
}

func (s *SecuritiesFinancingTransactionDetails29) SetMaturityDateModification(value string) {
	s.MaturityDateModification = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails29) SetInterestPayment(value string) {
	s.InterestPayment = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails29) AddVariableRateSupport() *RateName1 {
	s.VariableRateSupport = new(RateName1)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails29) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails29) SetTransactionCallDelay(value string) {
	s.TransactionCallDelay = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails29) AddAccruedInterestAmount() *AmountAndDirection21 {
	s.AccruedInterestAmount = new(AmountAndDirection21)
	return s.AccruedInterestAmount
}

func (s *SecuritiesFinancingTransactionDetails29) AddTerminationTransactionAmount() *AmountAndDirection21 {
	s.TerminationTransactionAmount = new(AmountAndDirection21)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails29) SetSecondLegNarrative(value string) {
	s.SecondLegNarrative = (*Max140Text)(&value)
}
