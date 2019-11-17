package model

// Identifies the details of the transaction.
type TransactionDetails90 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity4Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent15Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails116 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity10Choice `xml:"PstngQty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection67 `xml:"PstngAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date/time at which the sender expects settlement.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate15Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties49 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties49 `xml:"RcvgSttlmPties,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *RestrictedFINXMax350Text `xml:"TxAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *TransactionDetails90) AddTransactionActivity() *TransactionActivity4Choice {
	t.TransactionActivity = new(TransactionActivity4Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails90) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent15Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent15Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails90) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails90) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails90) AddSettlementParameters() *SettlementDetails116 {
	t.SettlementParameters = new(SettlementDetails116)
	return t.SettlementParameters
}

func (t *TransactionDetails90) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	t.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return t.PlaceOfTrade
}

func (t *TransactionDetails90) AddSafekeepingPlace() *SafeKeepingPlace2 {
	t.SafekeepingPlace = new(SafeKeepingPlace2)
	return t.SafekeepingPlace
}

func (t *TransactionDetails90) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	t.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return t.PlaceOfClearing
}

func (t *TransactionDetails90) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails90) AddPostingQuantity() *Quantity10Choice {
	t.PostingQuantity = new(Quantity10Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails90) AddPostingAmount() *AmountAndDirection67 {
	t.PostingAmount = new(AmountAndDirection67)
	return t.PostingAmount
}

func (t *TransactionDetails90) AddTradeDate() *TradeDate6Choice {
	t.TradeDate = new(TradeDate6Choice)
	return t.TradeDate
}

func (t *TransactionDetails90) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails90) AddSettlementDate() *SettlementDate15Choice {
	t.SettlementDate = new(SettlementDate15Choice)
	return t.SettlementDate
}

func (t *TransactionDetails90) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails90) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails90) AddDeliveringSettlementParties() *SettlementParties49 {
	t.DeliveringSettlementParties = new(SettlementParties49)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails90) AddReceivingSettlementParties() *SettlementParties49 {
	t.ReceivingSettlementParties = new(SettlementParties49)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails90) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}

func (t *TransactionDetails90) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}
