package model

// Elements characterising a financial instrument.
type FinancialInstrumentAttributes20 struct {

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification5 `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Specifies the form, this is, ownership, of the security.
	RegistrationForm *FormOfSecurity2Choice `xml:"RegnForm,omitempty"`

	// Specifies the frequency of an interest payment.
	PaymentFrequency *Frequency3Choice `xml:"PmtFrqcy,omitempty"`

	// Status of payment of a security at a particular time.
	PaymentStatus *SecuritiesPaymentStatus2Choice `xml:"PmtSts,omitempty"`

	// Indicates the direction of payment for asset or mortgage backed securities, this is, whether the repaid capital is distributed (payment direction is down) or capitalized (payment direction is up).
	PaymentDirection *PaymentDirection2Choice `xml:"PmtDrctn,omitempty"`

	// Specifies the frequency of change to the variable rate of an interest bearing instrument.
	VariableRateChangeFrequency *Frequency3Choice `xml:"VarblRateChngFrqcy,omitempty"`

	// Indicates the  level of priority to claim on income and assets of the company in case of the payment of dividends and in the event of a bankruptcy, for example, ordinary/common stocks, preferred stocks, subordinated debt, etc.
	PreferenceToIncome *PreferenceToIncome2Choice `xml:"PrefToIncm,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI) codification, for example, common share with voting rights, fully paid, or registered.
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised (American, European, Bermudan)
	OptionStyle *OptionStyle4Choice `xml:"OptnStyle,omitempty"`

	// Specifies whether it is a Call option (right to purchase a specific underlying asset) or a Put option (right to sell a specific underlying asset).
	OptionType *OptionType2Choice `xml:"OptnTp,omitempty"`

	// Currency in which a security is issued or redenominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	CouponDate *ISODate `xml:"CpnDt,omitempty"`

	// Date on which a privilege expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date at which the interest rate of an interest bearing security will be calculated and reset, according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Next date at which the issuer has the right to pay the security prior to maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date at which the holder has the right to ask for redemption of the security prior to final maturity.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// First date at which a security begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Date at which the first interest payment is due to holders of the security.
	FirstPaymentDate *ISODate `xml:"FrstPmtDt,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that was applicable before the current factor and defines the outstanding principal of the financial instrument (for factored securities).
	PreviousFactor *BaseOneRate `xml:"PrvsFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 defining the outstanding principal of the financial instrument (for factored securities).
	CurrentFactor *BaseOneRate `xml:"CurFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that will be applicable as of the next factor date and defines the outstanding principal of the financial instrument (for factored securities).
	NextFactor *BaseOneRate `xml:"NxtFctr,omitempty"`

	// Per annum ratio of interest paid to the principal amount of the financial instrument for a specific period of time.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Specifies the reference rate for fixed income instruments where the
	// price of the instrument is indexed to the price of an underlying benchmark.
	IndexRateBasis *PercentageRate `xml:"IndxRateBsis,omitempty"`

	// Number of the coupon attached to the physical security.
	CouponAttachedNumber *Number2Choice `xml:"CpnAttchdNb,omitempty"`

	// Number identifying a group of underlying assets assigned by the issuer of a factored security.
	PoolNumber *Number2Choice `xml:"PoolNb,omitempty"`

	// Indicates whether the interest rate of an interest bearing instrument is reset periodically.
	VariableRateIndicator *YesNoIndicator `xml:"VarblRateInd,omitempty"`

	// Indicates whether the issuer has the right to pay the security prior to maturity. Also called RetractableIndicator.
	CallableIndicator *YesNoIndicator `xml:"CllblInd,omitempty"`

	// Indicates whether the holder has the right to ask for redemption of the security prior to final maturity. Also called RedeemableIndicator.
	PutableIndicator *YesNoIndicator `xml:"PutblInd,omitempty"`

	// Value of the price, for example, as a currency and value per unit or as a percentage.
	MarketOrIndicativePrice *PriceType1Choice `xml:"MktOrIndctvPric,omitempty"`

	// Predetermined price at which the holder of a derivative will buy or sell the underlying instrument.
	ExercisePrice *Price2 `xml:"ExrcPric,omitempty"`

	// Pre-determined price at which the holder of a right is entitled to buy the underlying instrument.
	SubscriptionPrice *Price2 `xml:"SbcptPric,omitempty"`

	// Price of one target security in the conversion.
	ConversionPrice *Price2 `xml:"ConvsPric,omitempty"`

	// Predetermined price at which the holder will have to buy or sell the underlying instrument.
	StrikePrice *Price2 `xml:"StrkPric,omitempty"`

	// Indicates the minimum tradable quantity of a security.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`

	// Identification of the underlying security by an ISIN.
	UnderlyingFinancialInstrumentIdentification []*SecurityIdentification14 `xml:"UndrlygFinInstrmId,omitempty"`

	// Provides additional information about the financial instrument in narrative form.
	FinancialInstrumentAttributeAdditionalDetails *Max350Text `xml:"FinInstrmAttrAddtlDtls,omitempty"`
}

