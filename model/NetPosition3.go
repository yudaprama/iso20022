package model

// Provides the net positions details.
type NetPosition3 struct {

	// Identifies the clearing member account at the Central counterparty through which the trade must be cleared (sometimes called position account).
	ClearingAccount *SecuritiesAccount18 `xml:"ClrAcct"`

	// Provides the identification for the non-clearing member.
	NonClearingMember *PartyIdentificationAndAccount31 `xml:"NonClrMmb,omitempty"`

	// An account opened by the central counterparty in the name of the clearing member or its settlement agent within the account structure, for settlement purposes (gives information about the clearing member/its settlement agent account at the central securities depository).
	DeliveryAccount *SecuritiesAccount19 `xml:"DlvryAcct,omitempty"`

	// Provides details about the security identification.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Provides the initial position amount.
	InitialPositionAmount *AmountAndDirection21 `xml:"InitlPosAmt,omitempty"`

	// Provides the net position amount.
	NetPositionAmount *AmountAndDirection21 `xml:"NetPosAmt"`

	// Interest that has accumulated on a bond since the last interest payment up to, but not including, the settlement date.
	AccruedInterestAmount *AmountAndDirection21 `xml:"AcrdIntrstAmt,omitempty"`

	// This is the price of the trade.
	AverageDealPrice *Price4 `xml:"AvrgDealPric,omitempty"`

	// Identifies the quantity of the trade leg.
	NetQuantity *FinancialInstrumentQuantity1Choice `xml:"NetQty"`

	// Indicates the securities movement direction, that is, whether this is a delivery or return.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Place at which a trade settles.
	Depository *PartyIdentification34Choice `xml:"Dpstry"`

	// Identifies the trading capacity of the seller.
	TradingCapacity *TradingCapacity5Code `xml:"TradgCpcty,omitempty"`

	// Place at which the security is traded.
	PlaceOfTrade *MarketIdentification20 `xml:"PlcOfTrad,omitempty"`

	// Provides the date of the trade.
	TradeDate *ISODate `xml:"TradDt,omitempty"`

	// Provides the contractual settlement date.
	SettlementDate *DateFormat15Choice `xml:"SttlmDt,omitempty"`

	// Provides the trade leg details such as trade leg identification and trade type.
	TradeLegDetails []*TradeLeg10 `xml:"TradLegDtls,omitempty"`
}

func (n *NetPosition3) AddClearingAccount() *SecuritiesAccount18 {
	n.ClearingAccount = new(SecuritiesAccount18)
	return n.ClearingAccount
}

func (n *NetPosition3) AddNonClearingMember() *PartyIdentificationAndAccount31 {
	n.NonClearingMember = new(PartyIdentificationAndAccount31)
	return n.NonClearingMember
}

func (n *NetPosition3) AddDeliveryAccount() *SecuritiesAccount19 {
	n.DeliveryAccount = new(SecuritiesAccount19)
	return n.DeliveryAccount
}

func (n *NetPosition3) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	n.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return n.FinancialInstrumentIdentification
}

func (n *NetPosition3) AddInitialPositionAmount() *AmountAndDirection21 {
	n.InitialPositionAmount = new(AmountAndDirection21)
	return n.InitialPositionAmount
}

func (n *NetPosition3) AddNetPositionAmount() *AmountAndDirection21 {
	n.NetPositionAmount = new(AmountAndDirection21)
	return n.NetPositionAmount
}

func (n *NetPosition3) AddAccruedInterestAmount() *AmountAndDirection21 {
	n.AccruedInterestAmount = new(AmountAndDirection21)
	return n.AccruedInterestAmount
}

func (n *NetPosition3) AddAverageDealPrice() *Price4 {
	n.AverageDealPrice = new(Price4)
	return n.AverageDealPrice
}

func (n *NetPosition3) AddNetQuantity() *FinancialInstrumentQuantity1Choice {
	n.NetQuantity = new(FinancialInstrumentQuantity1Choice)
	return n.NetQuantity
}

func (n *NetPosition3) SetSecuritiesMovementType(value string) {
	n.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (n *NetPosition3) AddDepository() *PartyIdentification34Choice {
	n.Depository = new(PartyIdentification34Choice)
	return n.Depository
}

func (n *NetPosition3) SetTradingCapacity(value string) {
	n.TradingCapacity = (*TradingCapacity5Code)(&value)
}

func (n *NetPosition3) AddPlaceOfTrade() *MarketIdentification20 {
	n.PlaceOfTrade = new(MarketIdentification20)
	return n.PlaceOfTrade
}

func (n *NetPosition3) SetTradeDate(value string) {
	n.TradeDate = (*ISODate)(&value)
}

func (n *NetPosition3) AddSettlementDate() *DateFormat15Choice {
	n.SettlementDate = new(DateFormat15Choice)
	return n.SettlementDate
}

func (n *NetPosition3) AddTradeLegDetails() *TradeLeg10 {
	newValue := new(TradeLeg10)
	n.TradeLegDetails = append(n.TradeLegDetails, newValue)
	return newValue
}
