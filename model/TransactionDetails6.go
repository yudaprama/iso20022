package model

// Identifies the details of the transaction.
type TransactionDetails6 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity1Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent1Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails2 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and wich facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection3 `xml:"PstngAmt,omitempty"`

	// Interest amount that has accrued in between coupon payment periods.
	AccruedInterestAmount *AmountAndDirection4 `xml:"AcrdIntrstAmt,omitempty"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, ie, payment is effected and securities are delivered.
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt,omitempty"`

	// Date and time assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	ValueDate *DateAndDateTimeChoice `xml:"ValDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties2 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties2 `xml:"RcvgSttlmPties,omitempty"`

	// Specifies whether it is the reversal of a previously reported movement.
	ReversalIndicator *YesNoIndicator `xml:"RvslInd,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *Max350Text `xml:"TxAddtlDtls,omitempty"`
}

func (t *TransactionDetails6) AddTransactionActivity() *TransactionActivity1Choice {
	t.TransactionActivity = new(TransactionActivity1Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails6) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent1Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent1Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails6) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails6) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails6) AddSettlementParameters() *SettlementDetails2 {
	t.SettlementParameters = new(SettlementDetails2)
	return t.SettlementParameters
}

func (t *TransactionDetails6) AddPlaceOfTrade() *MarketIdentification4 {
	t.PlaceOfTrade = new(MarketIdentification4)
	return t.PlaceOfTrade
}

func (t *TransactionDetails6) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails6) SetPlaceOfClearing(value string) {
	t.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (t *TransactionDetails6) AddPostingQuantity() *Quantity6Choice {
	t.PostingQuantity = new(Quantity6Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails6) SetNumberOfDaysAccrued(value string) {
	t.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (t *TransactionDetails6) AddPostingAmount() *AmountAndDirection3 {
	t.PostingAmount = new(AmountAndDirection3)
	return t.PostingAmount
}

func (t *TransactionDetails6) AddAccruedInterestAmount() *AmountAndDirection4 {
	t.AccruedInterestAmount = new(AmountAndDirection4)
	return t.AccruedInterestAmount
}

func (t *TransactionDetails6) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails6) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	t.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return t.EffectiveSettlementDate
}

func (t *TransactionDetails6) AddSettlementDate() *SettlementDate1Choice {
	t.SettlementDate = new(SettlementDate1Choice)
	return t.SettlementDate
}

func (t *TransactionDetails6) AddValueDate() *DateAndDateTimeChoice {
	t.ValueDate = new(DateAndDateTimeChoice)
	return t.ValueDate
}

func (t *TransactionDetails6) AddDeliveringSettlementParties() *SettlementParties2 {
	t.DeliveringSettlementParties = new(SettlementParties2)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails6) AddReceivingSettlementParties() *SettlementParties2 {
	t.ReceivingSettlementParties = new(SettlementParties2)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails6) SetReversalIndicator(value string) {
	t.ReversalIndicator = (*YesNoIndicator)(&value)
}

func (t *TransactionDetails6) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*Max350Text)(&value)
}
