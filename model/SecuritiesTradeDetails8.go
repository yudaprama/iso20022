package model

// Details of the securities trade.
type SecuritiesTradeDetails8 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Identification of an account owner transaction that could potentially match with the allegement notified.
	AccountOwnerTransactionIdentification []*Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing1Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting1Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`
}

func (s *SecuritiesTradeDetails8) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails8) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails8) AddAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = append(s.AccountOwnerTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails8) AddPlaceOfTrade() *MarketIdentification4 {
	s.PlaceOfTrade = new(MarketIdentification4)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails8) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails8) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails8) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails8) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesTradeDetails8) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails8) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails8) AddOpeningClosing() *OpeningClosing1Choice {
	s.OpeningClosing = new(OpeningClosing1Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails8) AddReporting() *Reporting1Choice {
	newValue := new(Reporting1Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails8) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails8) AddTypeOfPrice() *TypeOfPrice3Choice {
	s.TypeOfPrice = new(TypeOfPrice3Choice)
	return s.TypeOfPrice
}
