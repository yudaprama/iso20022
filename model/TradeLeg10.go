package model

// Provides the trade leg details.
type TradeLeg10 struct {

	// Unambiguous identification of the transaction (that is the trade leg) as know by the instructing party.
	TradeLegIdentification *Max35Text `xml:"TradLegId"`

	// Reference allocated by the broker dealer.
	TradeIdentification *Max35Text `xml:"TradId,omitempty"`

	// Unique reference assigned by the trading venue when the trade is executed.
	TradeExecutionIdentification *Max35Text `xml:"TradExctnId"`

	// Identifies the order sent by the final investor to an intermediary in order to initiate a trade in the former's name. This identification is then matched with the equivalent trade by the clearing.
	OrderIdentification *Max35Text `xml:"OrdrId,omitempty"`

	// Identifies the portion of assets within a determined trade that shall be allocated to different clients.
	AllocationIdentification *Max35Text `xml:"AllcnId,omitempty"`

	// Provides the date and time of trade transaction.
	TradeDate *ISODate `xml:"TradDt"`

	// Date and time used to determine the price applicable to a trade. If the trade is registered "after market" hours, the trading price will the price of the day but the actual trade date will be the next working day.
	TransactionDateAndTime *ISODateTime `xml:"TxDtAndTm,omitempty"`

	// Provides the contractual settlement date.
	SettlementDate *DateFormat15Choice `xml:"SttlmDt"`

	// Specifies the ISO code of the trade currency.
	TradingCurrency *CurrencyCode `xml:"TradgCcy,omitempty"`

	// Identifies the trade leg indicator which gives the trade side (buy or sell).
	BuySellIndicator *Side1Code `xml:"BuySellInd"`

	// Identifies the quantity of the trade leg.
	TradeQuantity *FinancialInstrumentQuantity1Choice `xml:"TradQty"`

	// Specifies the price of the traded financial instrument.
	DealPrice *Price4 `xml:"DealPric"`

	// Principal amount of a trade (price multiplied by quantity).
	GrossAmount *AmountAndDirection21 `xml:"GrssAmt,omitempty"`

	// Place at which the security is traded.
	PlaceOfTrade *MarketIdentification84 `xml:"PlcOfTrad"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification85 `xml:"PlcOfListg,omitempty"`

	// Identifies the type of trade transaction.
	TradeType *TradeType1Code `xml:"TradTp"`

	// Indicates that the trade is for settlement of an exercised derivatives contract.
	DerivativeRelatedTrade *YesNoIndicator `xml:"DerivRltdTrad,omitempty"`

	// Party that identifies a broker when required (for example, authorised broker or prime broker).
	Broker *PartyIdentificationAndAccount100 `xml:"Brkr,omitempty"`

	// Provides the identification of the trading party.
	TradingParty *PartyIdentification35Choice `xml:"TradgPty"`

	// Indicates in which session the transaction/operation was executed by the final investor or an intermediary.
	TradeRegistrationOrigin *Max35Text `xml:"TradRegnOrgn,omitempty"`

	// Identifier of the trading participant's account at the trading venue using the venue's coding system.
	TradingPartyAccount *SecuritiesAccount19 `xml:"TradgPtyAcct,omitempty"`

	// Specifies the role of the trading party in the transaction.
	TradingCapacity *TradingCapacity5Code `xml:"TradgCpcty"`

	// Indicates how a trade is maintained in the clearing account.
	TradePostingCode *TradePosting1Code `xml:"TradPstngCd,omitempty"`

	// Place where the securities are safe-kept, physically or notionally. This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat7Choice `xml:"SfkpgPlc,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`
}

func (t *TradeLeg10) SetTradeLegIdentification(value string) {
	t.TradeLegIdentification = (*Max35Text)(&value)
}

func (t *TradeLeg10) SetTradeIdentification(value string) {
	t.TradeIdentification = (*Max35Text)(&value)
}

func (t *TradeLeg10) SetTradeExecutionIdentification(value string) {
	t.TradeExecutionIdentification = (*Max35Text)(&value)
}

func (t *TradeLeg10) SetOrderIdentification(value string) {
	t.OrderIdentification = (*Max35Text)(&value)
}

func (t *TradeLeg10) SetAllocationIdentification(value string) {
	t.AllocationIdentification = (*Max35Text)(&value)
}

func (t *TradeLeg10) SetTradeDate(value string) {
	t.TradeDate = (*ISODate)(&value)
}

func (t *TradeLeg10) SetTransactionDateAndTime(value string) {
	t.TransactionDateAndTime = (*ISODateTime)(&value)
}

func (t *TradeLeg10) AddSettlementDate() *DateFormat15Choice {
	t.SettlementDate = new(DateFormat15Choice)
	return t.SettlementDate
}

func (t *TradeLeg10) SetTradingCurrency(value string) {
	t.TradingCurrency = (*CurrencyCode)(&value)
}

func (t *TradeLeg10) SetBuySellIndicator(value string) {
	t.BuySellIndicator = (*Side1Code)(&value)
}

func (t *TradeLeg10) AddTradeQuantity() *FinancialInstrumentQuantity1Choice {
	t.TradeQuantity = new(FinancialInstrumentQuantity1Choice)
	return t.TradeQuantity
}

func (t *TradeLeg10) AddDealPrice() *Price4 {
	t.DealPrice = new(Price4)
	return t.DealPrice
}

func (t *TradeLeg10) AddGrossAmount() *AmountAndDirection21 {
	t.GrossAmount = new(AmountAndDirection21)
	return t.GrossAmount
}

func (t *TradeLeg10) AddPlaceOfTrade() *MarketIdentification84 {
	t.PlaceOfTrade = new(MarketIdentification84)
	return t.PlaceOfTrade
}

func (t *TradeLeg10) AddPlaceOfListing() *MarketIdentification85 {
	t.PlaceOfListing = new(MarketIdentification85)
	return t.PlaceOfListing
}

func (t *TradeLeg10) SetTradeType(value string) {
	t.TradeType = (*TradeType1Code)(&value)
}

func (t *TradeLeg10) SetDerivativeRelatedTrade(value string) {
	t.DerivativeRelatedTrade = (*YesNoIndicator)(&value)
}

func (t *TradeLeg10) AddBroker() *PartyIdentificationAndAccount100 {
	t.Broker = new(PartyIdentificationAndAccount100)
	return t.Broker
}

func (t *TradeLeg10) AddTradingParty() *PartyIdentification35Choice {
	t.TradingParty = new(PartyIdentification35Choice)
	return t.TradingParty
}

func (t *TradeLeg10) SetTradeRegistrationOrigin(value string) {
	t.TradeRegistrationOrigin = (*Max35Text)(&value)
}

func (t *TradeLeg10) AddTradingPartyAccount() *SecuritiesAccount19 {
	t.TradingPartyAccount = new(SecuritiesAccount19)
	return t.TradingPartyAccount
}

func (t *TradeLeg10) SetTradingCapacity(value string) {
	t.TradingCapacity = (*TradingCapacity5Code)(&value)
}

func (t *TradeLeg10) SetTradePostingCode(value string) {
	t.TradePostingCode = (*TradePosting1Code)(&value)
}

func (t *TradeLeg10) AddSafekeepingPlace() *SafekeepingPlaceFormat7Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat7Choice)
	return t.SafekeepingPlace
}

func (t *TradeLeg10) AddSafekeepingAccount() *SecuritiesAccount19 {
	t.SafekeepingAccount = new(SecuritiesAccount19)
	return t.SafekeepingAccount
}
