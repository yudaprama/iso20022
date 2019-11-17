package model

// Details of the securities trade.
type SecuritiesTradeDetails48 struct {

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
	TradeDate *TradeDate5Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate9Choice `xml:"SttlmDt"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting7Choice `xml:"Rptg,omitempty"`

	// Details about the financial instrument quantity involved in the transfer.
	QuantityDetails *Quantity11 `xml:"QtyDtls"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails100 `xml:"SttlmParams,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties36 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties36 `xml:"RcvgSttlmPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection52 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts29 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties26 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTradeDetails48) SetNotificationSenderTransactionIdentification(value string) {
	s.NotificationSenderTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails48) SetNotificationReceiverTransactionIdentification(value string) {
	s.NotificationReceiverTransactionIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails48) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesTradeDetails48) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesTradeDetails48) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesTradeDetails48) AddTradeDate() *TradeDate5Choice {
	s.TradeDate = new(TradeDate5Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails48) AddSettlementDate() *SettlementDate9Choice {
	s.SettlementDate = new(SettlementDate9Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails48) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails48) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesTradeDetails48) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesTradeDetails48) AddReporting() *Reporting7Choice {
	newValue := new(Reporting7Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails48) AddQuantityDetails() *Quantity11 {
	s.QuantityDetails = new(Quantity11)
	return s.QuantityDetails
}

func (s *SecuritiesTradeDetails48) AddSettlementParameters() *SettlementDetails100 {
	s.SettlementParameters = new(SettlementDetails100)
	return s.SettlementParameters
}

func (s *SecuritiesTradeDetails48) AddDeliveringSettlementParties() *SettlementParties36 {
	s.DeliveringSettlementParties = new(SettlementParties36)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeDetails48) AddReceivingSettlementParties() *SettlementParties36 {
	s.ReceivingSettlementParties = new(SettlementParties36)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeDetails48) AddSettlementAmount() *AmountAndDirection52 {
	s.SettlementAmount = new(AmountAndDirection52)
	return s.SettlementAmount
}

func (s *SecuritiesTradeDetails48) AddOtherAmounts() *OtherAmounts29 {
	s.OtherAmounts = new(OtherAmounts29)
	return s.OtherAmounts
}

func (s *SecuritiesTradeDetails48) AddOtherBusinessParties() *OtherParties26 {
	s.OtherBusinessParties = new(OtherParties26)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeDetails48) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
