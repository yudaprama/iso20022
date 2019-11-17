package model

// Details of the securities trade.
type SecuritiesTradeDetails38 struct {

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification78 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, ie, payment is effected and securities are delivered.
	EffectiveSettlementDate *SettlementDate3Choice `xml:"FctvSttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting2Choice `xml:"Rptg,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity1Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator1Choice `xml:"TradOrgtrRole,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *Max350Text `xml:"FxAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails38) AddPlaceOfTrade() *MarketIdentification78 {
	s.PlaceOfTrade = new(MarketIdentification78)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails38) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails38) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails38) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails38) AddEffectiveSettlementDate() *SettlementDate3Choice {
	s.EffectiveSettlementDate = new(SettlementDate3Choice)
	return s.EffectiveSettlementDate
}

func (s *SecuritiesTradeDetails38) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails38) AddReporting() *Reporting2Choice {
	newValue := new(Reporting2Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails38) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails38) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails38) AddInvestorCapacity() *InvestorCapacity1Choice {
	s.InvestorCapacity = new(InvestorCapacity1Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails38) AddTradeOriginatorRole() *TradeOriginator1Choice {
	s.TradeOriginatorRole = new(TradeOriginator1Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails38) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}

func (s *SecuritiesTradeDetails38) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*Max350Text)(&value)
}
