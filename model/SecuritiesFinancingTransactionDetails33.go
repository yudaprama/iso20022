package model

// Details of the closing of the securities financing transaction.
type SecuritiesFinancingTransactionDetails33 struct {

	// Unambiguous identification of the underlying securities financing trade as assigned by the instructing party. The identification is common to all collateral pieces (one or many).
	SecuritiesFinancingTradeIdentification *RestrictedFINXMax16Text `xml:"SctiesFincgTradId,omitempty"`

	// Unambiguous identification of the second leg of the transaction as known by the account owner (or the instructing party acting on its behalf).
	ClosingLegIdentification *RestrictedFINXMax16Text `xml:"ClsgLegId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Context, or geographic environment, in which trading parties may meet in order to negotiate and execute trades among themselves.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity10Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities at the opening of a securities financing transaction.
	OpeningSettlementAmount *AmountAndDirection67 `xml:"OpngSttlmAmt,omitempty"`

	// Total amount of money to be settled to terminate the transaction.
	TerminationTransactionAmount *AmountAndDirection59 `xml:"TermntnTxAmt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	OpeningSettlementDate *SettlementDate15Choice `xml:"OpngSttlmDt"`

	// Closing date/time or maturity date/time of the transaction.
	TerminationDate *TerminationDate5Choice `xml:"TermntnDt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date/time at which the sender expects settlement.
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
	SettlementParameters *SettlementDetails107 `xml:"SttlmParams,omitempty"`

	// Specifies whether the rate is fixed or variable.
	RateType *RateType67Choice `xml:"RateTp,omitempty"`

	// Index or support rate used together with the spread to calculate the
	// repurchase rate.
	VariableRateSupport *RateName2 `xml:"VarblRateSpprt,omitempty"`

	// Rate to be used to recalculate the repurchase amount.
	RepurchaseRate *Rate2 `xml:"RpRate,omitempty"`

	// Percentage mark-up on a loan consideration used to reflect the lender's risk.
	StockLoanMargin *Rate2 `xml:"StockLnMrgn,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	SecuritiesHaircut *Rate2 `xml:"SctiesHrcut,omitempty"`

	// Interest rate to be paid on the transaction amount, as agreed between the counterparties.
	PricingRate *RateOrName2Choice `xml:"PricgRate,omitempty"`

	// Repurchase spread expressed as a rate; margin over or under an index that determines the repurchase rate.
	Spread *Rate2 `xml:"Sprd,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties44 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties44 `xml:"RcvgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification110 `xml:"Invstr,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *RestrictedFINXMax350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (s *SecuritiesFinancingTransactionDetails33) SetSecuritiesFinancingTradeIdentification(value string) {
	s.SecuritiesFinancingTradeIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails33) SetClosingLegIdentification(value string) {
	s.ClosingLegIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails33) SetPoolIdentification(value string) {
	s.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails33) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails33) SetTripartyAgentCollateralTransactionIdentification(value string) {
	s.TripartyAgentCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails33) SetClientTripartyCollateralTransactionIdentification(value string) {
	s.ClientTripartyCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesFinancingTransactionDetails33) AddAccountOwner() *PartyIdentification109 {
	s.AccountOwner = new(PartyIdentification109)
	return s.AccountOwner
}

func (s *SecuritiesFinancingTransactionDetails33) AddSafekeepingAccount() *SecuritiesAccount30 {
	s.SafekeepingAccount = new(SecuritiesAccount30)
	return s.SafekeepingAccount
}

func (s *SecuritiesFinancingTransactionDetails33) AddSafekeepingPlace() *SafeKeepingPlace2 {
	s.SafekeepingPlace = new(SafeKeepingPlace2)
	return s.SafekeepingPlace
}

func (s *SecuritiesFinancingTransactionDetails33) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return s.PlaceOfTrade
}

func (s *SecuritiesFinancingTransactionDetails33) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesFinancingTransactionDetails33) AddSettlementQuantity() *Quantity10Choice {
	s.SettlementQuantity = new(Quantity10Choice)
	return s.SettlementQuantity
}

