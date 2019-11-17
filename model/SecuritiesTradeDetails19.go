package model

// Details of the securities trade.
type SecuritiesTradeDetails19 struct {

	// Unambiguous identification of the transaction as known by the notification receiver.
	NotificationSenderTransactionIdentification *Max35Text `xml:"NtfctnSndrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the notification receiver.
	NotificationReceiverTransactionIdentification *Max35Text `xml:"NtfctnRcvrTxId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate1Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate1Choice `xml:"SttlmDt"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes20 `xml:"FinInstrmAttrbts,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting1Choice `xml:"Rptg,omitempty"`

	// Details about the financial instrument quantity involved in the transfer.
	QuantityDetails *Quantity5 `xml:"QtyDtls"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails4 `xml:"SttlmParams,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties10 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties10 `xml:"RcvgSttlmPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection7 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts2 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties10 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTradeDetails19) SetNotificationSenderTransactionIdentification(value string) {
	s.NotificationSenderTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails19) SetNotificationReceiverTransactionIdentification(value string) {
	s.NotificationReceiverTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails19) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails19) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesTradeDetails19) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesTradeDetails19) AddTradeDate() *TradeDate1Choice {
	s.TradeDate = new(TradeDate1Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails19) AddSettlementDate() *SettlementDate1Choice {
	s.SettlementDate = new(SettlementDate1Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails19) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails19) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesTradeDetails19) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes20 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes20)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesTradeDetails19) AddReporting() *Reporting1Choice {
	newValue := new(Reporting1Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails19) AddQuantityDetails() *Quantity5 {
	s.QuantityDetails = new(Quantity5)
	return s.QuantityDetails
}

func (s *SecuritiesTradeDetails19) AddSettlementParameters() *SettlementDetails4 {
	s.SettlementParameters = new(SettlementDetails4)
	return s.SettlementParameters
}

func (s *SecuritiesTradeDetails19) AddDeliveringSettlementParties() *SettlementParties10 {
	s.DeliveringSettlementParties = new(SettlementParties10)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeDetails19) AddReceivingSettlementParties() *SettlementParties10 {
	s.ReceivingSettlementParties = new(SettlementParties10)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeDetails19) AddSettlementAmount() *AmountAndDirection7 {
	s.SettlementAmount = new(AmountAndDirection7)
	return s.SettlementAmount
}

func (s *SecuritiesTradeDetails19) AddOtherAmounts() *OtherAmounts2 {
	s.OtherAmounts = new(OtherAmounts2)
	return s.OtherAmounts
}

func (s *SecuritiesTradeDetails19) AddOtherBusinessParties() *OtherParties10 {
	s.OtherBusinessParties = new(OtherParties10)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeDetails19) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
