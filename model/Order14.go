package model

// Intention to transfer an ownership of a financial instrument.
type Order14 struct {

	// Specifies the type of business process.
	BusinessProcessType *BusinessProcessType1Choice `xml:"BizPrcTp,omitempty"`

	// Unique identifier for Order as assigned by sell-side.
	OrderIdentification []*Max35Text `xml:"OrdrId,omitempty"`

	// Unique identifier for the order as assigned by the buy-side. Uniqueness must be guaranteed within a single trading day. Firms, particularly those that electronically submit multi-day orders, trade globally or throughout market close periods, should ensure uniqueness across days, for example by embedding a date within the ClientOrderIdentification element.
	ClientOrderIdentification []*Max35Text `xml:"ClntOrdrId,omitempty"`

	// Assigned by the party that originates the order. Can be used to provide the ClientOrderIdentification used by an exchange or executing system.
	SecondaryClientOrderIdentification []*Max35Text `xml:"ScndryClntOrdrId,omitempty"`

	// Unique identifier for a list, as assigned by the trading party. The identifier must be unique within a single trading day.
	ListIdentification []*Max35Text `xml:"ListId,omitempty"`

	// Coded list to specify the side of the order.
	Side *Side3Code `xml:"Sd"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt,omitempty"`

	// Specifies the type of transaction of which the order is a component.
	TradeTransactionType *TradeType3Choice `xml:"TradTxTp,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TradeTransactionCondition []*TradeTransactionCondition4Choice `xml:"TradTxCond,omitempty"`

	// Transaction is a pre-advice, that is, for matching purposes only.
	PreAdvice *YesNoIndicator `xml:"PreAdvc,omitempty"`

	// Market in which a trade transaction is to be or has been executed.
	PlaceOfTrade *MarketIdentification77 `xml:"PlcOfTrad,omitempty"`

	// Date and time when an entry is posted to an account on the account servicer's books.
	OrderBookingDate *ISODate `xml:"OrdrBookgDt,omitempty"`

	// Indicates the date and time of the agreement in principal between counter-parties prior to actual trade date.
	// Used with fixed income for municipal new issue markets.
	TradeOriginationDate *ISODateTime `xml:"TradOrgtnDt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate4Choice `xml:"TradDt"`

	// Processing date of the trading session.
	ProcessingDate *TradeDate4Choice `xml:"PrcgDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate8Choice `xml:"SttlmDt"`

	// Valuation point, or valuation date of the portfolio (underlying assets). This is also known as price date.
	NAVDate *DateAndDateTime1Choice `xml:"NAVDt,omitempty"`

	// Quantity of financial instrument bought or sold which is less than the quantity of financial instrument ordered.
	PartialFillDetails []*PartialFill1 `xml:"PrtlFillDtls,omitempty"`

	// Quantity of financial instrument that is being confirmed for the account.The quantity of the security to be settled.
	ConfirmationQuantity *Quantity6Choice `xml:"ConfQty"`

	// Breakdown of a quantity into lots such as tax lots, instrument series, etc.
	QuantityBreakdown []*QuantityBreakdown11 `xml:"QtyBrkdwn,omitempty"`

	// Principal amount of a trade (price multiplied by quantity).
	GrossTradeAmount *AmountAndDirection29 `xml:"GrssTradAmt,omitempty"`

	// Amount of money for which goods or services are offered, sold, or bought.
	DealPrice *Price4 `xml:"DealPric"`

	// Specifies the type of transaction price.
	TypeOfPrice *TypeOfPrice10Choice `xml:"TpOfPric,omitempty"`

	// Identifies whether an order is a margin order or a non-margin order. This is primarily used when sending orders to Japanese exchanges to indicate sell margin or buy to cover. The same tag could be assigned also by buy-side to indicate the intent to sell or buy margin and the sell-side to accept or reject (base on some validation criteria) the margin request.
	CashMargin *CashMarginOrder1Code `xml:"CshMrgn,omitempty"`

	// Amount of money due to a party as compensation for a service.
	Commission *Commission16 `xml:"Comssn,omitempty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Specifies the number of days from trade date that the counterparty on the other side of the trade should be "given up" or divulged.
	GiveUpNumberOfDays *Max3Number `xml:"GvUpNbOfDays,omitempty"`

	// Indicates whether the trade is cum interest or ex interest.
	InterestType *InterestType2Code `xml:"IntrstTp,omitempty"`

	// Interest amount that has accrued in between two periods, for example, in between interest payment periods.
	AccruedInterestAmount *AmountAndDirection29 `xml:"AcrdIntrstAmt,omitempty"`

	// Interest rate that has been accrued in between coupon payment periods.
	AccruedInterestPercentage *PercentageRate `xml:"AcrdIntrstPctg,omitempty"`

	// Specifies the regulatory conditions type of the trade.
	TradeRegulatoryConditionsType *TradeRegulatoryConditions1Code `xml:"TradRgltryCondsTp,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Counterparties eligibility as defined by article 24 of the EU MiFID Directive applicable to transactions executed by investment firms for eligible counterparties.
	OrderOriginatorEligibility *Eligibility1Code `xml:"OrdrOrgtrElgblty,omitempty"`

	// Indicates whether the resulting position after a  trade should be an opening position or closing position. Used for omnibus accounting - where accounts are held on a gross basis instead of being netted together.
	PositionEffect *PositionEffect2Code `xml:"PosFct,omitempty"`

	// Indicates whether the derivative product is covered or not by an underlying financial instrument position.
	DerivativeCovered *YesNoIndicator `xml:"DerivCvrd,omitempty"`

	// Type of charge/tax basis.
	ChargeTaxBasisType *ChargeTaxBasisType1Choice `xml:"ChrgTaxBsisTp,omitempty"`

	// Specifies the type of capital gain.
	CapitalGainType *EUCapitalGainType2Choice `xml:"CptlGnTp,omitempty"`

	// Provides the matching status of the trade confirmation.
	MatchStatus *MatchingStatus8Choice `xml:"MtchSts,omitempty"`

	// Specifies the type of pay-in call report.
	CallInType *CallIn1Code `xml:"CallInTp,omitempty"`

	// Type of yield at which the transaction was effected.
	YieldType *YieldCalculation2 `xml:"YldTp,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting5Choice `xml:"Rptg,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters3 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Provides additional details of the trade process not included within structured fields of this message.
	AdditionalTradeInstructionProcessingInformation *Max350Text `xml:"AddtlTradInstrPrcgInf,omitempty"`
}

