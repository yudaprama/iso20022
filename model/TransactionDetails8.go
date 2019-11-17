package model

// Identifies the details of the transaction.
type TransactionDetails8 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity1Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent2Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails7 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection8 `xml:"PstngAmt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	ExpectedSettlementDate *DateAndDateTimeChoice `xml:"XpctdSttlmDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate2Choice `xml:"SttlmDt"`

	// Date and time after the settlement date specified in the trade, used for pool trades resulting from the original To Be Assigned (TBA) securities.
	LateDeliveryDate *DateAndDateTimeChoice `xml:"LateDlvryDt,omitempty"`

	// For against payment transactions, the value date/time at which the account servicer expects the settlement amount to be credited or debited.
	ExpectedValueDate *DateAndDateTimeChoice `xml:"XpctdValDt,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties2 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties2 `xml:"RcvgSttlmPties,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *Max350Text `xml:"TxAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension2 `xml:"Xtnsn,omitempty"`
}

func (t *TransactionDetails8) AddTransactionActivity() *TransactionActivity1Choice {
	t.TransactionActivity = new(TransactionActivity1Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails8) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent2Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent2Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails8) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails8) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails8) AddSettlementParameters() *SettlementDetails7 {
	t.SettlementParameters = new(SettlementDetails7)
	return t.SettlementParameters
}

func (t *TransactionDetails8) AddPlaceOfTrade() *MarketIdentification4 {
	t.PlaceOfTrade = new(MarketIdentification4)
	return t.PlaceOfTrade
}

func (t *TransactionDetails8) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails8) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails8) AddPostingQuantity() *Quantity6Choice {
	t.PostingQuantity = new(Quantity6Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails8) AddPostingAmount() *AmountAndDirection8 {
	t.PostingAmount = new(AmountAndDirection8)
	return t.PostingAmount
}

func (t *TransactionDetails8) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails8) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails8) AddSettlementDate() *SettlementDate2Choice {
	t.SettlementDate = new(SettlementDate2Choice)
	return t.SettlementDate
}

func (t *TransactionDetails8) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails8) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails8) AddDeliveringSettlementParties() *SettlementParties2 {
	t.DeliveringSettlementParties = new(SettlementParties2)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails8) AddReceivingSettlementParties() *SettlementParties2 {
	t.ReceivingSettlementParties = new(SettlementParties2)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails8) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*Max350Text)(&value)
}

func (t *TransactionDetails8) AddExtension() *Extension2 {
	newValue := new(Extension2)
	t.Extension = append(t.Extension, newValue)
	return newValue
}
