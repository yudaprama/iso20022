package model

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails8 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *Max35Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *Max35Text `xml:"ClsgLegId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *AmountAndDirection8 `xml:"OpngSttlmAmt,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection4 `xml:"TermntnTxAmt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	OpeningSettlementDate *SettlementDate2Choice `xml:"OpngSttlmDt"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate2Choice `xml:"TermntnDt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType1Code `xml:"SctiesFincgTxTp"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails29 `xml:"SttlmParams,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType5Choice `xml:"RateTp,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName1 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName1Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties11 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties11 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification37Choice `xml:"Invstr,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails8) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails8) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails8) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails8) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails8) SetTripartyAgentCollateralTransactionIdentification(value string) {
	s.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails8) SetClientTripartyCollateralTransactionIdentification(value string) {
	s.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails8) AddAccountOwner() *PartyIdentification36Choice {
	s.AccountOwner = new(PartyIdentification36Choice)
	return s.AccountOwner
}

func (s *SecuritiesFinancingTransactionDetails8) AddSafekeepingAccount() *SecuritiesAccount13 {
	s.SafekeepingAccount = new(SecuritiesAccount13)
	return s.SafekeepingAccount
}

func (s *SecuritiesFinancingTransactionDetails8) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesFinancingTransactionDetails8) AddPlaceOfTrade() *MarketIdentification4 {
	s.PlaceOfTrade = new(MarketIdentification4)
	return s.PlaceOfTrade
}

func (s *SecuritiesFinancingTransactionDetails8) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingTransactionDetails8) AddSettlementQuantity() *Quantity6Choice {
	s.SettlementQuantity = new(Quantity6Choice)
	return s.SettlementQuantity
}

func (s *SecuritiesFinancingTransactionDetails8) AddOpeningSettlementAmount() *AmountAndDirection8 {
	s.OpeningSettlementAmount = new(AmountAndDirection8)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingTransactionDetails8) AddTerminationTransactionAmount() *AmountAndDirection4 {
	s.TerminationTransactionAmount = new(AmountAndDirection4)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails8) AddOpeningSettlementDate() *SettlementDate2Choice {
	s.OpeningSettlementDate = new(SettlementDate2Choice)
	return s.OpeningSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails8) AddTerminationDate() *TerminationDate2Choice {
	s.TerminationDate = new(TerminationDate2Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails8) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesFinancingTransactionDetails8) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	s.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return s.ExpectedSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails8) AddExpectedValueDate() *DateAndDateTimeChoice {
	s.ExpectedValueDate = new(DateAndDateTimeChoice)
	return s.ExpectedValueDate
}

func (s *SecuritiesFinancingTransactionDetails8) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesFinancingTransactionDetails8) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails8) SetSecuritiesFinancingTransactionType(value string) {
	s.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType1Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails8) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails8) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails8) AddSettlementParameters() *SettlementDetails29 {
	s.SettlementParameters = new(SettlementDetails29)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingTransactionDetails8) AddRateType() *RateType5Choice {
	s.RateType = new(RateType5Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails8) AddVariableRateSupport() *RateName1 {
	s.VariableRateSupport = new(RateName1)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails8) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails8) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails8) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails8) AddPricingRate() *RateOrName1Choice {
	s.PricingRate = new(RateOrName1Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails8) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails8) AddDeliveringSettlementParties() *SettlementParties11 {
	s.DeliveringSettlementParties = new(SettlementParties11)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails8) AddReceivingSettlementParties() *SettlementParties11 {
	s.ReceivingSettlementParties = new(SettlementParties11)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails8) AddInvestor() *PartyIdentification37Choice {
	s.Investor = new(PartyIdentification37Choice)
	return s.Investor
}

func (s *SecuritiesFinancingTransactionDetails8) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}
