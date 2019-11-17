package model

// Details of the closing of the securities financing transaction.
type SecuritiesFinancing10 struct {

	// Date/Time at which rate change has taken place.
	RateChangeDate *ISODateTime `xml:"RateChngDt,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType19Choice `xml:"RateTp,omitempty"`

	// Specifies whether the collateral position should be subject to automatic revaluation by the account servicer.
	Revaluation *Revaluation2Choice `xml:"Rvaltn,omitempty"`

	// Legal framework of the transaction.
	LegalFramework *LegalFramework1Code `xml:"LglFrmwk,omitempty"`

	// Identifies the computation method of accrued interest of the related financial instrument.
	InterestComputationMethod *InterestComputationMethod2Choice `xml:"IntrstCmptnMtd,omitempty"`

	// Index or support rate used together with the spread to calculate the repurchase rate.
	VariableRateSupport *RateName1 `xml:"VarblRateSpprt,omitempty"`

	// Repurchase rate used to calculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName1Choice `xml:"PricgRate,omitempty"`

	// Margin over or under an index that determines the repurchase rate, expressed as a rate or an amount.
	SpreadRate *SpreadRate1 `xml:"SprdRate,omitempty"`

	// Indicates whether or not the trade is callable.
	CallableTradeIndicator *YesNoIndicator `xml:"CllblTradInd,omitempty"`

	// Minimum number of days' notice a counterparty needs for terminating the transaction.
	TransactionCallDelay *Max3NumericText `xml:"TxCallDely,omitempty"`

	// Interest amount that has accrued in between two periods, for example, in between interest payment periods.
	AccruedInterestAmount *AmountAndDirection5 `xml:"AcrdIntrstAmt,omitempty"`

	// Interest rate that has been accrued in between coupon payment periods.
	AccruedInterestPercentage *PercentageRate `xml:"AcrdIntrstPctg,omitempty"`

	// Fixed amount of money that has to be paid (instead of interest) in the case of a recall or at the closing date.
	ForfeitAmount *AmountAndDirection5 `xml:"FrftAmt,omitempty"`

	// Difference between the amount of money of the first leg and the amount of the second leg of the transaction.
	PremiumAmount *AmountAndDirection5 `xml:"PrmAmt,omitempty"`

	// Amount of money to be settled per piece of collateral to close the transaction.
	ClosingAmountPerPiecesOfCollateral *AmountAndDirection5 `xml:"ClsgAmtPerPcsOfColl,omitempty"`

	// Indicates the total Number of collateral instructions involved in the transaction.
	TotalNumberOfCollateralInstructions *Max3NumericText `xml:"TtlNbOfCollInstrs,omitempty"`

	// Provides details for the securities financing transaction.
	FinancingAgreement *Agreement3 `xml:"FincgAgrmt,omitempty"`

	// Method applied to a lending transaction.
	LendingTransactionMethod *LendingTransactionMethod1Choice `xml:"LndgTxMtd,omitempty"`

	// Indicates if the contract is with or without an exchange of collateral.
	LendingWithCollateral *YesNoIndicator `xml:"LndgWthColl,omitempty"`

	// Identifies the underlying reason for the borrowing, for instance, sale on my behalf or on behalf of a third party.
	BorrowingReason *BorrowingReason1Choice `xml:"BrrwgRsn,omitempty"`

	// Indicates the type of collateral, for insatnce, security, bond, etc.
	CollateralType *CollateralType1Choice `xml:"CollTp,omitempty"`

	// Indicates whether or not the contract terms changed.
	ContractTermsModificationChanged *YesNoIndicator `xml:"CtrctTermsModChngd,omitempty"`

	// Interest rate to be paid as agreed between the counterparties.
	InterestRate *Rate2 `xml:"IntrstRate,omitempty"`

	// Rate to be paid by the Borrower to the Lender for the securities borrowed.
	BorrowingRate *Rate2 `xml:"BrrwgRate,omitempty"`

	// Method used to calculate the standard collateral amount.
	StandardCollateralRatio *Rate2 `xml:"StdCollRatio,omitempty"`

	// Percentage of earnings paid to shareholders in dividends.
	DividendRatio *Rate2 `xml:"DvddRatio,omitempty"`

	// Number of days the securities are lent or borrowed where the contract has an agreed closing date.
	NumberOfDaysLendingBorrowing *Number21Choice `xml:"NbOfDaysLndgBrrwg,omitempty"`

	// Specifies the standard collateral amount.
	StandardCollateralAmount *AmountAndDirection5 `xml:"StdCollAmt,omitempty"`

	// Interest rate tax that has been accrued in between coupon payment periods.
	AccruedInterestTax *YesNoIndicator `xml:"AcrdIntrstTax,omitempty"`

	// Number of days accrued at the instant of closing trade.
	EndNumberOfDaysAccrued *Max3Number `xml:"EndNbOfDaysAcrd,omitempty"`

	// End ratio of principal outstanding to the original balance.
	EndFactor *BaseOneRate `xml:"EndFctr,omitempty"`

	// Type of securities lending.
	SecuritiesLendingType *SecuritiesLendingType1Choice `xml:"SctiesLndgTp,omitempty"`

	// Indicates the possibility to terminate the securitiesc lending contract either by the borrower or lender before the expiration date.
	Reversible *Reversible1Choice `xml:"Rvsbl,omitempty"`

	// This is the minimum date at which the Borrower is allowed to give back the securities.
	MinimumDateForCallBack *ISODate `xml:"MinDtForCallBck,omitempty"`

	// Indicates that the contract can be rolled over.
	RollOver *YesNoIndicator `xml:"RollOver,omitempty"`

	// Indicates whether the securities lending fees can be paid periodically or at the end of the contract.
	PeriodicPayment *YesNoIndicator `xml:"PrdcPmt,omitempty"`

	// Indicates whether the trade is executed ex coupon.
	ExCoupon *YesNoIndicator `xml:"ExCpn,omitempty"`
}

func (s *SecuritiesFinancing10) SetRateChangeDate(value string) {
	s.RateChangeDate = (*ISODateTime)(&value)
}

func (s *SecuritiesFinancing10) AddRateType() *RateType19Choice {
	s.RateType = new(RateType19Choice)
	return s.RateType
}

func (s *SecuritiesFinancing10) AddRevaluation() *Revaluation2Choice {
	s.Revaluation = new(Revaluation2Choice)
	return s.Revaluation
}

func (s *SecuritiesFinancing10) SetLegalFramework(value string) {
	s.LegalFramework = (*LegalFramework1Code)(&value)
}

func (s *SecuritiesFinancing10) AddInterestComputationMethod() *InterestComputationMethod2Choice {
	s.InterestComputationMethod = new(InterestComputationMethod2Choice)
	return s.InterestComputationMethod
}

func (s *SecuritiesFinancing10) AddVariableRateSupport() *RateName1 {
	s.VariableRateSupport = new(RateName1)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancing10) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancing10) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancing10) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancing10) AddPricingRate() *RateOrName1Choice {
	s.PricingRate = new(RateOrName1Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancing10) AddSpreadRate() *SpreadRate1 {
	s.SpreadRate = new(SpreadRate1)
	return s.SpreadRate
}

func (s *SecuritiesFinancing10) SetCallableTradeIndicator(value string) {
	s.CallableTradeIndicator = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancing10) SetTransactionCallDelay(value string) {
	s.TransactionCallDelay = (*Max3NumericText)(&value)
}

func (s *SecuritiesFinancing10) AddAccruedInterestAmount() *AmountAndDirection5 {
	s.AccruedInterestAmount = new(AmountAndDirection5)
	return s.AccruedInterestAmount
}

func (s *SecuritiesFinancing10) SetAccruedInterestPercentage(value string) {
	s.AccruedInterestPercentage = (*PercentageRate)(&value)
}

func (s *SecuritiesFinancing10) AddForfeitAmount() *AmountAndDirection5 {
	s.ForfeitAmount = new(AmountAndDirection5)
	return s.ForfeitAmount
}

func (s *SecuritiesFinancing10) AddPremiumAmount() *AmountAndDirection5 {
	s.PremiumAmount = new(AmountAndDirection5)
	return s.PremiumAmount
}

func (s *SecuritiesFinancing10) AddClosingAmountPerPiecesOfCollateral() *AmountAndDirection5 {
	s.ClosingAmountPerPiecesOfCollateral = new(AmountAndDirection5)
	return s.ClosingAmountPerPiecesOfCollateral
}

func (s *SecuritiesFinancing10) SetTotalNumberOfCollateralInstructions(value string) {
	s.TotalNumberOfCollateralInstructions = (*Max3NumericText)(&value)
}

func (s *SecuritiesFinancing10) AddFinancingAgreement() *Agreement3 {
	s.FinancingAgreement = new(Agreement3)
	return s.FinancingAgreement
}

func (s *SecuritiesFinancing10) AddLendingTransactionMethod() *LendingTransactionMethod1Choice {
	s.LendingTransactionMethod = new(LendingTransactionMethod1Choice)
	return s.LendingTransactionMethod
}

func (s *SecuritiesFinancing10) SetLendingWithCollateral(value string) {
	s.LendingWithCollateral = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancing10) AddBorrowingReason() *BorrowingReason1Choice {
	s.BorrowingReason = new(BorrowingReason1Choice)
	return s.BorrowingReason
}

func (s *SecuritiesFinancing10) AddCollateralType() *CollateralType1Choice {
	s.CollateralType = new(CollateralType1Choice)
	return s.CollateralType
}

func (s *SecuritiesFinancing10) SetContractTermsModificationChanged(value string) {
	s.ContractTermsModificationChanged = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancing10) AddInterestRate() *Rate2 {
	s.InterestRate = new(Rate2)
	return s.InterestRate
}

func (s *SecuritiesFinancing10) AddBorrowingRate() *Rate2 {
	s.BorrowingRate = new(Rate2)
	return s.BorrowingRate
}

func (s *SecuritiesFinancing10) AddStandardCollateralRatio() *Rate2 {
	s.StandardCollateralRatio = new(Rate2)
	return s.StandardCollateralRatio
}

func (s *SecuritiesFinancing10) AddDividendRatio() *Rate2 {
	s.DividendRatio = new(Rate2)
	return s.DividendRatio
}

func (s *SecuritiesFinancing10) AddNumberOfDaysLendingBorrowing() *Number21Choice {
	s.NumberOfDaysLendingBorrowing = new(Number21Choice)
	return s.NumberOfDaysLendingBorrowing
}

func (s *SecuritiesFinancing10) AddStandardCollateralAmount() *AmountAndDirection5 {
	s.StandardCollateralAmount = new(AmountAndDirection5)
	return s.StandardCollateralAmount
}

func (s *SecuritiesFinancing10) SetAccruedInterestTax(value string) {
	s.AccruedInterestTax = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancing10) SetEndNumberOfDaysAccrued(value string) {
	s.EndNumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesFinancing10) SetEndFactor(value string) {
	s.EndFactor = (*BaseOneRate)(&value)
}

func (s *SecuritiesFinancing10) AddSecuritiesLendingType() *SecuritiesLendingType1Choice {
	s.SecuritiesLendingType = new(SecuritiesLendingType1Choice)
	return s.SecuritiesLendingType
}

func (s *SecuritiesFinancing10) AddReversible() *Reversible1Choice {
	s.Reversible = new(Reversible1Choice)
	return s.Reversible
}

func (s *SecuritiesFinancing10) SetMinimumDateForCallBack(value string) {
	s.MinimumDateForCallBack = (*ISODate)(&value)
}

func (s *SecuritiesFinancing10) SetRollOver(value string) {
	s.RollOver = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancing10) SetPeriodicPayment(value string) {
	s.PeriodicPayment = (*YesNoIndicator)(&value)
}

func (s *SecuritiesFinancing10) SetExCoupon(value string) {
	s.ExCoupon = (*YesNoIndicator)(&value)
}