func (s *SecuritiesFinancingTransactionDetails33) AddOpeningSettlementAmount() *AmountAndDirection67 {
	s.OpeningSettlementAmount = new(AmountAndDirection67)
	return s.OpeningSettlementAmount
}

func (s *SecuritiesFinancingTransactionDetails33) AddTerminationTransactionAmount() *AmountAndDirection59 {
	s.TerminationTransactionAmount = new(AmountAndDirection59)
	return s.TerminationTransactionAmount
}

func (s *SecuritiesFinancingTransactionDetails33) AddOpeningSettlementDate() *SettlementDate15Choice {
	s.OpeningSettlementDate = new(SettlementDate15Choice)
	return s.OpeningSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails33) AddTerminationDate() *TerminationDate5Choice {
	s.TerminationDate = new(TerminationDate5Choice)
	return s.TerminationDate
}

func (s *SecuritiesFinancingTransactionDetails33) AddTradeDate() *TradeDate6Choice {
	s.TradeDate = new(TradeDate6Choice)
	return s.TradeDate
}

func (s *SecuritiesFinancingTransactionDetails33) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	s.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return s.ExpectedSettlementDate
}

func (s *SecuritiesFinancingTransactionDetails33) AddExpectedValueDate() *DateAndDateTimeChoice {
	s.ExpectedValueDate = new(DateAndDateTimeChoice)
	return s.ExpectedValueDate
}

func (s *SecuritiesFinancingTransactionDetails33) AddLateDeliveryDate() *DateAndDateTimeChoice {
	s.LateDeliveryDate = new(DateAndDateTimeChoice)
	return s.LateDeliveryDate
}

func (s *SecuritiesFinancingTransactionDetails33) AddRateChangeDate() *DateAndDateTimeChoice {
	s.RateChangeDate = new(DateAndDateTimeChoice)
	return s.RateChangeDate
}

func (s *SecuritiesFinancingTransactionDetails33) SetSecuritiesFinancingTransactionType(value string) {
	s.SecuritiesFinancingTransactionType = (*SecuritiesFinancingTransactionType1Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails33) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails33) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesFinancingTransactionDetails33) AddSettlementParameters() *SettlementDetails107 {
	s.SettlementParameters = new(SettlementDetails107)
	return s.SettlementParameters
}

func (s *SecuritiesFinancingTransactionDetails33) AddRateType() *RateType67Choice {
	s.RateType = new(RateType67Choice)
	return s.RateType
}

func (s *SecuritiesFinancingTransactionDetails33) AddVariableRateSupport() *RateName2 {
	s.VariableRateSupport = new(RateName2)
	return s.VariableRateSupport
}

func (s *SecuritiesFinancingTransactionDetails33) AddRepurchaseRate() *Rate2 {
	s.RepurchaseRate = new(Rate2)
	return s.RepurchaseRate
}

func (s *SecuritiesFinancingTransactionDetails33) AddStockLoanMargin() *Rate2 {
	s.StockLoanMargin = new(Rate2)
	return s.StockLoanMargin
}

func (s *SecuritiesFinancingTransactionDetails33) AddSecuritiesHaircut() *Rate2 {
	s.SecuritiesHaircut = new(Rate2)
	return s.SecuritiesHaircut
}

func (s *SecuritiesFinancingTransactionDetails33) AddPricingRate() *RateOrName2Choice {
	s.PricingRate = new(RateOrName2Choice)
	return s.PricingRate
}

func (s *SecuritiesFinancingTransactionDetails33) AddSpread() *Rate2 {
	s.Spread = new(Rate2)
	return s.Spread
}

func (s *SecuritiesFinancingTransactionDetails33) AddDeliveringSettlementParties() *SettlementParties44 {
	s.DeliveringSettlementParties = new(SettlementParties44)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails33) AddReceivingSettlementParties() *SettlementParties44 {
	s.ReceivingSettlementParties = new(SettlementParties44)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesFinancingTransactionDetails33) AddInvestor() *PartyIdentification110 {
	s.Investor = new(PartyIdentification110)
	return s.Investor
}

func (s *SecuritiesFinancingTransactionDetails33) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	s.SettlementInstructionProcessingAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}
