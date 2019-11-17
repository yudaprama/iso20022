package model

// Identifies the details of the transaction.
type TransactionDetails91 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity4Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent16Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails117 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *PlaceOfTradeIdentification2 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafeKeepingPlace2 `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *PlaceOfClearingIdentification1 `xml:"PlcOfClr,omitempty"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity10Choice `xml:"PstngQty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection18 `xml:"PstngAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection59 `xml:"AcrdIntrstAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, that is, payment is effected and securities are delivered.
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate12Choice `xml:"SttlmDt,omitempty"`

	// Date and time assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties49 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties49 `xml:"RcvgSttlmPties,omitempty"`

	// Specifies whether it is the reversal of a previously reported movement.
	ReversalIndicator *YesNoIndicator `xml:"RvslInd,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *RestrictedFINXMax350Text `xml:"TxAddtlDtls,omitempty"`
}

func (t *TransactionDetails91) AddTransactionActivity() *TransactionActivity4Choice {
	t.TransactionActivity = new(TransactionActivity4Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails91) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent16Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent16Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails91) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails91) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails91) AddSettlementParameters() *SettlementDetails117 {
	t.SettlementParameters = new(SettlementDetails117)
	return t.SettlementParameters
}

func (t *TransactionDetails91) AddPlaceOfTrade() *PlaceOfTradeIdentification2 {
	t.PlaceOfTrade = new(PlaceOfTradeIdentification2)
	return t.PlaceOfTrade
}

func (t *TransactionDetails91) AddSafekeepingPlace() *SafeKeepingPlace2 {
	t.SafekeepingPlace = new(SafeKeepingPlace2)
	return t.SafekeepingPlace
}

func (t *TransactionDetails91) AddPlaceOfClearing() *PlaceOfClearingIdentification1 {
	t.PlaceOfClearing = new(PlaceOfClearingIdentification1)
	return t.PlaceOfClearing
}

func (t *TransactionDetails91) AddPostingQuantity() *Quantity10Choice {
	t.PostingQuantity = new(Quantity10Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails91) SetNumberOfDaysAccrued(value string) {
	t.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (t *TransactionDetails91) AddPostingAmount() *AmountAndDirection18 {
	t.PostingAmount = new(AmountAndDirection18)
	return t.PostingAmount
}

func (t *TransactionDetails91) AddAccruedInterestAmount() *AmountAndDirection59 {
	t.AccruedInterestAmount = new(AmountAndDirection59)
	return t.AccruedInterestAmount
}

func (t *TransactionDetails91) AddTradeDate() *TradeDate6Choice {
	t.TradeDate = new(TradeDate6Choice)
	return t.TradeDate
}

func (t *TransactionDetails91) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	t.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return t.EffectiveSettlementDate
}

func (t *TransactionDetails91) AddSettlementDate() *SettlementDate12Choice {
	t.SettlementDate = new(SettlementDate12Choice)
	return t.SettlementDate
}

func (t *TransactionDetails91) AddValueDate() *DateAndDateTimeChoice {
	t.ValueDate = new(DateAndDateTimeChoice)
	return t.ValueDate
}

func (t *TransactionDetails91) AddDeliveringSettlementParties() *SettlementParties49 {
	t.DeliveringSettlementParties = new(SettlementParties49)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails91) AddReceivingSettlementParties() *SettlementParties49 {
	t.ReceivingSettlementParties = new(SettlementParties49)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails91) SetReversalIndicator(value string) {
	t.ReversalIndicator = (*YesNoIndicator)(&value)
}

func (t *TransactionDetails91) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}
