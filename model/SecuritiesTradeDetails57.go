package model

// Details of the securities trade.
type SecuritiesTradeDetails57 struct {

	// Unambiguous identification of the transaction as known by the notification receiver.
	NotificationSenderTransactionIdentification *RestrictedFINXMax16Text `xml:"NtfctnSndrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the notification receiver.
	NotificationReceiverTransactionIdentification *RestrictedFINXMax16Text `xml:"NtfctnRcvrTxId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Specifies the date/time on which the trade was executed.
	TradeDate *TradeDate6Choice `xml:"TradDt,omitempty"`

	// Date and time at which the securities are to be delivered or received.
	SettlementDate *SettlementDate12Choice `xml:"SttlmDt"`

	// Number of days on which the interest rate accrues (daily accrual note).
	NumberOfDaysAccrued *Max3Number `xml:"NbOfDaysAcrd,omitempty"`

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Specifies that a trade is to be reported to a third party.
	Reporting []*Reporting8Choice `xml:"Rptg,omitempty"`

	// Details about the financial instrument quantity involved in the transfer.
	QuantityDetails *Quantity12 `xml:"QtyDtls"`

	// Parameters applied to the settlement of a security transfer.
	SettlementParameters *SettlementDetails103 `xml:"SttlmParams,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties44 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties44 `xml:"RcvgSttlmPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection57 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts33 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties30 `xml:"OthrBizPties,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesTradeDetails57) SetNotificationSenderTransactionIdentification(value string) {
	s.NotificationSenderTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails57) SetNotificationReceiverTransactionIdentification(value string) {
	s.NotificationReceiverTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails57) SetCommonIdentification(value string) {
	s.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SecuritiesTradeDetails57) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SecuritiesTradeDetails57) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SecuritiesTradeDetails57) AddTradeDate() *TradeDate6Choice {
	s.TradeDate = new(TradeDate6Choice)
	return s.TradeDate
}

func (s *SecuritiesTradeDetails57) AddSettlementDate() *SettlementDate12Choice {
	s.SettlementDate = new(SettlementDate12Choice)
	return s.SettlementDate
}

func (s *SecuritiesTradeDetails57) SetNumberOfDaysAccrued(value string) {
	s.NumberOfDaysAccrued = (*Max3Number)(&value)
}

func (s *SecuritiesTradeDetails57) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesTradeDetails57) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesTradeDetails57) AddReporting() *Reporting8Choice {
	newValue := new(Reporting8Choice)
	s.Reporting = append(s.Reporting, newValue)
	return newValue
}

func (s *SecuritiesTradeDetails57) AddQuantityDetails() *Quantity12 {
	s.QuantityDetails = new(Quantity12)
	return s.QuantityDetails
}

func (s *SecuritiesTradeDetails57) AddSettlementParameters() *SettlementDetails103 {
	s.SettlementParameters = new(SettlementDetails103)
	return s.SettlementParameters
}

func (s *SecuritiesTradeDetails57) AddDeliveringSettlementParties() *SettlementParties44 {
	s.DeliveringSettlementParties = new(SettlementParties44)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesTradeDetails57) AddReceivingSettlementParties() *SettlementParties44 {
	s.ReceivingSettlementParties = new(SettlementParties44)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesTradeDetails57) AddSettlementAmount() *AmountAndDirection57 {
	s.SettlementAmount = new(AmountAndDirection57)
	return s.SettlementAmount
}

func (s *SecuritiesTradeDetails57) AddOtherAmounts() *OtherAmounts33 {
	s.OtherAmounts = new(OtherAmounts33)
	return s.OtherAmounts
}

func (s *SecuritiesTradeDetails57) AddOtherBusinessParties() *OtherParties30 {
	s.OtherBusinessParties = new(OtherParties30)
	return s.OtherBusinessParties
}

func (s *SecuritiesTradeDetails57) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
