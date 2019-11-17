package model

// Details of the securities trade.
type SecuritiesTradeDetails58 struct {

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate12Choice `xml:"SttlmDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, ie, payment is effected and securities are delivered.
	EffectiveSettlementDate *SettlementDate14Choice `xml:"FctvSttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price3 `xml:"DealPric,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting9Choice `xml:"Rptg,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition6Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity5Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator4Choice `xml:"TradOrgtrRole,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *RestrictedFINXMax350Text `xml:"FxAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails58) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails58) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails58) AddTradeDate() *TradeDate6Choice {
	s.TradeDate = new(TradeDate6Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails58) AddSettlementDate() *SettlementDate12Choice {
	s.SettlementDate = new(SettlementDate12Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails58) AddEffectiveSettlementDate() *SettlementDate14Choice {
	s.EffectiveSettlementDate = new(SettlementDate14Choice)
	return s.EffectiveSettlementDate
}

func (s *SecuritiesTradeDetails58) AddDealPrice() *Price3 {
	s.DealPrice = new(Price3)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails58) AddReporting() *Reporting9Choice {
	newValue := new(Reporting9Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails58) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails58) AddTradeTransactionCondition() *TradeTransactionCondition6Choice {
	newValue := new(TradeTransactionCondition6Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails58) AddInvestorCapacity() *InvestorCapacity5Choice {
	s.InvestorCapacity = new(InvestorCapacity5Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails58) AddTradeOriginatorRole() *TradeOriginator4Choice {
	s.TradeOriginatorRole = new(TradeOriginator4Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails58) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

func (s *SecuritiesTradeDetails58) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}
