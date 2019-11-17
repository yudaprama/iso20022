package model

// Description of the financial instrument.
type FinancialInstrumentAttributes17 struct {

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId"`

	// Quantity of entitled intermediate securities based on the balance of underlying securities.
	Quantity *DecimalNumber `xml:"Qty,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableEntitlementStatusTypeFormat1Choice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType9Choice `xml:"FrctnDspstn,omitempty"`

	// Quantity of intermediate securities awarded for a given quantity of underlying security.
	IntermediateSecuritiesToUnderlyingRatio *QuantityToQuantityRatio1 `xml:"IntrmdtSctiesToUndrlygRatio,omitempty"`

	// Last reported/known price of a financial instrument in a market.
	MarketPrice *AmountPrice2 `xml:"MktPric,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *DateFormat16Choice `xml:"XpryDt"`

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateFormat16Choice `xml:"PstngDt"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period4 `xml:"TradgPrd,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *BalanceFormat1Choice `xml:"UinstdBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat1Choice `xml:"InstdBal,omitempty"`
}

func (f *FinancialInstrumentAttributes17) AddSecurityIdentification() *SecurityIdentification14 {
	f.SecurityIdentification = new(SecurityIdentification14)
	return f.SecurityIdentification
}

func (f *FinancialInstrumentAttributes17) SetQuantity(value string) {
	f.Quantity = (*DecimalNumber)(&value)
}

func (f *FinancialInstrumentAttributes17) AddRenounceableEntitlementStatusType() *RenounceableEntitlementStatusTypeFormat1Choice {
	f.RenounceableEntitlementStatusType = new(RenounceableEntitlementStatusTypeFormat1Choice)
	return f.RenounceableEntitlementStatusType
}

func (f *FinancialInstrumentAttributes17) AddFractionDisposition() *FractionDispositionType9Choice {
	f.FractionDisposition = new(FractionDispositionType9Choice)
	return f.FractionDisposition
}

func (f *FinancialInstrumentAttributes17) AddIntermediateSecuritiesToUnderlyingRatio() *QuantityToQuantityRatio1 {
	f.IntermediateSecuritiesToUnderlyingRatio = new(QuantityToQuantityRatio1)
	return f.IntermediateSecuritiesToUnderlyingRatio
}

func (f *FinancialInstrumentAttributes17) AddMarketPrice() *AmountPrice2 {
	f.MarketPrice = new(AmountPrice2)
	return f.MarketPrice
}

func (f *FinancialInstrumentAttributes17) AddExpiryDate() *DateFormat16Choice {
	f.ExpiryDate = new(DateFormat16Choice)
	return f.ExpiryDate
}

func (f *FinancialInstrumentAttributes17) AddPostingDate() *DateFormat16Choice {
	f.PostingDate = new(DateFormat16Choice)
	return f.PostingDate
}

func (f *FinancialInstrumentAttributes17) AddTradingPeriod() *Period4 {
	f.TradingPeriod = new(Period4)
	return f.TradingPeriod
}

func (f *FinancialInstrumentAttributes17) AddUninstructedBalance() *BalanceFormat1Choice {
	f.UninstructedBalance = new(BalanceFormat1Choice)
	return f.UninstructedBalance
}

func (f *FinancialInstrumentAttributes17) AddInstructedBalance() *BalanceFormat1Choice {
	f.InstructedBalance = new(BalanceFormat1Choice)
	return f.InstructedBalance
}