func (o *Order14) AddBusinessProcessType() *BusinessProcessType1Choice {
	o.BusinessProcessType = new(BusinessProcessType1Choice)
	return o.BusinessProcessType
}

func (o *Order14) AddOrderIdentification(value string) {
	o.OrderIdentification = append(o.OrderIdentification, (*Max35Text)(&value))
}

func (o *Order14) AddClientOrderIdentification(value string) {
	o.ClientOrderIdentification = append(o.ClientOrderIdentification, (*Max35Text)(&value))
}

func (o *Order14) AddSecondaryClientOrderIdentification(value string) {
	o.SecondaryClientOrderIdentification = append(o.SecondaryClientOrderIdentification, (*Max35Text)(&value))
}

func (o *Order14) AddListIdentification(value string) {
	o.ListIdentification = append(o.ListIdentification, (*Max35Text)(&value))
}

func (o *Order14) SetSide(value string) {
	o.Side = (*Side3Code)(&value)
}

func (o *Order14) SetPayment(value string) {
	o.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (o *Order14) AddTradeTransactionType() *TradeType3Choice {
	o.TradeTransactionType = new(TradeType3Choice)
	return o.TradeTransactionType
}

func (o *Order14) AddTradeTransactionCondition() *TradeTransactionCondition4Choice {
	newValue := new(TradeTransactionCondition4Choice)
	o.TradeTransactionCondition = append(o.TradeTransactionCondition, newValue)
	return newValue
}

func (o *Order14) SetPreAdvice(value string) {
	o.PreAdvice = (*YesNoIndicator)(&value)
}

func (o *Order14) AddPlaceOfTrade() *MarketIdentification77 {
	o.PlaceOfTrade = new(MarketIdentification77)
	return o.PlaceOfTrade
}

func (o *Order14) SetOrderBookingDate(value string) {
	o.OrderBookingDate = (*ISODate)(&value)
}

func (o *Order14) SetTradeOriginationDate(value string) {
	o.TradeOriginationDate = (*ISODateTime)(&value)
}

func (o *Order14) AddTradeDate() *TradeDate4Choice {
	o.TradeDate = new(TradeDate4Choice)
	return o.TradeDate
}

func (o *Order14) AddProcessingDate() *TradeDate4Choice {
	o.ProcessingDate = new(TradeDate4Choice)
	return o.ProcessingDate
}

func (o *Order14) AddSettlementDate() *SettlementDate8Choice {
	o.SettlementDate = new(SettlementDate8Choice)
	return o.SettlementDate
}

func (o *Order14) AddNAVDate() *DateAndDateTime1Choice {
	o.NAVDate = new(DateAndDateTime1Choice)
	return o.NAVDate
}

func (o *Order14) AddPartialFillDetails() *PartialFill1 {
	newValue := new(PartialFill1)
	o.PartialFillDetails = append(o.PartialFillDetails, newValue)
	return newValue
}

func (o *Order14) AddConfirmationQuantity() *Quantity6Choice {
	o.ConfirmationQuantity = new(Quantity6Choice)
	return o.ConfirmationQuantity
}

func (o *Order14) AddQuantityBreakdown() *QuantityBreakdown11 {
	newValue := new(QuantityBreakdown11)
	o.QuantityBreakdown = append(o.QuantityBreakdown, newValue)
	return newValue
}

func (o *Order14) AddGrossTradeAmount() *AmountAndDirection29 {
	o.GrossTradeAmount = new(AmountAndDirection29)
	return o.GrossTradeAmount
}

func (o *Order14) AddDealPrice() *Price4 {
	o.DealPrice = new(Price4)
	return o.DealPrice
}

func (o *Order14) AddTypeOfPrice() *TypeOfPrice10Choice {
	o.TypeOfPrice = new(TypeOfPrice10Choice)
	return o.TypeOfPrice
}

func (o *Order14) SetCashMargin(value string) {
	o.CashMargin = (*CashMarginOrder1Code)(&value)
}

func (o *Order14) AddCommission() *Commission16 {
	o.Commission = new(Commission16)
	return o.Commission
}

func (o *Order14) SetNumberOfDaysAccrued(value string) {
	o.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (o *Order14) SetGiveUpNumberOfDays(value string) {
	o.GiveUpNumberOfDays = (*Max3Number)(&value)
}

func (o *Order14) SetInterestType(value string) {
	o.InterestType = (*InterestType2Code)(&value)
}

func (o *Order14) AddAccruedInterestAmount() *AmountAndDirection29 {
	o.AccruedInterestAmount = new(AmountAndDirection29)
	return o.AccruedInterestAmount
}

func (o *Order14) SetAccruedInterestPercentage(value string) {
	o.AccruedInterestPercentage = (*PercentageRate)(&value)
}

func (o *Order14) SetTradeRegulatoryConditionsType(value string) {
	o.TradeRegulatoryConditionsType = (*TradeRegulatoryConditions1Code)(&value)
}

func (o *Order14) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	o.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return o.CurrencyToBuyOrSell
}

func (o *Order14) SetOrderOriginatorEligibility(value string) {
	o.OrderOriginatorEligibility = (*Eligibility1Code)(&value)
}

func (o *Order14) SetPositionEffect(value string) {
	o.PositionEffect = (*PositionEffect2Code)(&value)
}

func (o *Order14) SetDerivativeCovered(value string) {
	o.DerivativeCovered = (*YesNoIndicator)(&value)
}

func (o *Order14) AddChargeTaxBasisType() *ChargeTaxBasisType1Choice {
	o.ChargeTaxBasisType = new(ChargeTaxBasisType1Choice)
	return o.ChargeTaxBasisType
}

func (o *Order14) AddCapitalGainType() *EUCapitalGainType2Choice {
	o.CapitalGainType = new(EUCapitalGainType2Choice)
	return o.CapitalGainType
}

func (o *Order14) AddMatchStatus() *MatchingStatus8Choice {
	o.MatchStatus = new(MatchingStatus8Choice)
	return o.MatchStatus
}

func (o *Order14) SetCallInType(value string) {
	o.CallInType = (*CallIn1Code)(&value)
}

func (o *Order14) AddYieldType() *YieldCalculation2 {
	o.YieldType = new(YieldCalculation2)
	return o.YieldType
}

func (o *Order14) AddReporting() *Reporting5Choice {
	newValue := new(Reporting5Choice)
	o.Reporting = append(o.Reporting, newValue)
	return newValue
}

func (o *Order14) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters3 {
	o.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters3)
	return o.AdditionalPhysicalOrRegistrationDetails
}

func (o *Order14) SetAdditionalTradeInstructionProcessingInformation(value string) {
	o.AdditionalTradeInstructionProcessingInformation = (*Max350Text)(&value)
}
