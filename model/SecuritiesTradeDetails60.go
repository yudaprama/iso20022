package model

// Details of the securities trade.
type SecuritiesTradeDetails60 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*RestrictedFINXMax16Text `xml:"CollTxId,omitempty"`

	// Identification of an account owner transaction that could potentially match with the allegement notified.
	AccountOwnerTransactionIdentification []*RestrictedFINXMax16Text `xml:"AcctOwnrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate12Choice `xml:"SttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price3 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition6Choice `xml:"TradTxCond,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice32Choice `xml:"TpOfPric,omitempty"`
}

func (s *SecuritiesTradeDetails60) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails60) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails60) AddAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = append(s.AccountOwnerTransactionIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails60) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails60) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails60) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails60) AddTradeDate() *TradeDate6Choice {
	s.TradeDate = new(TradeDate6Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails60) AddSettlementDate() *SettlementDate12Choice {
	s.SettlementDate = new(SettlementDate12Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails60) AddDealPrice() *Price3 {
	s.DealPrice = new(Price3)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails60) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails60) AddTradeTransactionCondition() *TradeTransactionCondition6Choice {
	newValue := new(TradeTransactionCondition6Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails60) AddTypeOfPrice() *TypeOfPrice32Choice {
	s.TypeOfPrice = new(TypeOfPrice32Choice)
	return s.TypeOfPrice
}
