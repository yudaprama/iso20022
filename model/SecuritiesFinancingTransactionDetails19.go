package model

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails19 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *Max35Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *Max35Text `xml:"ClsgLegId,omitempty"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate2Choice `xml:"TermntnDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Earliest date/time at which the call back can take place.
	EarliestCallBackDate *DateAndDateTimeChoice `xml:"EarlstCallBckDt,omitempty"`

	// Date/time at which the commission is calculated.
	CommissionCalculationDate *DateAndDateTimeChoice `xml:"ComssnClctnDt,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType5Choice `xml:"RateTp,omitempty"`

	// Specifies whether the collateral position should be subject to automatic revaluation by the account servicer.
	Revaluation *RevaluationIndicator1Choice `xml:"Rvaltn,omitempty"`

	// Legal framework of the transaction.
	LegalFramework *LegalFramework1Choice `xml:"LglFrmwk,omitempty"`

	// Identifies the computation method of accrued interest of the related financial instrument.
	InterestComputationMethod *InterestComputationMethodFormat1Choice `xml:"IntrstCmptnMtd,omitempty"`

	// Specifies whether the maturity date of the securities financing transaction may be modified.
	MaturityDateModification *YesNoIndicator `xml:"MtrtyDtMod,omitempty"`

	// Specifies whether the interest is to be paid to the collateral taker. If set to no, the interest is paid to the collateral giver.
	InterestPayment *YesNoIndicator `xml:"IntrstPmt,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName1 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate paid in the context of a securities financing transaction.
	ChargesRate *Rate2 `xml:"ChrgsRate,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName1Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Minimum number of days' notice a counterparty needs for terminating the transaction.
	TransactionCallDelay *Exact3NumericText `xml:"TxCallDely,omitempty"`

	// Total number of collateral instructions involved in the transaction.
	TotalNumberOfCollateralInstructions *Exact3NumericText `xml:"TtlNbOfCollInstrs,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection4 `xml:"LclBrkrComssn,omitempty"`

	// Principal amount of a trade (for second leg).
	DealAmount *AmountAndDirection4 `xml:"DealAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection4 `xml:"AcrdIntrstAmt,omitempty"`

	// Fixed amount of money that has to be paid (instead of interest) in the case of a recall or at the closing date.
	ForfeitAmount *AmountAndDirection4 `xml:"FrftAmt,omitempty"`

	// Difference between the amount of money of the first leg and the amount of the second leg of the transaction.
	PremiumAmount *AmountAndDirection4 `xml:"PrmAmt,omitempty"`

	// Amount of money to be settled per piece of collateral to terminate the transaction.
	TerminationAmountPerPieceOfCollateral *AmountAndDirection4 `xml:"TermntnAmtPerPcOfColl,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection4 `xml:"TermntnTxAmt,omitempty"`

	// Provides additional information about the second leg in narrative form.
	SecondLegNarrative *Max140Text `xml:"ScndLegNrrtv,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails19) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails19) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails19) AddTerminationDate() *TerminationDate2Choice {
	s.TerminationDate = new(TerminationDate2Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails19) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails19) AddEarliestCallBackDate() *DateAndDateTimeChoice {
	s.EarliestCallBackDate = new(DateAndDateTimeChoice)
	return s.EarliestCallBackDate
}

func (s *SecuritiesFinancingTransactionDetails19) AddCommissionCalculationDate() *DateAndDateTimeChoice {
	s.CommissionCalculationDate = new(DateAndDateTimeChoice)
	return s.CommissionCalculationDate
}

func (s *SecuritiesFinancingTransactionDetails19) AddRateType() *RateType5Choice {
	s.RateType = new(RateType5Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails19) AddRevaluation() *RevaluationIndicator1Choice {
	s.Revaluation = new(RevaluationIndicator1Choice)
	return s.Revaluation
}

func (s *SecuritiesFinancingTransactionDetails19) AddLegalFramework() *LegalFramework1Choice {
	s.LegalFramework = new(LegalFramework1Choice)
	return s.LegalFramework
}

func (s *SecuritiesFinancingTransactionDetails19) AddInterestComputationMethod() *InterestComputationMethodFormat1Choice {
	s.InterestComputationMethod = new(InterestComputationMethodFormat1Choice)
	return s.InterestComputationMethod
}

func (s *SecuritiesFinancingTransactionDetails19) SetMaturityDateModification(value string) {
	s.MaturityDateModification = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails19) SetInterestPayment(value string) {
	s.InterestPayment = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails19) AddVariableRateSupport() *RateName1 {
	s.VariableRateSupport = new(RateName1)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails19) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails19) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails19) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails19) AddChargesRate() *Rate2 {
	s.ChargesRate = new(Rate2)
	return s.ChargesRate
}

func (s *SecuritiesFinancingTransactionDetails19) AddPricingRate() *RateOrName1Choice {
	s.PricingRate = new(RateOrName1Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails19) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails19) SetTransactionCallDelay(value string) {
	s.TransactionCallDelay = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails19) SetTotalNumberOfCollateralInstructions(value string) {
	s.TotalNumberOfCollateralInstructions = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails19) AddLocalBrokerCommission() *AmountAndDirection4 {
	s.LocalBrokerCommission = new(AmountAndDirection4)
	return s.LocalBrokerCommission
}

func (s *SecuritiesFinancingTransactionDetails19) AddDealAmount() *AmountAndDirection4 {
	s.DealAmount = new(AmountAndDirection4)
	return s.DealAmount
}

func (s *SecuritiesFinancingTransactionDetails19) AddAccruedInterestAmount() *AmountAndDirection4 {
	s.AccruedInterestAmount = new(AmountAndDirection4)
	return s.AccruedInterestAmount
}

func (s *SecuritiesFinancingTransactionDetails19) AddForfeitAmount() *AmountAndDirection4 {
	s.ForfeitAmount = new(AmountAndDirection4)
	return s.ForfeitAmount
}

func (s *SecuritiesFinancingTransactionDetails19) AddPremiumAmount() *AmountAndDirection4 {
	s.PremiumAmount = new(AmountAndDirection4)
	return s.PremiumAmount
}

func (s *SecuritiesFinancingTransactionDetails19) AddTerminationAmountPerPieceOfCollateral() *AmountAndDirection4 {
	s.TerminationAmountPerPieceOfCollateral = new(AmountAndDirection4)
	return s.TerminationAmountPerPieceOfCollateral
}

func (s *SecuritiesFinancingTransactionDetails19) AddTerminationTransactionAmount() *AmountAndDirection4 {
	s.TerminationTransactionAmount = new(AmountAndDirection4)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails19) SetSecondLegNarrative(value string) {
	s.SecondLegNarrative = (*Max140Text)(&value)
}
