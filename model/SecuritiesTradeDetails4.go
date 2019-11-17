package model

// Details of the securities trade.
type SecuritiesTradeDetails4 struct {

	// Identification of an account owner transaction that could potentially match with the allegement notified.
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unambiguous identification of a collateral transaction as assigned by the instructing party.
	CollateralTransactionIdentification []*Max35Text `xml:"CollTxId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Status of the allegement.
	Status *AllegementStatus1Choice `xml:"Sts,omitempty"`

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

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes8 `xml:"FinInstrmAttrbts,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies additional information relative to the processing of the trade.
	OpeningClosing *OpeningClosing1Choice `xml:"OpngClsg,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting1Choice `xml:"Rptg,omitempty"`

	// Details about the financial instrument quantity and the account involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount8 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *SecuritiesFinancingTransactionDetails1 `xml:"SctiesFincgDtls,omitempty"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails5 `xml:"SttlmParams"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties5 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties5 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties3 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection2 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts3 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties3 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension2 `xml:"Xtnsn,omitempty"`
}

func (s *SecuritiesTradeDetails4) SetAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails4) SetAccountServicerTransactionIdentification(value string) {
	s.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails4) SetMarketInfrastructureTransactionIdentification(value string) {
	s.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails4) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails4) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails4) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails4) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails4) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesTradeDetails4) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesTradeDetails4) AddStatus() *AllegementStatus1Choice {
	s.Status = new(AllegementStatus1Choice)
	return s.Status
}

func (s *SecuritiesTradeDetails4) AddPlaceOfTrade() *MarketIdentification4 {
	s.PlaceOfTrade = new(MarketIdentification4)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails4) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails4) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails4) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails4) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesTradeDetails4) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails4) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails4) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesTradeDetails4) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes8 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes8)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesTradeDetails4) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails4) AddOpeningClosing() *OpeningClosing1Choice {
	s.OpeningClosing = new(OpeningClosing1Choice)
	return s.OpeningClosing
}

func (s *SecuritiesTradeDetails4) AddTypeOfPrice() *TypeOfPrice3Choice {
	s.TypeOfPrice = new(TypeOfPrice3Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails4) AddReporting() *Reporting1Choice {
	newValue := new(Reporting1Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails4) AddQuantityAndAccountDetails() *QuantityAndAccount8 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount8)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesTradeDetails4) AddSecuritiesFinancingDetails() *SecuritiesFinancingTransactionDetails1 {
	s.SecuritiesFinancingDetails = new(SecuritiesFinancingTransactionDetails1)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesTradeDetails4) AddSettlementParameters() *SettlementDetails5 {
	s.SettlementParameters = new(SettlementDetails5)
	return s.SettlementParameters
}

func (s *SecuritiesTradeDetails4) AddDeliveringSettlementParties() *SettlementParties5 {
	s.DeliveringSettlementParties = new(SettlementParties5)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeDetails4) AddReceivingSettlementParties() *SettlementParties5 {
	s.ReceivingSettlementParties = new(SettlementParties5)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeDetails4) AddCashParties() *CashParties3 {
	s.CashParties = new(CashParties3)
	return s.CashParties
}

func (s *SecuritiesTradeDetails4) AddSettlementAmount() *AmountAndDirection2 {
	s.SettlementAmount = new(AmountAndDirection2)
	return s.SettlementAmount
}

func (s *SecuritiesTradeDetails4) AddOtherAmounts() *OtherAmounts3 {
	s.OtherAmounts = new(OtherAmounts3)
	return s.OtherAmounts
}

func (s *SecuritiesTradeDetails4) AddOtherBusinessParties() *OtherParties3 {
	s.OtherBusinessParties = new(OtherParties3)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeDetails4) AddExtension() *Extension2 {
	newValue := new(Extension2)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
