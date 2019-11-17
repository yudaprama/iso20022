package model

// Details of the securities trade.
type SecuritiesTradeDetails17 struct {

	// Identification of an account owner transaction that could potentially match with the allegement notified.
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

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

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt"`

	// Specifies the price of the traded financial instrument.
	// This is the deal price of the individual trade transaction.
	// If there is only one trade transaction for the execution of the trade, then the deal price could equal the executed trade price (unless, for example, the price includes commissions or rounding, or some other factor has been applied to the deal price or the executed trade price, or both).
	DealPrice *Price2 `xml:"DealPric,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes15 `xml:"FinInstrmAttrbts,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition1Choice `xml:"TradTxCond,omitempty"`

	// Specifies the type of price and information about the price.
	TypeOfPrice *TypeOfPrice3Choice `xml:"TpOfPric,omitempty"`

	// Details about the financial instrument quantity and the account involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount14 `xml:"QtyAndAcctDtls"`

	// Details of the closing of the securities financing transaction.
	SecuritiesFinancingDetails *SecuritiesFinancingTransactionDetails7 `xml:"SctiesFincgDtls,omitempty"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails25 `xml:"SttlmParams"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties11 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties11 `xml:"RcvgSttlmPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection22 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts8 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties11 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTradeDetails17) SetAccountOwnerTransactionIdentification(value string) {
	s.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails17) SetAccountServicerTransactionIdentification(value string) {
	s.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails17) SetMarketInfrastructureTransactionIdentification(value string) {
	s.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails17) SetProcessorTransactionIdentification(value string) {
	s.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails17) AddTradeIdentification(value string) {
	s.TradeIdentification = append(s.TradeIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails17) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails17) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails17) AddCollateralTransactionIdentification(value string) {
	s.CollateralTransactionIdentification = append(s.CollateralTransactionIdentification, (*Max35Text)(&value))
}

func (s *SecuritiesTradeDetails17) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesTradeDetails17) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesTradeDetails17) AddStatus() *AllegementStatus1Choice {
	s.Status = new(AllegementStatus1Choice)
	return s.Status
}

func (s *SecuritiesTradeDetails17) AddPlaceOfTrade() *MarketIdentification4 {
	s.PlaceOfTrade = new(MarketIdentification4)
	return s.PlaceOfTrade
}

func (s *SecuritiesTradeDetails17) SetPlaceOfClearing(value string) {
	s.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (s *SecuritiesTradeDetails17) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails17) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails17) AddDealPrice() *Price2 {
	s.DealPrice = new(Price2)
	return s.DealPrice
}

func (s *SecuritiesTradeDetails17) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails17) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesTradeDetails17) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes15 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes15)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesTradeDetails17) AddTradeTransactionCondition() *TradeTransactionCondition1Choice {
	newValue := new(TradeTransactionCondition1Choice)
	s.TradeTransactionCondition = append(s.TradeTransactionCondition, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails17) AddTypeOfPrice() *TypeOfPrice3Choice {
	s.TypeOfPrice = new(TypeOfPrice3Choice)
	return s.TypeOfPrice
}

func (s *SecuritiesTradeDetails17) AddQuantityAndAccountDetails() *QuantityAndAccount14 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount14)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesTradeDetails17) AddSecuritiesFinancingDetails() *SecuritiesFinancingTransactionDetails7 {
	s.SecuritiesFinancingDetails = new(SecuritiesFinancingTransactionDetails7)
	return s.SecuritiesFinancingDetails
}

func (s *SecuritiesTradeDetails17) AddSettlementParameters() *SettlementDetails25 {
	s.SettlementParameters = new(SettlementDetails25)
	return s.SettlementParameters
}

func (s *SecuritiesTradeDetails17) AddDeliveringSettlementParties() *SettlementParties11 {
	s.DeliveringSettlementParties = new(SettlementParties11)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeDetails17) AddReceivingSettlementParties() *SettlementParties11 {
	s.ReceivingSettlementParties = new(SettlementParties11)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeDetails17) AddSettlementAmount() *AmountAndDirection22 {
	s.SettlementAmount = new(AmountAndDirection22)
	return s.SettlementAmount
}

func (s *SecuritiesTradeDetails17) AddOtherAmounts() *OtherAmounts8 {
	s.OtherAmounts = new(OtherAmounts8)
	return s.OtherAmounts
}

func (s *SecuritiesTradeDetails17) AddOtherBusinessParties() *OtherParties11 {
	s.OtherBusinessParties = new(OtherParties11)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeDetails17) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