func (f *FinancialInstrumentAttributes20) AddPlaceOfListing() *MarketIdentification5 {
	f.PlaceOfListing = new(MarketIdentification5)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes20) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes20) AddRegistrationForm() *FormOfSecurity2Choice {
	f.RegistrationForm = new(FormOfSecurity2Choice)
	return f.RegistrationForm
}

func (f *FinancialInstrumentAttributes20) AddPaymentFrequency() *Frequency3Choice {
	f.PaymentFrequency = new(Frequency3Choice)
	return f.PaymentFrequency
}

func (f *FinancialInstrumentAttributes20) AddPaymentStatus() *SecuritiesPaymentStatus2Choice {
	f.PaymentStatus = new(SecuritiesPaymentStatus2Choice)
	return f.PaymentStatus
}

func (f *FinancialInstrumentAttributes20) AddPaymentDirection() *PaymentDirection2Choice {
	f.PaymentDirection = new(PaymentDirection2Choice)
	return f.PaymentDirection
}

func (f *FinancialInstrumentAttributes20) AddVariableRateChangeFrequency() *Frequency3Choice {
	f.VariableRateChangeFrequency = new(Frequency3Choice)
	return f.VariableRateChangeFrequency
}

func (f *FinancialInstrumentAttributes20) AddPreferenceToIncome() *PreferenceToIncome2Choice {
	f.PreferenceToIncome = new(PreferenceToIncome2Choice)
	return f.PreferenceToIncome
}

func (f *FinancialInstrumentAttributes20) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes20) AddOptionStyle() *OptionStyle4Choice {
	f.OptionStyle = new(OptionStyle4Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes20) AddOptionType() *OptionType2Choice {
	f.OptionType = new(OptionType2Choice)
	return f.OptionType
}

func (f *FinancialInstrumentAttributes20) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes20) SetCouponDate(value string) {
	f.CouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetFirstPaymentDate(value string) {
	f.FirstPaymentDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetCurrentFactor(value string) {
	f.CurrentFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetNextFactor(value string) {
	f.NextFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes20) SetIndexRateBasis(value string) {
	f.IndexRateBasis = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes20) AddCouponAttachedNumber() *Number2Choice {
	f.CouponAttachedNumber = new(Number2Choice)
	return f.CouponAttachedNumber
}

func (f *FinancialInstrumentAttributes20) AddPoolNumber() *Number2Choice {
	f.PoolNumber = new(Number2Choice)
	return f.PoolNumber
}

func (f *FinancialInstrumentAttributes20) SetVariableRateIndicator(value string) {
	f.VariableRateIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes20) SetCallableIndicator(value string) {
	f.CallableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes20) SetPutableIndicator(value string) {
	f.PutableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes20) AddMarketOrIndicativePrice() *PriceType1Choice {
	f.MarketOrIndicativePrice = new(PriceType1Choice)
	return f.MarketOrIndicativePrice
}

func (f *FinancialInstrumentAttributes20) AddExercisePrice() *Price2 {
	f.ExercisePrice = new(Price2)
	return f.ExercisePrice
}

func (f *FinancialInstrumentAttributes20) AddSubscriptionPrice() *Price2 {
	f.SubscriptionPrice = new(Price2)
	return f.SubscriptionPrice
}

func (f *FinancialInstrumentAttributes20) AddConversionPrice() *Price2 {
	f.ConversionPrice = new(Price2)
	return f.ConversionPrice
}

func (f *FinancialInstrumentAttributes20) AddStrikePrice() *Price2 {
	f.StrikePrice = new(Price2)
	return f.StrikePrice
}

func (f *FinancialInstrumentAttributes20) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes20) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes20) AddUnderlyingFinancialInstrumentIdentification() *SecurityIdentification14 {
	newValue := new(SecurityIdentification14)
	f.UnderlyingFinancialInstrumentIdentification = append(f.UnderlyingFinancialInstrumentIdentification, newValue)
	return newValue
}

func (f *FinancialInstrumentAttributes20) SetFinancialInstrumentAttributeAdditionalDetails(value string) {
	f.FinancialInstrumentAttributeAdditionalDetails = (*Max350Text)(&value)
}
