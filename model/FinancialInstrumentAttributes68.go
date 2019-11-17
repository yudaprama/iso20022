package model

// Description of the financial instrument.
type FinancialInstrumentAttributes68 struct {

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification19 `xml:"SctyId"`

	// Quantity of entitled intermediate securities based on the balance of underlying securities.
	Quantity *DecimalNumber `xml:"Qty,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableEntitlementStatusTypeFormat3Choice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType25Choice `xml:"FrctnDspstn,omitempty"`

	// Quantity of intermediate securities awarded for a given quantity of underlying security.
	IntermediateSecuritiesToUnderlyingRatio *QuantityToQuantityRatio1 `xml:"IntrmdtSctiesToUndrlygRatio,omitempty"`

	// Last reported/known price of a financial instrument in a market.
	MarketPrice *AmountPrice2 `xml:"MktPric,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *DateFormat30Choice `xml:"XpryDt"`

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateFormat30Choice `xml:"PstngDt"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period4 `xml:"TradgPrd,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *BalanceFormat5Choice `xml:"UinstdBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat5Choice `xml:"InstdBal,omitempty"`
}

func (f *FinancialInstrumentAttributes68) AddSecurityIdentification() *SecurityIdentification19 {
	f.SecurityIdentification = new(SecurityIdentification19)
	return f.SecurityIdentification
}

func (f *FinancialInstrumentAttributes68) SetQuantity(value string) {
	f.Quantity = (*DecimalNumber)(&value)
}

func (f *FinancialInstrumentAttributes68) AddRenounceableEntitlementStatusType() *RenounceableEntitlementStatusTypeFormat3Choice {
	f.RenounceableEntitlementStatusType = new(RenounceableEntitlementStatusTypeFormat3Choice)
	return f.RenounceableEntitlementStatusType
}

func (f *FinancialInstrumentAttributes68) AddFractionDisposition() *FractionDispositionType25Choice {
	f.FractionDisposition = new(FractionDispositionType25Choice)
	return f.FractionDisposition
}

func (f *FinancialInstrumentAttributes68) AddIntermediateSecuritiesToUnderlyingRatio() *QuantityToQuantityRatio1 {
	f.IntermediateSecuritiesToUnderlyingRatio = new(QuantityToQuantityRatio1)
	return f.IntermediateSecuritiesToUnderlyingRatio
}

func (f *FinancialInstrumentAttributes68) AddMarketPrice() *AmountPrice2 {
	f.MarketPrice = new(AmountPrice2)
	return f.MarketPrice
}

func (f *FinancialInstrumentAttributes68) AddExpiryDate() *DateFormat30Choice {
	f.ExpiryDate = new(DateFormat30Choice)
	return f.ExpiryDate
}

func (f *FinancialInstrumentAttributes68) AddPostingDate() *DateFormat30Choice {
	f.PostingDate = new(DateFormat30Choice)
	return f.PostingDate
}

func (f *FinancialInstrumentAttributes68) AddTradingPeriod() *Period4 {
	f.TradingPeriod = new(Period4)
	return f.TradingPeriod
}

func (f *FinancialInstrumentAttributes68) AddUninstructedBalance() *BalanceFormat5Choice {
	f.UninstructedBalance = new(BalanceFormat5Choice)
	return f.UninstructedBalance
}

func (f *FinancialInstrumentAttributes68) AddInstructedBalance() *BalanceFormat5Choice {
	f.InstructedBalance = new(BalanceFormat5Choice)
	return f.InstructedBalance
}
