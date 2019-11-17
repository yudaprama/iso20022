package model

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails35 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *Max35Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *Max35Text `xml:"ClsgLegId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the triparty-agent's/service-provider's point of view.
	TripartyAgentServiceProviderCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtSvcPrvdrCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`

	// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *AmountAndDirection51 `xml:"OpngSttlmAmt,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection21 `xml:"TermntnTxAmt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	OpeningSettlementDate *SettlementDate10Choice `xml:"OpngSttlmDt"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate4Choice `xml:"TermntnDt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date/time at which the sender expects settlement.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date/Time at which rate change has taken place.
	RateChangeDate *DateAndDateTimeChoice `xml:"RateChngDt,omitempty"`

	// Specifies the type of securities financing transaction, that is, repurchase agreement, reverse repurchase agreement, securities lending or securities borrowing.
	SecuritiesFinancingTransactionType *SecuritiesFinancingTransactionType2Code `xml:"SctiesFincgTxTp"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails98 `xml:"SttlmParams,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType35Choice `xml:"RateTp,omitempty"`

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
	DeliveringSettlementParties *SettlementParties36 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties36 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification99 `xml:"Invstr,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails35) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails35) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails35) SetPoolIdentification(value string) {
	s.PoolIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails35) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails35) SetTripartyAgentServiceProviderCollateralTransactionIdentification(value string) {
	s.TripartyAgentServiceProviderCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails35) SetClientTripartyCollateralTransactionIdentification(value string) {
	s.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails35) AddAccountOwner() *PartyIdentification98 {
	s.AccountOwner = new(PartyIdentification98)
	return s.AccountOwner
}

func (s *SecuritiesFinancingTransactionDetails35) AddSafekeepingAccount() *SecuritiesAccount19 {
	s.SafekeepingAccount = new(SecuritiesAccount19)
	return s.SafekeepingAccount
}

func (s *SecuritiesFinancingTransactionDetails35) AddSafekeepingPlace() *SafeKeepingPlace1 {
	s.SafekeepingPlace = new(SafeKeepingPlace1)
	return s.SafekeepingPlace
}

func (s *SecuritiesFinancingTransactionDetails35) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return s.PlaceOfTrade
}

func (s *SecuritiesFinancingTransactionDetails35) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingTransactionDetails35) AddSettlementQuantity() *Quantity6Choice {
	s.SettlementQuantity = new(Quantity6Choice)
	return s.SettlementQuantity
}

func (s *SecuritiesFinancingTransactionDetails35) AddOpeningSettlementAmount() *AmountAndDirection51 {
	s.OpeningSettlementAmount = new(AmountAndDirection51)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingTransactionDetails35) AddTerminationTransactionAmount() *AmountAndDirection21 {
	s.TerminationTransactionAmount = new(AmountAndDirection21)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails35) AddOpeningSettlementDate() *SettlementDate10Choice {
	s.OpeningSettlementDate = new(SettlementDate10Choice)
	return s.OpeningSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails35) AddTerminationDate() *TerminationDate4Choice {
	s.TerminationDate = new(TerminationDate4Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails35) AddTradeDate() *TradeDate5Choice {
	s.TradeDate = new(TradeDate5Choice)
	return s.TradeDate
}

func (s *SecuritiesFinancingTransactionDetails35) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	s.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return s.ExpectedSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails35) AddExpectedValueDate() *DateAndDateTimeChoice {
	s.ExpectedValueDate = new(DateAndDateTimeChoice)
	return s.ExpectedValueDate
}

func (s *SecuritiesFinancingTransactionDetails35) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesFinancingTransactionDetails35) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails35) SetSecuritiesFinancingTransactionType(value string) {
	s.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType2Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails35) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails35) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails35) AddSettlementParameters() *SettlementDetails98 {
	s.SettlementParameters = new(SettlementDetails98)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingTransactionDetails35) AddRateType() *RateType35Choice {
	s.RateType = new(RateType35Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails35) AddVariableRateSupport() *RateName1 {
	s.VariableRateSupport = new(RateName1)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails35) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails35) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails35) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails35) AddPricingRate() *RateOrName1Choice {
	s.PricingRate = new(RateOrName1Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails35) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails35) AddDeliveringSettlementParties() *SettlementParties36 {
	s.DeliveringSettlementParties = new(SettlementParties36)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails35) AddReceivingSettlementParties() *SettlementParties36 {
	s.ReceivingSettlementParties = new(SettlementParties36)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails35) AddInvestor() *PartyIdentification99 {
	s.Investor = new(PartyIdentification99)
	return s.Investor
}

func (s *SecuritiesFinancingTransactionDetails35) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}
