package model

// Details of the securities trade.
type SecuritiesTradeDetails66 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*RestrictedFINXMax16Text `xml:"CollTxId,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate12Choice `xml:"SttlmDt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price3 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing4Choice `xml:"OpngClsg,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting9Choice `xml:"Rptg,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition6Choice `xml:"TradTxCond,omitempty"`

	// Specifies the role of the investor in the transaction.
	InvestorCapacity *InvestorCapacity5Choice `xml:"InvstrCpcty,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradeOriginatorRole *TradeOriginator4Choice `xml:"TradOrgtrRole,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice32Choice `xml:"TpOfPric,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Provides the matching status of the instruction.
	MatchingStatus *MatchingStatus28Choice `xml:"MtchgSts,omitempty"`

	// Status of affirmation of a trade.
	AffirmationStatus *AffirmationStatus9Choice `xml:"AffirmSts,omitempty"`

	// Provides additional details pertaining to foreign exchange instructions.
	FXAdditionalDetails *RestrictedFINXMax350Text `xml:"FxAddtlDtls,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesTradeDetails66) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails66) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (s *SecuritiesTradeDetails66) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails66) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	s.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return s.PlaceOfClearing
}

func (s *SecuritiesTradeDetails66) AddTradeDate() *TradeDate6Choice {
	s.TradeDate = new(TradeDate6Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails66) AddSettlementDate() *SettlementDate12Choice {
	s.SettlementDate = new(SettlementDate12Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails66) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesTradeDetails66) AddDealPrice() *Price3 {
	s.DealPrice = new(Price3)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails66) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails66) AddOpeningClosing() *OpeningClosing4Choice {
	s.OpeningClosing = new(OpeningClosing4Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails66) AddReporting() *Reporting9Choice {
	newValue := new(Reporting9Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails66) AddTradeTransactionCondition() *TradeTransactionCondition6Choice {
	newValue := new(TradeTransactionCondition6Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails66) AddInvestorCapacity() *InvestorCapacity5Choice {
	s.InvestorCapacity = new(InvestorCapacity5Choice)
	return s.InvestorCapacity
}

func (s *SecuritiesTradeDetails66) AddTradeOriginatorRole() *TradeOriginator4Choice {
	s.TradeOriginatorRole = new(TradeOriginator4Choice)
	return s.TradeOriginatorRole
}

func (s *SecuritiesTradeDetails66) AddTypeOfPrice() *TypeOfPrice32Choice {
	s.TypeOfPrice = new(TypeOfPrice32Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails66) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SecuritiesTradeDetails66) AddMatchingStatus() *MatchingStatus28Choice {
	s.MatchingStatus = new(MatchingStatus28Choice)
	return s.MatchingStatus
}

func (s *SecuritiesTradeDetails66) AddAffirmationStatus() *AffirmationStatus9Choice {
	s.AffirmationStatus = new(AffirmationStatus9Choice)
	return s.AffirmationStatus
}

func (s *SecuritiesTradeDetails66) SetFXAdditionalDetails(value string) {
	s.FXAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

func (s *SecuritiesTradeDetails66) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}
