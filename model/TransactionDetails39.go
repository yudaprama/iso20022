package model

// Identifies the details of the transaction.
type TransactionDetails39 struct {

	// Specifies the type of activity to which this instruction relates.
	TransactionActivity *TransactionActivity1Choice `xml:"TxActvty"`

	// Choice of type for the transaction reported.
	SettlementTransactionOrCorporateActionEventType *SettlementOrCorporateActionEvent6Choice `xml:"SttlmTxOrCorpActnEvtTp,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails47 `xml:"SttlmParams,omitempty"`

	// Market in which a trade transaction has been executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat3Choice `xml:"SfkpgPlc,omitempty"`

	// Infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
	PlaceOfClearing *AnyBICIdentifier `xml:"PlcOfClr,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Quantity of financial instrument (to be) posted to the safekeeping account.
	PostingQuantity *Quantity6Choice `xml:"PstngQty"`

	// Amount of money that is to be/was posted to the account.
	PostingAmount *AmountAndDirection8 `xml:"PstngAmt,omitempty"`

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
	DeliveringSettlementParties *SettlementParties13 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties13 `xml:"RcvgSttlmPties,omitempty"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionAdditionalDetails *Max350Text `xml:"TxAddtlDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (t *TransactionDetails39) AddTransactionActivity() *TransactionActivity1Choice {
	t.TransactionActivity = new(TransactionActivity1Choice)
	return t.TransactionActivity
}

func (t *TransactionDetails39) AddSettlementTransactionOrCorporateActionEventType() *SettlementOrCorporateActionEvent6Choice {
	t.SettlementTransactionOrCorporateActionEventType = new(SettlementOrCorporateActionEvent6Choice)
	return t.SettlementTransactionOrCorporateActionEventType
}

func (t *TransactionDetails39) SetSecuritiesMovementType(value string) {
	t.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (t *TransactionDetails39) SetPayment(value string) {
	t.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (t *TransactionDetails39) AddSettlementParameters() *SettlementDetails47 {
	t.SettlementParameters = new(SettlementDetails47)
	return t.SettlementParameters
}

func (t *TransactionDetails39) AddPlaceOfTrade() *MarketIdentification4 {
	t.PlaceOfTrade = new(MarketIdentification4)
	return t.PlaceOfTrade
}

func (t *TransactionDetails39) AddSafekeepingPlace() *SafekeepingPlaceFormat3Choice {
	t.SafekeepingPlace = new(SafekeepingPlaceFormat3Choice)
	return t.SafekeepingPlace
}

func (t *TransactionDetails39) SetPlaceOfClearing(value string) {
	t.PlaceOfClearing = (*AnyBICIdentifier)(&value)
}

func (t *TransactionDetails39) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	t.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return t.FinancialInstrumentIdentification
}

func (t *TransactionDetails39) AddPostingQuantity() *Quantity6Choice {
	t.PostingQuantity = new(Quantity6Choice)
	return t.PostingQuantity
}

func (t *TransactionDetails39) AddPostingAmount() *AmountAndDirection8 {
	t.PostingAmount = new(AmountAndDirection8)
	return t.PostingAmount
}

func (t *TransactionDetails39) AddTradeDate() *TradeDate1Choice {
	t.TradeDate = new(TradeDate1Choice)
	return t.TradeDate
}

func (t *TransactionDetails39) AddExpectedSettlementDate() *DateAndDateTimeChoice {
	t.ExpectedSettlementDate = new(DateAndDateTimeChoice)
	return t.ExpectedSettlementDate
}

func (t *TransactionDetails39) AddSettlementDate() *SettlementDate2Choice {
	t.SettlementDate = new(SettlementDate2Choice)
	return t.SettlementDate
}

func (t *TransactionDetails39) AddLateDeliveryDate() *DateAndDateTimeChoice {
	t.LateDeliveryDate = new(DateAndDateTimeChoice)
	return t.LateDeliveryDate
}

func (t *TransactionDetails39) AddExpectedValueDate() *DateAndDateTimeChoice {
	t.ExpectedValueDate = new(DateAndDateTimeChoice)
	return t.ExpectedValueDate
}

func (t *TransactionDetails39) AddDeliveringSettlementParties() *SettlementParties13 {
	t.DeliveringSettlementParties = new(SettlementParties13)
	return t.DeliveringSettlementParties
}

func (t *TransactionDetails39) AddReceivingSettlementParties() *SettlementParties13 {
	t.ReceivingSettlementParties = new(SettlementParties13)
	return t.ReceivingSettlementParties
}

func (t *TransactionDetails39) SetTransactionAdditionalDetails(value string) {
	t.TransactionAdditionalDetails = (*Max350Text)(&value)
}

func (t *TransactionDetails39) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	t.SupplementaryData = append(t.SupplementaryData, newValue)
	return newValue
}
