package model

// Description of the financial instrument.
type FinancialInstrumentAttributes73 struct {

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification20 `xml:"SctyId"`

	// Quantity of entitled intermediate securities based on the balance of underlying securities.
	Quantity *RestrictedFINDecimalNumber `xml:"Qty,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableEntitlementStatusTypeFormat4Choice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType32Choice `xml:"FrctnDspstn,omitempty"`

	// Quantity of intermediate securities awarded for a given quantity of underlying security.
	IntermediateSecuritiesToUnderlyingRatio *QuantityToQuantityRatio2 `xml:"IntrmdtSctiesToUndrlygRatio,omitempty"`

	// Last reported/known price of a financial instrument in a market.
	MarketPrice *AmountPrice4 `xml:"MktPric,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *DateFormat41Choice `xml:"XpryDt"`

	// Date of the posting (credit or debit) to the account.
	PostingDate *DateFormat41Choice `xml:"PstngDt"`

	// Period during which intermediate or outturn securities are tradable in a secondary market.
	TradingPeriod *Period4 `xml:"TradgPrd,omitempty"`

	// Balance of uninstructed position.
	UninstructedBalance *BalanceFormat7Choice `xml:"UinstdBal,omitempty"`

	// Balance of instructed position.
	InstructedBalance *BalanceFormat7Choice `xml:"InstdBal,omitempty"`
}

func (f *FinancialInstrumentAttributes73) AddSecurityIdentification() *SecurityIdentification20 {
	f.SecurityIdentification = new(SecurityIdentification20)
	return f.SecurityIdentification
}

func (f *FinancialInstrumentAttributes73) SetQuantity(value string) {
	f.Quantity = (*RestrictedFINDecimalNumber)(&value)
}

func (f *FinancialInstrumentAttributes73) AddRenounceableEntitlementStatusType() *RenounceableEntitlementStatusTypeFormat4Choice {
	f.RenounceableEntitlementStatusType = new(RenounceableEntitlementStatusTypeFormat4Choice)
	return f.RenounceableEntitlementStatusType
}

func (f *FinancialInstrumentAttributes73) AddFractionDisposition() *FractionDispositionType32Choice {
	f.FractionDisposition = new(FractionDispositionType32Choice)
	return f.FractionDisposition
}

func (f *FinancialInstrumentAttributes73) AddIntermediateSecuritiesToUnderlyingRatio() *QuantityToQuantityRatio2 {
	f.IntermediateSecuritiesToUnderlyingRatio = new(QuantityToQuantityRatio2)
	return f.IntermediateSecuritiesToUnderlyingRatio
}

func (f *FinancialInstrumentAttributes73) AddMarketPrice() *AmountPrice4 {
	f.MarketPrice = new(AmountPrice4)
	return f.MarketPrice
}

func (f *FinancialInstrumentAttributes73) AddExpiryDate() *DateFormat41Choice {
	f.ExpiryDate = new(DateFormat41Choice)
	return f.ExpiryDate
}

func (f *FinancialInstrumentAttributes73) AddPostingDate() *DateFormat41Choice {
	f.PostingDate = new(DateFormat41Choice)
	return f.PostingDate
}

func (f *FinancialInstrumentAttributes73) AddTradingPeriod() *Period4 {
	f.TradingPeriod = new(Period4)
	return f.TradingPeriod
}

func (f *FinancialInstrumentAttributes73) AddUninstructedBalance() *BalanceFormat7Choice {
	f.UninstructedBalance = new(BalanceFormat7Choice)
	return f.UninstructedBalance
}

func (f *FinancialInstrumentAttributes73) AddInstructedBalance() *BalanceFormat7Choice {
	f.InstructedBalance = new(BalanceFormat7Choice)
	return f.InstructedBalance
}
