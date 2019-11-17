package model

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails32 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *RestrictedFINXMax16Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *RestrictedFINXMax16Text `xml:"ClsgLegId,omitempty"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate5Choice `xml:"TermntnDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Earliest date/time at which the call back can take place.
	EarliestCallBackDate *DateAndDateTimeChoice `xml:"EarlstCallBckDt,omitempty"`

	// Date/time at which the commission is calculated.
	CommissionCalculationDate *DateAndDateTimeChoice `xml:"ComssnClctnDt,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType67Choice `xml:"RateTp,omitempty"`

	// Specifies whether the collateral position should be subject to automatic revaluation by the account servicer.
	Revaluation *RevaluationIndicator4Choice `xml:"Rvaltn,omitempty"`

	// Legal framework of the transaction.
	LegalFramework *LegalFramework4Choice `xml:"LglFrmwk,omitempty"`

	// Identifies the computation method of accrued interest of the related financial instrument.
	InterestComputationMethod *InterestComputationMethodFormat5Choice `xml:"IntrstCmptnMtd,omitempty"`

	// Specifies whether the maturity date of the securities financing transaction may be modified.
	MaturityDateModification *YesNoIndicator `xml:"MtrtyDtMod,omitempty"`

	// Specifies whether the interest is to be paid to the collateral taker. If set to no, the interest is paid to the collateral giver.
	InterestPayment *YesNoIndicator `xml:"IntrstPmt,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName2 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate paid in the context of a securities financing transaction.
	ChargesRate *Rate2 `xml:"ChrgsRate,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName2Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Minimum number of days' notice a counterparty needs for terminating the transaction.
	TransactionCallDelay *Exact3NumericText `xml:"TxCallDely,omitempty"`

	// Total number of collateral instructions involved in the transaction.
	TotalNumberOfCollateralInstructions *Exact3NumericText `xml:"TtlNbOfCollInstrs,omitempty"`

	// Amount of commission paid to a local broker.
	LocalBrokerCommission *AmountAndDirection59 `xml:"LclBrkrComssn,omitempty"`

	// Principal amount of a trade (for second leg).
	DealAmount *AmountAndDirection59 `xml:"DealAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection59 `xml:"AcrdIntrstAmt,omitempty"`

	// Fixed amount of money that has to be paid (instead of interest) in the case of a recall or at the closing date.
	ForfeitAmount *AmountAndDirection59 `xml:"FrftAmt,omitempty"`

	// Difference between the amount of money of the first leg and the amount of the second leg of the transaction.
	PremiumAmount *AmountAndDirection59 `xml:"PrmAmt,omitempty"`

	// Amount of money to be settled per piece of collateral to terminate the transaction.
	TerminationAmountPerPieceOfCollateral *AmountAndDirection59 `xml:"TermntnAmtPerPcOfColl,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection59 `xml:"TermntnTxAmt,omitempty"`

	// Provides additional information about the second leg in narrative form.
	SecondLegNarrative *RestrictedFINXMax140Text `xml:"ScndLegNrrtv,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails32) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails32) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails32) AddTerminationDate() *TerminationDate5Choice {
	s.TerminationDate = new(TerminationDate5Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails32) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails32) AddEarliestCallBackDate() *DateAndDateTimeChoice {
	s.EarliestCallBackDate = new(DateAndDateTimeChoice)
	return s.EarliestCallBackDate
}

func (s *SecuritiesFinancingTransactionDetails32) AddCommissionCalculationDate() *DateAndDateTimeChoice {
	s.CommissionCalculationDate = new(DateAndDateTimeChoice)
	return s.CommissionCalculationDate
}

func (s *SecuritiesFinancingTransactionDetails32) AddRateType() *RateType67Choice {
	s.RateType = new(RateType67Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails32) AddRevaluation() *RevaluationIndicator4Choice {
	s.Revaluation = new(RevaluationIndicator4Choice)
	return s.Revaluation
}

func (s *SecuritiesFinancingTransactionDetails32) AddLegalFramework() *LegalFramework4Choice {
	s.LegalFramework = new(LegalFramework4Choice)
	return s.LegalFramework
}

func (s *SecuritiesFinancingTransactionDetails32) AddInterestComputationMethod() *InterestComputationMethodFormat5Choice {
	s.InterestComputationMethod = new(InterestComputationMethodFormat5Choice)
	return s.InterestComputationMethod
}

func (s *SecuritiesFinancingTransactionDetails32) SetMaturityDateModification(value string) {
	s.MaturityDateModification = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails32) SetInterestPayment(value string) {
	s.InterestPayment = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancingTransactionDetails32) AddVariableRateSupport() *RateName2 {
	s.VariableRateSupport = new(RateName2)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails32) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails32) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails32) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails32) AddChargesRate() *Rate2 {
	s.ChargesRate = new(Rate2)
	return s.ChargesRate
}

func (s *SecuritiesFinancingTransactionDetails32) AddPricingRate() *RateOrName2Choice {
	s.PricingRate = new(RateOrName2Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails32) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails32) SetTransactionCallDelay(value string) {
	s.TransactionCallDelay = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails32) SetTotalNumberOfCollateralInstructions(value string) {
	s.TotalNumberOfCollateralInstructions = (*Exact3NumericText)(&value)
}

func (s *SecuritiesFinancingTransactionDetails32) AddLocalBrokerCommission() *AmountAndDirection59 {
	s.LocalBrokerCommission = new(AmountAndDirection59)
	return s.LocalBrokerCommission
}

func (s *SecuritiesFinancingTransactionDetails32) AddDealAmount() *AmountAndDirection59 {
	s.DealAmount = new(AmountAndDirection59)
	return s.DealAmount
}

func (s *SecuritiesFinancingTransactionDetails32) AddAccruedInterestAmount() *AmountAndDirection59 {
	s.AccruedInterestAmount = new(AmountAndDirection59)
	return s.AccruedInterestAmount
}

func (s *SecuritiesFinancingTransactionDetails32) AddForfeitAmount() *AmountAndDirection59 {
	s.ForfeitAmount = new(AmountAndDirection59)
	return s.ForfeitAmount
}

func (s *SecuritiesFinancingTransactionDetails32) AddPremiumAmount() *AmountAndDirection59 {
	s.PremiumAmount = new(AmountAndDirection59)
	return s.PremiumAmount
}

func (s *SecuritiesFinancingTransactionDetails32) AddTerminationAmountPerPieceOfCollateral() *AmountAndDirection59 {
	s.TerminationAmountPerPieceOfCollateral = new(AmountAndDirection59)
	return s.TerminationAmountPerPieceOfCollateral
}

func (s *SecuritiesFinancingTransactionDetails32) AddTerminationTransactionAmount() *AmountAndDirection59 {
	s.TerminationTransactionAmount = new(AmountAndDirection59)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails32) SetSecondLegNarrative(value string) {
	s.SecondLegNarrative = (*RestrictedFINXMax140Text)(&value)
}
