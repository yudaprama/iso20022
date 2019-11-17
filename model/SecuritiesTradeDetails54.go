package model

// Details of the securities trade.
type SecuritiesTradeDetails54 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Identification of an account owner transaction that could potentially match with the allegement notified.
	AccountOwnerTransactionIdentification []*Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate9Choice `xml:"SttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition5Choice `xml:"TradTxCond,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice29Choice `xml:"TpOfPric,omitempty"`
}

func (s *SecuritiesTradeDetails54) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails54) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails54) AddAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = append(s.AccountOwnerTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails54) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails54) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails54) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails54) AddTradeDate() *TradeDate5Choice {
	s.TradeDate = new(TradeDate5Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails54) AddSettlementDate() *SettlementDate9Choice {
	s.SettlementDate = new(SettlementDate9Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails54) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails54) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails54) AddTradeTransactionCondition() *TradeTransactionCondition5Choice {
	newValue := new(TradeTransactionCondition5Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails54) AddTypeOfPrice() *TypeOfPrice29Choice {
	s.TypeOfPrice = new(TypeOfPrice29Choice)
	return s.TypeOfPrice
}
