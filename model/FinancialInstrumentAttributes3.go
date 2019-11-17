package model

// Description of the financial instrument.
type FinancialInstrumentAttributes3 struct {

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification11 `xml:"SctyId"`

	// Quantity of entitled intermediate securities based on the balance of underlying securities.
	Quantity *DecimalNumber `xml:"Qty,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableEntitlementStatusTypeFormat1Choice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType3Choice `xml:"FrctnDspstn,omitempty"`

	// Quantity of intermediate securities awarded for a given quantity of underlying security.
	IntermediateSecuritiesToUnderlyingRatio *QuantityToQuantityRatio1 `xml:"IntrmdtSctiesToUndrlygRatio,omitempty"`

	// Last reported/known price of a financial instrument in a market.
	MarketPrice *AmountPrice2 `xml:"MktPric,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *DateFormat5Choice `xml:"XpryDt"`

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateFormat5Choice `xml:"PstngDt"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period3 `xml:"TradgPrd,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *BalanceFormat1Choice `xml:"UinstdBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat1Choice `xml:"InstdBal,omitempty"`
}

func (f *FinancialInstrumentAttributes3) AddSecurityIdentification() *SecurityIdentification11 {
	f.SecurityIdentification = new(SecurityIdentification11)
	return f.SecurityIdentification
}

func (f *FinancialInstrumentAttributes3) SetQuantity(value string) {
	f.Quantity = (*DecimalNumber)(&value)
}

func (f *FinancialInstrumentAttributes3) AddRenounceableEntitlementStatusType() *RenounceableEntitlementStatusTypeFormat1Choice {
	f.RenounceableEntitlementStatusType = new(RenounceableEntitlementStatusTypeFormat1Choice)
	return f.RenounceableEntitlementStatusType
}

func (f *FinancialInstrumentAttributes3) AddFractionDisposition() *FractionDispositionType3Choice {
	f.FractionDisposition = new(FractionDispositionType3Choice)
	return f.FractionDisposition
}

func (f *FinancialInstrumentAttributes3) AddIntermediateSecuritiesToUnderlyingRatio() *QuantityToQuantityRatio1 {
	f.IntermediateSecuritiesToUnderlyingRatio = new(QuantityToQuantityRatio1)
	return f.IntermediateSecuritiesToUnderlyingRatio
}

func (f *FinancialInstrumentAttributes3) AddMarketPrice() *AmountPrice2 {
	f.MarketPrice = new(AmountPrice2)
	return f.MarketPrice
}

func (f *FinancialInstrumentAttributes3) AddExpiryDate() *DateFormat5Choice {
	f.ExpiryDate = new(DateFormat5Choice)
	return f.ExpiryDate
}

func (f *FinancialInstrumentAttributes3) AddPostingDate() *DateFormat5Choice {
	f.PostingDate = new(DateFormat5Choice)
	return f.PostingDate
}

func (f *FinancialInstrumentAttributes3) AddTradingPeriod() *Period3 {
	f.TradingPeriod = new(Period3)
	return f.TradingPeriod
}

func (f *FinancialInstrumentAttributes3) AddUninstructedBalance() *BalanceFormat1Choice {
	f.UninstructedBalance = new(BalanceFormat1Choice)
	return f.UninstructedBalance
}

func (f *FinancialInstrumentAttributes3) AddInstructedBalance() *BalanceFormat1Choice {
	f.InstructedBalance = new(BalanceFormat1Choice)
	return f.InstructedBalance
}
