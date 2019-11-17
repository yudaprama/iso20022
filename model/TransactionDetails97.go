package model

// Specifies the details of the transaction.
type TransactionDetails97 struct {

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*Max35Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the triparty-agent's/service-provider's point of view.
	TripartyAgentServiceProviderCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtSvcPrvdrCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *Max35Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Unique identification assigned to the instruction by the client.
	ClientCollateralInstructionIdentification *Max35Text `xml:"ClntCollInstrId,omitempty"`

	// Unique identification assigned to the instruction by the triparty-agent/service-provider.
	TripartyAgentServiceProviderCollateralInstructionIdentification *Max35Text `xml:"TrptyAgtSvcPrvdrCollInstrId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount24 `xml:"SfkpgAcct"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace1 `xml:"SfkpgPlc,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification1 `xml:"PlcOfTrad,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Total quantity of securities to be settled.
	SettlementQuantity *Quantity6Choice `xml:"SttlmQty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection51 `xml:"SttlmAmt,omitempty"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// Date/time at which the sender expects settlement.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// For against payment transactions, the value date/time at which the Sender expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate10Choice `xml:"SttlmDt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Time stamp on when the transaction is acknowledged.
	AcknowledgedStatusTimeStamp *ISODateTime `xml:"AckdStsTmStmp,omitempty"`

	// Time stamp on when the transaction is matched.
	MatchedStatusTimeStamp *ISODateTime `xml:"MtchdStsTmStmp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails121 `xml:"SttlmParams"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties40 `xml:"RcvgSttlmPties,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties40 `xml:"DlvrgSttlmPties,omitempty"`

	// Party, either an individual or organisation, whose assets are being invested.
	Investor *PartyIdentification99 `xml:"Invstr,omitempty"`

	// Foreign Financial Institution which has been authorised by local authorities to act as account management institution in the country.
	QualifiedForeignIntermediary *PartyIdentification100 `xml:"QlfdFrgnIntrmy,omitempty"`

	// Provides additional settlement processing information which can not be included within the structured fields of the message.
	SettlementInstructionProcessingAdditionalDetails *Max350Text `xml:"SttlmInstrPrcgAddtlDtls,omitempty"`
}

func (t *TransactionDetails97) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*Max35Text)(&value))
}

func (t *TransactionDetails97) SetPoolIdentification(value string) {
	t.PoolIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails97) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails97) SetTripartyAgentServiceProviderCollateralTransactionIdentification(value string) {
	t.TripartyAgentServiceProviderCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails97) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails97) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails97) SetTripartyAgentServiceProviderCollateralInstructionIdentification(value string) {
	t.TripartyAgentServiceProviderCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails97) AddAccountOwner() *PartyIdentification98 {
	t.AccountOwner = new(PartyIdentification98)
	return t.AccountOwner
}

func (t *TransactionDetails97) AddSafekeepingAccount() *SecuritiesAccount24 {
	t.SafekeepingAccount = new(SecuritiesAccount24)
	return t.SafekeepingAccount
}

func (t *TransactionDetails97) AddSafekeepingPlace() *SafeKeepingPlace1 {
	t.SafekeepingPlace = new(SafeKeepingPlace1)
	return t.SafekeepingPlace
}

func (t *TransactionDetails97) AddPlaceOfTrade() *PlaceOfTradeIdentification1 {
	t.PlaceOfTrade = new(PlaceOfTradeIdentification1)
	return t.PlaceOfTrade
}

func (t *TransactionDetails97) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	t.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return t.PlaceOfClearing
}

func (t *TransactionDetails97) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails97) AddSettlementQuantity() *Quantity6Choice {
	t.SettlementQuantity = new(Quantity6Choice)
	return t.SettlementQuantity
}

func (t *TransactionDetails97) AddSettlementAmount() *AmountAndDirection51 {
	t.SettlementAmount = new(AmountAndDirection51)
	return t.SettlementAmount
}

func (t *TransactionDetails97) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails97) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails97) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails97) AddSettlementDate() *SettlementDate10Choice {
	t.SettlementDate = new(SettlementDate10Choice)
	return t.SettlementDate
}

func (t *TransactionDetails97) AddTradeDate() *TradeDate5Choice {
	t.TradeDate = new(TradeDate5Choice)
	return t.TradeDate
}

func (t *TransactionDetails97) SetAcknowledgedStatusTimeStamp(value string) {
	t.AcknowledgedStatusTimeStamp = (*ISODateTime)(&value)
}

func (t *TransactionDetails97) SetMatchedStatusTimeStamp(value string) {
	t.MatchedStatusTimeStamp = (*ISODateTime)(&value)
}

func (t *TransactionDetails97) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails97) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails97) AddSettlementParameters() *SettlementDetails121 {
	t.SettlementParameters = new(SettlementDetails121)
	return t.SettlementParameters
}

func (t *TransactionDetails97) AddReceivingSettlementParties() *SettlementParties40 {
	t.ReceivingSettlementParties = new(SettlementParties40)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails97) AddDeliveringSettlementParties() *SettlementParties40 {
	t.DeliveringSettlementParties = new(SettlementParties40)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails97) AddInvestor() *PartyIdentification99 {
	t.Investor = new(PartyIdentification99)
	return t.Investor
}

func (t *TransactionDetails97) AddQualifiedForeignIntermediary() *PartyIdentification100 {
	t.QualifiedForeignIntermediary = new(PartyIdentification100)
	return t.QualifiedForeignIntermediary
}

func (t *TransactionDetails97) SetSettlementInstructionProcessingAdditionalDetails(value string) {
	t.SettlementInstructionProcessingAdditionalDetails = (*Max350Text)(&value)
}
